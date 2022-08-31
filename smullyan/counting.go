package smullyan

import (
	. "github.com/etnz/boolgebra"
	"github.com/etnz/permute"
)

//counting.go holds the function relative to counting, like Exactly or AtMost

// Exactly retuns an expression that is true if and only if exactly 'i' terms are True.
//
// This is the Or() of and And() of all the i-subsets of terms
func Exactly[T comparable](i int, terms ...Expr[T]) Expr[T] {
	return quantified(i, identity[T], Not[T], terms...)
}

// AtMost retuns an expression that is true if and only if at most 'i' terms are True.
func AtMost[T comparable](i int, terms ...Expr[T]) Expr[T] {
	return quantified(i, truth[T], Not[T], terms...)
}

// AtLeast retuns an expression that is true if and only if at least 'i' terms are True.
func AtLeast[T comparable](i int, terms ...Expr[T]) Expr[T] {
	return quantified(i, identity[T], truth[T], terms...)
}

func identity[T comparable](x Expr[T]) Expr[T] { return x }
func truth[T comparable](x Expr[T]) Expr[T]    { return Lit[T](true) }

// quantifier returns Or ( And( f(p,p')) )where:
// p is a subset of terms, p' is the complement ( remain terms)
// and f is either:
// - for Exactly we need to build a minterm : with all terms in p, and their negation in p'
// - for AtLeast: we just need to build with those in p
// - for Atmost: we just need to build those with p'
//
// therefore we defined two simple function fp(Expr) and fp' accordingly
func quantified[T comparable](i int, f, g func(x Expr[T]) Expr[T], terms ...Expr[T]) Expr[T] {

	p := subid(i)     // the slice of indices starting at identify, always
	var ors []Expr[T] // all ands to be Or()ed

	for ok := true; ok; ok = nextsubset(p, len(terms)) {
		res := make([]Expr[T], 0, len(terms))
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

// subid return the identity permutation
func subid(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = i
	}
	return res
}

func complement(p []int, n int) (c []int) {
	i := 0 // the absolute index, that will sweep all values from 0 to n
	for _, pi := range p {
		for ; i < pi; i++ {
			// the next item to belong to p is pi, i was the last one, so all between belong to the complement
			c = append(c, i)
		}
		i++
	}
	// do til the end
	for ; i < n; i++ {
		// the next item to belong to p is pi, i was the last one, so all between belong to the complement
		c = append(c, i)
	}
	return
}

// nextsubset updates 'p' a 'subset' and return true until it has gone through all
// the possible subsets.
//
// 'p' is a slice of indices from the original list of items (of length 'n')
//
// The first subset to use is [ 0, 1, 2, ...].
func nextsubset(p []int, n int) bool {
	// s := &[2]int{0, 0}
	return permute.SubsetRevolvingDoorNext(p, n, &[2]int{0, 0})
}
