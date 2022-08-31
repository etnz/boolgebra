package boolgebra

// secondary.go holds functions to handle Expr that are implemented
// using primary ones

// Impl returns the logical implication of A,B
//
// see https://en.wikipedia.org/wiki/Boolean_algebra#Secondary_operations
func Impl[T comparable](A, B Expr[T]) Expr[T] {
	return Or(Not(A), B)
}

// Eq returns the logical equality of A,B ( A and B have both the same truth value)
//
// It can also be the logical equivalence A <=> B. Both are in fact the same boolean function.
//
// see https://en.wikipedia.org/wiki/Boolean_algebra#Secondary_operations
func Eq[T comparable](A, B Expr[T]) Expr[T] { return Or(And(Not(A), Not(B)), And(A, B)) }

// Xor returns the logical Xor
//
// see https://en.wikipedia.org/wiki/Boolean_algebra#Secondary_operations
func Xor[T comparable](A, B Expr[T]) Expr[T] { return Or(And(A, Not(B)), And(Not(A), B)) }

// Neq returns the logical '!='
//
// It is the same as Xor.
//
// see https://en.wikipedia.org/wiki/Boolean_algebra#Secondary_operations
func Neq[T comparable](A, B Expr[T]) Expr[T] { return Xor(A, B) }
