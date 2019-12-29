package boolgebra

// primary.go contains functions/methods for the Expr type based on the
// expression/minterm types

// Lit returns an Expr equivalent to a boolean literal 'val'
func Lit(val bool) Expr {
	if val {
		return minterm{} // true is by definition an empty minterm ( neutral for product)
	} else {
		return expression{} // false is an empty expression (neutral for sum)
	}
}

// ID returns an Expr equivalent to a single ID 'id'
func ID(id string) Expr { return minterm{id: true} }

// Or return the conjunction of all the expression passed in parameter.
//
// By convention, if 'x' is empty it returns Lit(false). See https://en.wikipedia.org/wiki/Empty_sum
func Or(x ...Expr) Expr {
	// start with the neutral of the Or i.e a false
	res := make(expression, 0)
	// scan all terms, in all expr
	for _, exp := range x {
		for i := 0; i < exp.Terms(); i++ {
			t := exp.Term(i)
			if t.Is(true) {
				return t //
			}
			if !t.Is(false) { // if this is the literal false, we can just skip it
				res = append(res, t.(minterm))
			}
		}
	}
	if len(res) == 1 {
		return res[0]
	}
	return res
}

// And returns the disjunction of all the expressions passed in parameters.
//
// By convention, if 'x' is empty it returns Lit(true). See  https://en.wikipedia.org/wiki/Empty_product
func And(expressions ...Expr) Expr {

	if len(expressions) == 0 {
		return Lit(true) // return the neutral of And operation by convention
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

	if x.Is(false) || y.Is(false) {
		return Lit(false)
	}

	if x.Is(true) {
		return y
	}
	if y.Is(true) {
		return x
	}

	// general case
	z := make(expression, 0, x.Terms()*y.Terms())
	// this is the big one: all terms from x multiplied by terms from y
	for i := 0; i < x.Terms(); i++ {
		m := x.Term(i).(minterm)

	product:
		for j := 0; j < y.Terms(); j++ {
			n := y.Term(j).(minterm)

			// compute the real m && n , this is basically a merge of all IDs
			// there is one special case: A & A' = false
			for k, v := range m {
				if w, exists := n[k]; exists && w != v {
					// the ID k exists in n too, and their values ( A or A') are different
					continue product
				}
			}

			// basic merge
			o := make(minterm)
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
func Not(x Expr) Expr { return x.Not() }

// Simplify returns a simpler version of 'x' by applying simplification rules.
func Simplify(x Expr) Expr {
	switch e := x.(type) {
	case expression:
		return reduce(e)
	default:
		return x // unchanged
	}
}

// Factor computes the greatest common factor between terms of x
//
// x = \sum t_i \arrow f \and \sum r_i
//
// x is currently a sum of terms, this function returns f and rem so that
//
//    x = And(f, rem)
//    f.Terms() ==1 : it's a minterm
//
func Factor(x Expr) (f, rem Expr) {
	var res minterm
	for i := 0; i < x.Terms(); i++ {
		m := x.Term(i).(minterm)
		if i == 0 {
			// special case for the first one, need to init the thing
			res = m
		}
		res = inter(res, m)
		if len(res) == 0 {
			return expression{res}, x // empty one
		}
	}
	// now for each minterm recompute the reminder

	r := expression{}
	for i := 0; i < x.Terms(); i++ {
		m := x.Term(i).(minterm)
		r = append(r, div(m, res))
	}

	return expression{res}, r

}
