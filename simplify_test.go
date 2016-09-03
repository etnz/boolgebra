package boolgebra

import "testing"

// list of test cases:
// first value is the source, second is the expected simplified expression
var SimplifyCases = [][2]Expression{
	// let's start with trivialities
	{True, True},
	{False, False},
	{A, A},
	{Not(A), Not(A)},
	{And(A, B), And(A, B)},
	{Or(A, B), Or(A, B)},

	// go for simple literal simplifications
	{Not(True), False},
	{And(True, A), A},
	{And(A, True), A},
	{And(False, A), False},
	{And(A, False), False},
	{Or(True, A), True},
	{Or(A, True), True},
	{Or(False, A), A},
	{Or(A, False), A},

	// some repetitions
	{Or(A, B, A), Or(A, B)},
	{Or(A, A, A, B), Or(A, B)},
	{Or(A, B, B), Or(A, B)},
	{Or(A, B, B, B), Or(A, B)},
	{Or(A, A, A), A},
	// And is not a special case, but I need to repeat it
	{And(A, B, A), And(A, B)},
	{And(A, A, A, B), And(A, B)},
	{And(A, B, B), And(A, B)},
	{And(A, B, B, B), And(A, B)},
	{And(A, A, A), A},

	// some buggy cases
	{Or(And(A, B), And(A, Not(B))), A},
	{Or(And(Not(B), Not(A)), A, B), True},
	// {And(Or(And(Not(A), Not(And(B, C))), And(A, And(B, C))), Or(And(Not(A), Not(Not(B))), And(A, Not(B)))), And(B, Not(A), Not(C))},
	// {And(Or(Not(A), B), Or(Not(A), C), Or(Not(B), Not(C), A), True, Or(B, A)), And(B, Not(A), Not(C))},
	// {And(Or(Not(A), B), Or(Not(A), C), Or(Not(B), Not(C), A), Or(B, A)), And(B, Not(A), Not(C))},
	// {Or(And(Not(A), B), And(Not(A), C), And(Not(B), Not(C), A), And(B, A)), Or(B, Not(A), Not(C))},

	{Impl(Or(p, And(q, r)), And(Or(p, q), Or(p, r))), True},
}

// TestEquality run a test case to make sure they are equals
func TestSimplify(t *testing.T) {

	for i, testcase := range SimplifyCases {
		x, res := testcase[0], testcase[1]
		y := Simplify(x)
		t.Logf("SimplifyCases[%d]: Simplify(%v) = %v", i, x, y)
		if !Equals(y, res) {
			t.Errorf("Error SimplifyCases[%d]: Simplify(%v) <> expected:\n%v\n%v", i, x, y, res)
		}
	}

	for i, testcase := range SimplifyCases {
		x, res := testcase[0], testcase[1]
		y := Simplify(expand(Eq(x, res)))
		t.Logf("SimplifyCases[%d]: Simplify(Eq(%v,%v)) = %v", i, x, res, y)
		if !y.Val() {
			t.Errorf("Error SimplifyCases[%d]: Simplify(Eq(%v,%v)) <> True:\n%v", i, x, res, y)
		}
	}

}

func TestBug(t *testing.T) {
	i := 26
	testcase := SimplifyCases[i]
	x, res := testcase[0], testcase[1]
	pb := x
	selems := SimplifyAll(pb.Elems()...)
	for i, p := range pb.Elems() {
		t.Logf("Simplify(%v) -> %v", p, selems[i])
	}

	y := Simplify(pb)
	t.Logf("SimplifyCases[%d]: \nSimplify(Eq(%v,%v)) = \nSimplify(%v) =  \n%v", i, x, res, pb, y)
	if !y.Val() {
		t.Errorf("Error SimplifyCases[%d]: Simplify(Eq(%v,%v)) <> True:\n%v", i, x, res, y)
	}
}
