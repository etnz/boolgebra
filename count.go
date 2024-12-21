package boolgebra

import "github.com/etnz/permute"

// count is a family of boolean algebra function to express quantities of true.

// AtLeast retuns an expression that is true if and only if at least 'n' terms are True.
func AtLeast(n int, terms ...Expr) Expr {
	// to assess that at least n terms in 'terms' are true
	// in one of the n-subsets of terms they must all be true.
	//
	// e.g. in AtLeast(2, a,b,c):
	// a&b | a&c | b&c must be true.
	if n == 0 {
		return Lit(true)
	}
	if n > len(terms) {
		return Lit(false)
	}
	if n == len(terms) {
		return And(terms...)
	}

	expr := Lit(false) // default for Or
	// Always start with the identity
	p := permute.New(n)
	// loop over all n-subset.
	for ok := true; ok; ok = nextsubset(p, len(terms)) {
		res := make([]Expr, 0, len(p))
		for _, i := range p {
			res = append(res, terms[i])
		}
		expr = Or(expr, And(res...))
	}
	return expr
}

// AtMost retuns an expression that is true if and only if at most 'n' terms are True.
func AtMost(n int, terms ...Expr) Expr {
	// m is the complement's size.
	m := len(terms) - n
	if m < 0 {
		return Lit(true)
	}

	// assess that among the negations of terms, there are atleast m
	negs := make([]Expr, 0, m)
	for _, x := range terms {
		negs = append(negs, Not(x))
	}
	return AtLeast(m, negs...)
}

// Exactly retuns an expression that is true if and only if exactly 'i' terms are True.
func Exactly(n int, terms ...Expr) Expr {
	// Exactly is imply Atleast N and at Most N
	return And(AtLeast(n, terms...), AtMost(n, terms...))
}

func nextsubset(p []int, n int) bool {
	return permute.SubsetRevolvingDoorNext(p, n, &[2]int{0, 0})
}
