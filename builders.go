package boolgebra

// TermBuilder can be used to efficiently append ID ( or Not(ID)) into a big And
type TermBuilder[T comparable] struct {
	isLitFalse bool
	m          Term[T]
}

// And append a variable to the current term
func (t *TermBuilder[T]) And(id T, val bool) {
	if t.isLitFalse {
		return
	} // nothing to do
	if t.m == nil {
		t.m = make(Term[T])
	}
	if prev, exists := t.m[id]; exists && prev != val {
		//attempt to do something like  AND(x, !x) which is always false therefore the result will always be Lit(false)
		t.isLitFalse = true
		t.m = nil // destroy every previous values
		return
	}
	t.m[id] = val
}

// IsFalse returns true if the term under construction is already degenerated to False
func (t TermBuilder[T]) IsFalse() bool { return t.isLitFalse }

func (t *TermBuilder[T]) Build() Expr[T] {
	if t.isLitFalse {
		t.isLitFalse = false // reset it
		return Lit[T](false)
	}
	res := t.m
	t.m = nil // destroy reference to m to avoid editing it anymore
	return Expr[T]{res}
}
