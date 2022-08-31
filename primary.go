package boolgebra

// primary.go contains functions/methods for the Expr type based on the
// expression/minterm types

// Lit returns an Expr equivalent to a boolean literal 'val'
func Lit[T comparable](val bool) Expr[T] {
	if val {
		return Expr[T]{Term[T]{}} // true is by definition an empty minterm ( neutral for product)
	} else {
		return Expr[T]{} // false is an empty expression (neutral for sum)
	}
}

// ID returns an Expr equivalent to a single ID 'id'
func ID[T comparable](id T) Expr[T] { return Expr[T]{Term[T]{id: true}} }

// Or return the conjunction of all the expression passed in parameter.
//
// By convention, if 'x' is empty it returns Lit(false). See https://en.wikipedia.org/wiki/Empty_sum
func Or[T comparable](x ...Expr[T]) Expr[T] {
	// start with the neutral of the Or i.e a false
	res := make(Expr[T], 0)
	// scan all terms, in all expr
	for _, exp := range x {
		for _, t := range exp {
			if t.isLiteral(true) {
				return Expr[T]{Term[T]{}}
			}
			if !t.isLiteral(false) { // if this is the literal false, we can just skip it
				res = append(res, t)
			}
		}
	}
	if len(res) == 1 {
		return Expr[T]{res[0]}
	}
	return res
}

// And returns the disjunction of all the expressions passed in parameters.
//
// By convention, if 'x' is empty it returns Lit(true). See  https://en.wikipedia.org/wiki/Empty_product
func And[T comparable](expressions ...Expr[T]) Expr[T] {

	if len(expressions) == 0 {
		return Lit[T](true) // return the neutral of And operation by convention
	}
	if len(expressions) == 1 {
		return expressions[0] // another common degenerated case
	}
	if len(expressions) > 2 { // general case, that rely on the case len() == 2 below
		// will recurse to the next case
		res := expressions[0]
		for i := 1; i < len(expressions); i++ {
			res = And(res, expressions[i])
		}
		return res
	}
	//if len(expressions) == 2 {
	// this is the only real case
	x, y := expressions[0], expressions[1]

	if x.isLiteral(false) || y.isLiteral(false) {
		return Lit[T](false)
	}

	if x.isLiteral(true) {
		return y
	}
	if y.isLiteral(true) {
		return x
	}

	// general case
	z := make(Expr[T], 0, len(x)*len(y))
	// this is the big one: all terms from x multiplied by terms from y
	for _, m := range x {

	product:
		for _, n := range y {

			// compute the real m && n , this is basically a merge of all IDs
			// there is one special case: A & A' = false
			for k, v := range m {
				if w, exists := n[k]; exists && w != v {
					// the ID k exists in n too, and their values ( A or A') are different
					continue product
				}
			}

			// basic merge
			o := make(Term[T])
			for k, v := range m {
				o[k] = v
			}
			for k, v := range n {
				o[k] = v
			}

			z = append(z, o)
		}
	}

	return z
}

// Not returns the negation of 'x'.
func Not[T comparable](x Expr[T]) Expr[T] { return x.Not() }

// Factor computes the greatest common factor between terms of x
//
// x = \sum t_i \arrow f \and \sum r_i
//
// x is currently a sum of terms, this function returns f and rem so that
//
//	x = And(f, rem)
//	f.Terms() ==1 : it's a minterm
func Factor[T comparable](x Expr[T]) (f, rem Expr[T]) {
	var res Term[T]
	for i, m := range x {
		if i == 0 {
			// special case for the first one, need to init the thing
			res = m
		}
		res = inter(res, m)
		if len(res) == 0 {
			return Expr[T]{res}, x // empty one
		}
	}
	// now for each minterm recompute the reminder

	r := Expr[T]{}
	for _, m := range x {
		r = append(r, div(m, res))
	}

	return Expr[T]{res}, r

}
