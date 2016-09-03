package boolgebra

import "testing"

// a few globals to simplify writting tests
var A = ID("A")
var B = ID("B")
var C = ID("C")
var D = ID("D")
var p = ID("p")
var q = ID("q")
var r = ID("r")

// var s = ID("s")

// list of expression, that:
//  have no repetition
//  are sorted in natural order
var EqualsCases = []Expression{
	// check trvial case (who gets the trivial cases wrong ?)
	True, False,
	A,
	Not(True),
	Not(False),
	Not(A),
	And(A, B),
	And(Not(A), B),
	And(Not(True), B),
	And(True, B),
	Or(A, B),
	Or(Not(A), B),
	Or(Not(True), B),
	Or(True, B),
	Not(Not(A)),
	Not(And(A, B)),
	Not(And(Not(A), B)),
	Not(And(Not(True), B)),
	Not(And(True, B)),
	Not(Or(A, B)),
	Not(Or(Not(A), B)),
	Not(Or(Not(True), B)),
	Not(Or(True, B)),
}

// TestEquality run a test case to make sure they are equals
func TestEquality(t *testing.T) {

	for i, testcase := range EqualsCases {
		x := testcase
		if !Equals(x, x) {
			t.Errorf("Error EqualsCases[%d]: %v  not equals to self", i, x)
		}
	}
	// the test case is so that there is no repetition
	for i, x := range EqualsCases {
		for j, y := range EqualsCases {
			// i == j <=> x == y
			if i == j {
				if !Equals(x, y) {
					t.Errorf("Error EqualsCases[%d,%d]: they should be equals:\n%v\n%v   ", i, j, x, y)
				}
			} else {
				if Equals(x, y) {
					t.Errorf("Error EqualsCases[%d,%d]: they should not be equals:\n%v\n%v   ", i, j, x, y)
				}
			}
		}
	}
}
