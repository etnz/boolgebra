package boole

import (
	"fmt"
	"testing"
)

func ExampleID() {
	A := ID("A")
	fmt.Println(A)
	//Output: "A"
}

func ExampleLit() {
	A := Lit(true)
	fmt.Println(A)

	B := Lit(false)
	fmt.Println(B)
	//Output: True
	// False
}

func ExampleNot() {
	// despite not being actually simplified, expression can only be a sum of prod.
	// and literal are represented as False = empty sum, true = sum of an empty prod
	fmt.Println(Not(ID("A")))
	fmt.Println(Not(Lit(true)))
	fmt.Println(Not(Lit(false)))
	//Output:
	// !"A"
	// False
	// True
}

func ExampleAnd() {
	A := ID("A")
	B := ID("B")
	fmt.Println(And(A, Lit(true)))
	fmt.Println(And(A, Lit(false)))
	fmt.Println(And(A, Not(B)))
	//Output:
	// "A"
	// False
	// "A" & !"B"
}

func ExampleOr() {

	A := ID("A")
	B := ID("B")

	fmt.Println(Or(A, Lit(true)))
	fmt.Println(Or(A, Lit(false)))
	fmt.Println(Or(A, Not(B)))
	//Output:
	// True
	// "A"
	// "A" | !"B"
}

// TestMinterm_PosLen make sure that we count the number of true correctly
func TestMinterm_PosLen(t *testing.T) {
	x := minterm{"A": true, "B": true, "C": false, "E": true}
	if x.PosLen() != 3 {
		t.Errorf("invalid minterm %v PosLen attribute got %v want 3", x, x.PosLen())
	}
}

// TestMinterm_combine gold test some minterm combinations
func TestCombine(t *testing.T) {

	x := minterm{"A": true, "B": true, "C": false, "E": true}
	var y, r, c minterm
	var ok bool

	// 1,0 -> _
	y = minterm{"A": true, "B": true, "C": false, "E": false}
	r = minterm{"A": true, "B": true, "C": false}
	c, ok = combine(x, y)
	if !ok {
		t.Errorf("failed to combine(%v,%v)", x, y)
	} else if !c.Equals(r) {
		t.Errorf("combine(%v,%v) got %v expected %v", x, y, c, r)
	}

	// 1,_ -> _
	y = minterm{"A": true, "B": true, "C": false}
	r = minterm{"A": true, "B": true, "C": false}
	c, ok = combine(x, y)
	if !ok {
		t.Errorf("failed to combine(%v,%v)", x, y)
	} else if !c.Equals(r) {
		t.Errorf("combine(%v,%v) got %v expected %v", x, y, c, r)
	}

	// 0,1 -> _
	y = minterm{"A": true, "B": true, "C": true, "E": true}
	r = minterm{"A": true, "B": true, "E": true}
	c, ok = combine(x, y)
	if !ok {
		t.Errorf("failed to combine(%v,%v)", x, y)
	} else if !c.Equals(r) {
		t.Errorf("combine(%v,%v) got %v expected %v", x, y, c, r)
	}

	// 0,_ -> _
	y = minterm{"A": true, "B": true, "E": true}
	r = minterm{"A": true, "B": true, "E": true}
	c, ok = combine(x, y)
	if !ok {
		t.Errorf("failed to combine(%v,%v)", x, y)
	} else if !c.Equals(r) {
		t.Errorf("combine(%v,%v) got %v expected %v", x, y, c, r)
	}
}
