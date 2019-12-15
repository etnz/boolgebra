package boole

import "testing"

//TestReduce check that we get the prime correctly using the wikipedia examples
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
	x := expression{m4, m8, m9, m10, m12, m11, m14, m15}
	//primes are
	//m4_12 := m("bc'd'")
	//m8_9_10_11 := m("ab'")
	//m8_10_12_14 := m("ad'")
	//m10_11_14_15 := m("ac")
	y := reduce(x)
	t.Logf("reduced = %v", y)

}

//newminterm creates a new minterm using ' at the end of the string to find out that its a neg
func m(x string) minterm {
	res := make(minterm)

	for i, id := range x {
		nextisquote := i+1 < len(x) && x[i+1] == '\''
		if id != '\'' {
			res[string(id)] = !nextisquote
		}
	}
	return res
}
