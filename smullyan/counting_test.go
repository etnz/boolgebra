package smullyan

import (
	"fmt"
	"testing"

	. "github.com/etnz/boolgebra"
)

func TestComplement(t *testing.T) {

	p := []int{0, 1, 4}
	exp := []int{2, 3, 5}
	c := complement(p, 6)
	if len(c) != len(exp) {
		t.Fatalf("invalid complement length got %v want %v", len(c), len(exp))
	}
	for i, v := range c {
		w := exp[i]
		if w != v {
			t.Errorf("invalid c[%d] got %v want %v", i, v, w)
		}
	}
}

func ExampleQuantified() {

	A := ID("A")
	B := ID("B")
	C := ID("C")
	fmt.Println(Exactly(2, A, B, C))
	fmt.Println(AtLeast(2, A, B, C))
	fmt.Println(AtMost(2, A, B, C))

	//Output:
	// Or(And("A", "B", Not("C")), And(Not("A"), "B", "C"), And("A", Not("B"), "C"))
	// Or(And("A", "B"), And("B", "C"), And("A", "C"))
	// Or(Not("C"), Not("A"), Not("B"))
}
