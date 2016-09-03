package boolgebra

import "fmt"

func ExampleString() {
	A, B := ID("A"), ID("B")
	X := Or(And(A, B), And(A, Not(B)))
	fmt.Println(String(X))
	// Output:
	// Or(And(A, B), And(A, Not(B)))
}
