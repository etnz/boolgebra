package boolgebra

import "testing"

// TestReduce check that we get the prime correctly using the wikipedia examples
func TestReduce(t *testing.T) {
	/*
		Number
		of 1s	Minterm	Binary
		Representation
		1	m4	0100
			m8	1000
		2	(m9)	1001
			m10	1010
			m12	1100
		3	m11	1011
			(m14)	1110
		4	m15	1111


	*/

	m4 := m("a'bc'd'")
	m8 := m("ab'c'd'")
	m9 := m("ab'c'd")
	m10 := m("ab'cd'")
	m12 := m("abc'd'")
	m11 := m("ab'cd")
	m14 := m("abcd'")
	m15 := m("abcd")
	x := Expr[string]{m4, m8, m9, m10, m12, m11, m14, m15}
	//primes are
	//m4_12 := m("bc'd'")
	//m8_9_10_11 := m("ab'")
	//m8_10_12_14 := m("ad'")
	//m10_11_14_15 := m("ac")
	y := Simplify(x)
	t.Logf("reduced = %v", y)

}

// newminterm creates a new minterm using ' at the end of the string to find out that its a neg
func m(x string) Term[string] {
	res := make(Term[string])

	for i, id := range x {
		nextisquote := i+1 < len(x) && x[i+1] == '\''
		if id != '\'' {
			res[string(id)] = !nextisquote
		}
	}
	return res
}

// TestPosLen make sure that we count the number of true correctly
func TestPositives(t *testing.T) {
	x := Term[string]{"A": true, "B": true, "C": false, "E": true}
	if positives(x) != 3 {
		t.Errorf("invalid minterm %v PosLen attribute got %v want 3", x, positives(x))
	}
}

// TestMinterm_combine gold test some minterm combinations
func TestCombine(t *testing.T) {

	x := Term[string]{"A": true, "B": true, "C": false, "E": true}
	var y, r, c Term[string]
	var ok bool

	// 1,0 -> _
	y = Term[string]{"A": true, "B": true, "C": false, "E": false}
	r = Term[string]{"A": true, "B": true, "C": false}
	c, ok = combine(x, y)
	if !ok {
		t.Errorf("failed to combine(%v,%v)", x, y)
	} else if !equals(c, r) {
		t.Errorf("combine(%v,%v) got %v expected %v", x, y, c, r)
	}

	// 1,_ -> _
	y = Term[string]{"A": true, "B": true, "C": false}
	r = Term[string]{"A": true, "B": true, "C": false}
	c, ok = combine(x, y)
	if !ok {
		t.Errorf("failed to combine(%v,%v)", x, y)
	} else if !equals(c, r) {
		t.Errorf("combine(%v,%v) got %v expected %v", x, y, c, r)
	}

	// 0,1 -> _
	y = Term[string]{"A": true, "B": true, "C": true, "E": true}
	r = Term[string]{"A": true, "B": true, "E": true}
	c, ok = combine(x, y)
	if !ok {
		t.Errorf("failed to combine(%v,%v)", x, y)
	} else if !equals(c, r) {
		t.Errorf("combine(%v,%v) got %v expected %v", x, y, c, r)
	}

	// 0,_ -> _
	y = Term[string]{"A": true, "B": true, "E": true}
	r = Term[string]{"A": true, "B": true, "E": true}
	c, ok = combine(x, y)
	if !ok {
		t.Errorf("failed to combine(%v,%v)", x, y)
	} else if !equals(c, r) {
		t.Errorf("combine(%v,%v) got %v expected %v", x, y, c, r)
	}
}
