package boolgebra

// expand X so that:
//
// there is at most one TypeOr element in the resulting tree, and it stands at the top of it.
//
// TypeNot Expressions only contains Identifiers
func expand(X Expression) Expression {

	switch X.Type() {

	case TypeNot:
		// expanding Not(Y) depends on Y's Type
		Y := X.Elems()[0]
		Y = expand(Y)
		switch Y.Type() {

		case TypeNot:
			// Not(Not(u)) = u
			return Y.Elems()[0]

		case TypeLiteral:
			return Lit(!Y.Val())

		case TypeOr:
			//progate Not to all elems
			// Not( Or(A,B)) = And(Not(A), Not(B))
			// so I will apply the function 'Not' to every
			// element of Y
			// the resulting 'And' might need further expansion
			// return expand(And(fmap(Y.Elems(), Not)...))
			return expand(fmapreduce(Not, iand, Y.Elems()...))

		case TypeAnd:
			// idem as for TypeOr, but the basic rule is now
			// Not( And(A,B)) = Or(Not(A), Not(B))
			// return expand(Or(fmap(Y.Elems(), Not)...))
			return expand(fmapreduce(Not, ior, Y.Elems()...))

		default:
			// only identifier remain
			return X
		}

	case TypeOr:
		// expanding Or(x...)
		// Or is already the top level operand
		// so I just need to expand elementwise, into 'y'
		return fmapreduce(expand, ior, X.Elems()...)

	case TypeAnd:

		// expanding And(x, ...)
		// And need to be distributed, i.e it can return an and iif all terms (expanded)
		// are Ands, or below
		return fmapreduce(expand, expandAnd, X.Elems()...)

	default:
		return X
	}
}

// ior, for inline Or, build a new expression equivalent to Or(x,y)
//
// if x or y are nil, they are ignored
//
// if x or y are already TypeOr expression, then they are reused
func ior(x, y Expression) Expression {
	// this is used in freduce, and y initial value can be nil
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}
	//append all terms, "openning" Ors
	result := append(forceOp(TypeOr, x), forceOp(TypeOr, y)...)
	return Or(result...)
}

// iand, for inline And, build a new expression equivalent to And(x,y)
//
// if x or y are nil, they are ignored
//
// if x or y are already TypeAnd expression, then they are reused
func iand(x, y Expression) Expression {
	// this is used in freduce, and y initial value can be nil
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}
	//append all terms, "openning" Ands
	result := append(forceOp(TypeAnd, x), forceOp(TypeAnd, y)...)
	return And(result...)
}

// expandAnd is like iand, a reduce function for And binary operator, that also takes care of expansion of and over 'Or'
func expandAnd(x, y Expression) Expression {
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}
	// there are two possible outcome, either x, or y is an Or, then we are going to build an OR,
	// or we are going to build and And (just like 'iand' above)

	if x.Type() == TypeOr || y.Type() == TypeOr {
		// we are going to build an or at the end
		// get both left and rights as generic Ors
		lefts, rights := forceOp(TypeOr, x), forceOp(TypeOr, y)

		// double sweep on left cross right
		result := make([]Expression, 0, len(lefts)*len(rights))
		for _, left := range lefts {
			for _, right := range rights {
				result = append(result, And(left, right))
			}
		}
		return Or(result...)

	} else {
		// none is an Or, therefore the outcome will be and And, this is a simple case
		return iand(x, y)
	}
}

// forceOp returns x as a slice of op terms.
//
// if x is an 'op' return Elems()
// otherwise returns []{x}
//
// this is a trick to handle Or (resp And) expression, and non-Or (resp And) expression as if non-Or expression where slice of size 1
func forceOp(op Type, x Expression) []Expression {

	if x.Type() == op {
		return x.Elems()
	} else {
		return []Expression{x}
	}
}

// fmap (functional map) map 'f' element wise, the function 'f' to every Expression in 'expressions'
//
// The result is sorted according to the natural order of Expressions
func fmap(f func(exp Expression) Expression, expressions ...Expression) []Expression {
	res := make([]Expression, len(expressions))
	for i, n := range expressions {
		res[i] = f(n)
	}
	return res
}

// freduce 'expression' by applying 'f(x, y)'
//
//     freduce(f, a,b,c,d)
//      = f(a,b) , c, d
//      = f( f(a,b), c), d
//      = f( f( f(a,b), c), d)
func freduce(f func(exp1, exp2 Expression) Expression, exp ...Expression) (res Expression) {
	switch len(exp) {
	case 0:
		return
	case 1:
		return exp[0]
	default:
		res = exp[0]
		for _, node := range exp[1:] {
			res = f(res, node)
		}
	}
	return
}

//fmapreduce combines a fmap, and a freduce.
func fmapreduce(mapf func(exp Expression) Expression, reducef func(exp1, exp2 Expression) Expression, exp ...Expression) Expression {
	return freduce(reducef, fmap(mapf, exp...)...)
}
