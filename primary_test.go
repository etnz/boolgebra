package boolgebra

import (
	"fmt"
	"testing"
)

func ExampleID() {
	A := ID("A")
	fmt.Println(A)
	//Output: "A"
}

// TesLit ensure that the basic true and false are working accordingly with Is(bool)
func TestLit(t *testing.T) {
	if !Lit[string](true).isLiteral(true) {
		t.Error("Lit(true).Is(true) must be true")
	}
	if !Lit[string](false).isLiteral(false) {
		t.Error("Lit(false).Is(false) must be true")
	}
}

func ExampleLit() {
	A := Lit[string](true)
	fmt.Println(A)

	B := Lit[string](false)
	fmt.Println(B)
	//Output:
	// Lit(true)
	// Lit(false)
}

func ExampleNot() {
	fmt.Println(Not(ID("A")))
	fmt.Println(Not(Lit[string](true)))
	fmt.Println(Not(Lit[string](false)))
	//Output:
	// Not("A")
	// Lit(false)
	// Lit(true)
}

func ExampleAnd() {
	A := ID("A")
	B := ID("B")
	C := ID("C")
	fmt.Println(And(A, Lit[string](true)))
	fmt.Println(And(A, Lit[string](false)))
	fmt.Println(And(A, B, C))
	fmt.Println(And(A, Not(B)))
	//Output:
	// "A"
	// Lit(false)
	// And("A", "B", "C")
	// And("A", Not("B"))
}

func ExampleOr() {

	A := ID("A")
	B := ID("B")

	fmt.Println(Or(A, Lit[string](true)))
	fmt.Println(Or(A, Lit[string](false)))
	fmt.Println(Or(A, Not(B)))
	//Output:
	// Lit(true)
	// "A"
	// Or("A", Not("B"))
}

func truthTester(t *testing.T, label string, z Expr[string], expected bool) {
	if !z.isLiteral(expected) {
		t.Errorf("%s: expected %v got %v", label, expected, z)
	}
}

func Test_truthTables(t *testing.T) {

	T := Lit[string](true)
	F := Lit[string](false)

	//Not
	truthTester(t, "Not(F)", Not(F), true)
	truthTester(t, "Not(T)", Not(T), false)

	//Or
	truthTester(t, "Or(F,F)", Or(F, F), false)
	truthTester(t, "Or(F,T)", Or(F, T), true)
	truthTester(t, "Or(T,F)", Or(T, F), true)
	truthTester(t, "Or(T,T)", Or(T, T), true)

	//And
	truthTester(t, "And(F,F)", And(F, F), false)
	truthTester(t, "And(F,T)", And(F, T), false)
	truthTester(t, "And(T,F)", And(T, F), false)
	truthTester(t, "And(T,T)", And(T, T), true)

}
