package boolgebra

import "testing"

func TestAtLeast(t *testing.T) {
	a := ID("a")
	b := ID("b")
	c := ID("c")
	d := ID("d")

	data := []struct {
		n    int
		want string
	}{
		{0, `true`},
		{1, `a | b | c | d`},
		{2, `a & b | b & c | a & c | c & d | b & d | a & d`},
		{3, `a & b & c | a & c & d | b & c & d | a & b & d`},
		{4, `a & b & c & d`},
		{5, `false`},
	}

	for _, td := range data {

		got := AtLeast(td.n, a, b, c, d)
		want, _ := Parse(td.want)

		t.Logf("AtLeast(%v, a,b,c,d) = \n%v want \n%v", td.n, got, want)
		if !Simplify(Xor(got, want)).Is(false) {
			t.Fail()
		}

	}
}

func TestAtMost(t *testing.T) {
	a := ID("a")
	b := ID("b")
	c := ID("c")
	d := ID("d")

	data := []struct {
		n    int
		want string
	}{
		{0, `not a & not b & not c & not d`},
		{1, `not a & not b & not c | not a & not c & not d | not b & not c & not d | not a & not b & not d`},
		{2, `not a & not b | not b & not c | not a & not c | not c & not d | not b & not d | not a & not d`},
		{3, `not a | not b | not c | not d`},
		{4, `true`},
		{5, `true`},
	}

	for _, td := range data {

		got := AtMost(td.n, a, b, c, d)
		want, _ := Parse(td.want)

		t.Logf("AtMost(%v, a,b,c,d) = \n%v want \n%v", td.n, got, want)
		if !Simplify(Xor(got, want)).Is(false) {
			t.Fail()
		}

	}
}
