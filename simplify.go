package boolgebra

// Simplify X by applying simplifications rules.
//
// The goal is to apply all possible simplification rules. Other rules like
// Commutativity, or Distributivity do not reduce complexity ( some even increases it),
// instead they operate like a tranformation that enable some simplification rule.
//
func Simplify(X Expression) Expression {
	// make sure that the X is in the best possible normal form
	X = expand(X)
	X = localSimplify(X)
	X = globalSimplify(X)

	return X
}

// simplifySpace explore the expression space through distributions,
// and bring back simplifications from there
func globalSimplify(X Expression) Expression {

	switch X.Type() {
	// distributivity happens only in Ands and Ors
	// to avoid implementing it twice, And case use the dual tranformation, back and forth
	case TypeAnd:
		Z := dual(X)                  // use the dual property to use 'Or' simplifications
		Z = globalSimplify(expand(Z)) // simplify this Or expr
		return dual(Z)                // dual back
	case TypeOr:

		X = fmapreduce(globalSimplify, ior, X.Elems()...)
		initialLength := len(X.Elems())
		for {
			// X has changed, it might benefit from local simpl
			X = localSimplify(X)

			// now we need to explore through distributivity
			// so we are going to pick one monomial M from the expression
			// and build Y,Z:
			// X = M and Y or Z
			// i.e factorizing as much as possible X
			// Z is not new, but Y is a totally different expression, and might contains some simplifications
			//
			// Therefore the strategy is to recurse into Y, and rebuild X with the 'simplified' version

			// try each monomial, factor by It, and simplify the new 'factored' expression
			// AB A!B C  ->  A.(  B  !B ) C
			// here   'B  !B' is a new expression, and just need to be simplified
			//log.Printf("Scanning monomials for %v", Or(xs...))
			for _, m := range monomials(X) {

				div, mod := divmod(X, m)

				// do the full simplification in this sub expression
				div = localSimplify(div)
				div = globalSimplify(div)

				// and rebuild the expression, now
				// X = m and div  or Mod

				q, _ := divmod(mod, expand(Not(m)))
				if q != False {
					X = localSimplify(expand(Or(div, mod)))
				} else {
					X = localSimplify(expand(Or(And(m, div), mod)))
				}

			}
			if len(X.Elems()) == initialLength {
				//there have NOT been any simplifications, stop it
				return X
			}
			return X
		}

	default: // no possible simplifications
		return X
	}
}

// localSimplify applies all simplification that match on X.
func localSimplify(X Expression) Expression {
	switch X.Type() {

	case TypeNot:
		// recursing into it first
		Y := localSimplify(X.Elems()[0])
		switch Y.Type() {
		case TypeLiteral:
			//      Not(False)   -> True
			//      Not(True)    -> False
			return Lit(!Y.Val())

		case TypeNot:
			//      Not(Not(x))  -> x
			return Y.Elems()[0]
		}

	case TypeAnd:
		// transform X to its dual form, to match rules on it
		// just because rules involving operators are all implemented using Or, only
		return dual(localSimplify(dual(X)))

	case TypeOr:

		// recurse first
		xs := fmap(localSimplify, X.Elems()...)

		for i := 0; i < len(xs); i++ {
			x := xs[i]

			if x.Type() == TypeLiteral && x.Val() {
				//      Or(True, x)  -> True          Annihilator
				// nb x.Val() is enough to test that x is a literal, and true
				return True
			}

			if x.Type() == TypeLiteral && !x.Val() {
				//      Or(False,x)  -> x             Identity
				xs = remove(xs, i)
				i--
				continue
			}

			// z is the current X without the current 'x'
			// X = x Or z
			z := Or(remove(xs, i)...) // remove copy the original slice, so xs will be unspoiled

			// if I can  write z as '!x Or y', then
			// => X = x Or (!x Or y)
			//      Or(x, Not(x)) -> True         Complementation
			notx := expand(Not(x))
			if findAssociation(notx, z) {
				return True
			}

			// if I can  write z as 'x Or y', then
			// => X = x Or (x Or y)
			//      Or(x, x) -> x     Idempotence
			if findAssociation(x, z) {
				xs = remove(xs, i)
				i--
				continue

			}
		}

		return Or(xs...)

	}
	return X
}

// SimplifyAll execute Simplify elementwise
func SimplifyAll(X ...Expression) []Expression { return fmap(Simplify, X...) }

// remove the ith value, is the complement of append
func remove(w []Expression, i int) []Expression {
	// protect against i outside of the box
	// if i < 0 {
	// 	i = len(w) + i%len(w)
	// } else {
	// 	i = i % len(w)
	// }

	last := len(w) - 1
	//allocate the final copy
	z := make([]Expression, last)
	if i > 0 {
		// there is a heading part, copy it
		copy(z[:i], w[:i])
	}
	if i < last {
		// there is a trailing part, copy it
		copy(z[i:], w[i+1:])
	}
	return z
}

func dual(X Expression) Expression {
	switch X.Type() {
	case TypeOr:
		return fmapreduce(dual, iand, X.Elems()...)
	case TypeAnd:
		return fmapreduce(dual, ior, X.Elems()...)
	case TypeNot:
		return Not(dual(X.Elems()[0]))
	case TypeLiteral:
		return Lit(!X.Val())

	default:
		return X
	}
}

// divmod returns the couple (q,m)quotient and modulo so that
//  X = Y and q or m
func divmod(X, Y Expression) (q, mod Expression) {
	//trivial cases first
	if Equals(X, Y) {
		return True, False
	}

	if X.Type() == TypeOr {
		qs := make([]Expression, 0)
		ms := make([]Expression, 0)
		for _, x := range X.Elems() {
			q, m := divmod(x, Y)
			qs, ms = append(qs, q), append(ms, m)
		}

		return localSimplify(Or(qs...)), localSimplify(Or(ms...))
	}

	if X.Type() == TypeAnd {
		for i, x := range X.Elems() {
			if Equals(x, Y) {
				// nb : as it is, this works only if Y is a monomial, so don't generalize the function too quickly
				// there is one ! great build up the result
				return And(remove(X.Elems(), i)...), False
			}
		}
	}
	// bad luck, no hit
	return False, X
}

// findAssociation computes 'y' if possible so that
//
//    z = x Or y
//
// if no possible returns (false)
func findAssociation(X, Z Expression) (ok bool) {

	xs := forceOp(TypeOr, X)
	zs := forceOp(TypeOr, Z)
	if len(zs) < len(xs) {
		// there is no chance that zs can be written xs + something !
		return false
	}

	for _, x := range xs {
		found := false // asumme I've not found it
		for i, z := range zs {

			// pb: equals is not strong enough, it need to NOT consider the And order
			if Equals(x, z) {
				// I've found it
				zs = remove(zs, i)
				found = true
				break
			}
		}
		if !found {
			// unfortunately, this x cannot be found in z
			return false
		}
	}
	return true

}
