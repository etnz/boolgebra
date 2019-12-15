package boole

import "log"

// Expr represent any boolean algebra expression.
//
// This is not an abstract syntax tree.
type Expr struct{ x expression }

func newExpr(x expression) Expr { return Expr{x} }

// String represent this Expression for debug only
func (x Expr) String() string { return x.x.String() }

// Lit returns an Expr equivalent to a boolean literal 'val'
func Lit(val bool) Expr { return newExpr(lit(val)) }

// ID returns an Expr equivalent to a single ID 'id'
func ID(id string) Expr { return newExpr(expression{minterm{id: true}}) }

// Or return the conjunction of all the expression passed in parameter.
//
// By convention, if 'x' is empty it returns Lit(false). See https://en.wikipedia.org/wiki/Empty_sum
func Or(x ...Expr) Expr {
	// start with the neutral of the Or i.e a false
	res := lit(false)
	for _, exp := range x {
		res = or(res, exp.x)
	}
	return newExpr(res)
}

// And returns the disjunction of all the expressions passed in parameters.
//
// By convention, if 'x' is empty it returns Lit(true). See  https://en.wikipedia.org/wiki/Empty_product
func And(x ...Expr) Expr {
	// start with the neutral of the and i.e a true
	res := lit(true)
	for _, exp := range x {
		res = and(res, exp.x)
	}
	return newExpr(res)
}

// Not returns the negation of 'x'.
func Not(x Expr) Expr { return newExpr(not(x.x)) }

// Simplify returns a simpler version of 'x' by applying simplification rules.
//
// So far there is no guarantee that the result is minimal (cannot be further simplified)
func Simplify(x Expr) Expr {
	return Expr{reduce(x.x)}
}

// Exactly retuns an expression that is true if and only if exactly 'i' terms are True.
//
// This is the Or() of and And() of all the i-subsets of terms
func Exactly(i int, terms ...Expr) Expr {
	return quantified(i, identity, Not, terms...)
}

// AtMost retuns an expression that is true if and only if at most 'i' terms are True.
func AtMost(i int, terms ...Expr) Expr {
	return quantified(i, truth, Not, terms...)
}

// AtLeast retuns an expression that is true if and only if at least 'i' terms are True.
func AtLeast(i int, terms ...Expr) Expr {
	return quantified(i, identity, truth, terms...)
}

func identity(x Expr) Expr { return x }
func truth(x Expr) Expr    { return Lit(true) }

// quantifier returns Or ( And( f(p,p')) )where:
// p is a subset of terms, p' is the complement ( remain terms)
// and f is either:
// - for Exactly we need to build a minterm : with all terms in p, and their negation in p'
// - for AtLeast: we just need to build with those in p
// - for Atmost: we just need to build those with p'
//
// therefore we defined two simple function fp(Expr) and fp' accordingly
func quantified(i int, f, g func(x Expr) Expr, terms ...Expr) Expr {

	p := subid(i)  // the slice of indices starting at identify, always
	var ors []Expr // all ands to be Or()ed

	for ok := true; ok; ok = nextsubset(p, len(terms)) {
		res := make([]Expr, 0, len(terms))
		c := complement(p, len(terms))

		for _, i := range p {
			res = append(res, f(terms[i]))
		}

		for _, i := range c {
			res = append(res, g(terms[i]))
		}
		ors = append(ors, And(res...))
	}
	return Or(ors...)
}

// Factor computes the greatest common factor between of terms in x ( x is a sum of prod)
func Factor(x Expr) Expr {
	var res minterm
	for i, m := range x.x {
		if i == 0 {
			// special case for the first one, need to init the thing
			res = m
		}
		res = inter(res, m)
		log.Printf("current factor = %v", res)
		if len(res) == 0 {
			return Expr{expression{res}} // empty one
		}
	}
	return Expr{expression{res}}

}
