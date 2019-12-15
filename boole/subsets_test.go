package boole

import (
	"testing"
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

func TestSubsetNext(t *testing.T) {
	p := []int{0, 0}
	var res []int // cumulate all slices here
	for nextsubset(p, 4) {
		res = append(res, p...)
	}
	t.Logf("last p = %v", p)

	exp := []int{0, 1, 1, 2, 0, 2, 2, 3, 1, 3, 0, 3}
	if len(exp) != len(res) {
		t.Errorf("invalid subsets length got %v want %v", len(exp), len(res))
	}
	for i, x := range exp {
		if res[i] != x {
			t.Errorf("invalid res got res[%d]=%v want %v", i, res[i], x)
		}
	}

}
