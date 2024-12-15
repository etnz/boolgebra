package boolgebra

import "testing"

func TestScan(t *testing.T) {
	data := []struct {
		src  string
		want []token
	}{
		{"", nil},
		{"abc", []token{{0, identifier, "abc"}}},
		{"a b", []token{{0, identifier, "a"}, {2, identifier, "b"}}},

		{"a and b", []token{{0, identifier, "a"}, {2, identifier, "and"}, {6, identifier, "b"}}},
		{"a & b", []token{{0, identifier, "a"}, {2, and, "&"}, {4, identifier, "b"}}},
		{"a | b", []token{{0, identifier, "a"}, {2, or, "|"}, {4, identifier, "b"}}},
		{"a ^ b", []token{{0, identifier, "a"}, {2, xor, "^"}, {4, identifier, "b"}}},
		{"a not b", []token{{0, identifier, "a"}, {2, not, "not"}, {6, identifier, "b"}}},
		{"not a b", []token{{0, not, "not"}, {4, identifier, "a"}, {6, identifier, "b"}}},
		{"a != b", []token{{0, identifier, "a"}, {2, neq, "!="}, {5, identifier, "b"}}},
		{"a => b", []token{{0, identifier, "a"}, {2, impl, "=>"}, {5, identifier, "b"}}},
		{"a <=> b", []token{{0, identifier, "a"}, {2, eq, "<=>"}, {6, identifier, "b"}}},
		{"a () b", []token{{0, identifier, "a"}, {2, lparen, "("}, {3, rparen, ")"}, {5, identifier, "b"}}},
	}

	for _, td := range data {
		t.Run(td.src, func(t *testing.T) {
			got := make([]token, 0, len(td.want))
			s := newParser([]byte(td.src))
			for {
				next := s.peek()
				tk := s.next()
				if tk.kind == eof {
					break
				}
				got = append(got, tk)
				if next != tk {
					t.Errorf("Unexpected peek function behavior got %v want %v", next, tk)
				}
			}

			// Compare results
			if len(td.want) != len(got) {
				t.Errorf("Error in scan(%q) size mismatch, got %v want %v", td.src, got, td.want)
				return
			}
			for i := range td.want {
				w, g := td.want[i], got[i]
				if w != g {
					t.Errorf("Error in scan(%q) content mismatch, got %v want %v", td.src, got, td.want)
					return
				}
			}
		})
	}
}

func TestParse(t *testing.T) {
	a := ID("a")
	b := ID("b")
	c := ID("c")
	data := []struct {
		src  string
		want Expr
	}{
		{"a b c", ID("a b c")},
		{"a b c | c d e", Or(ID("a b c"), ID("c d e"))},
		{"a & b", And(a, b)},
		{"a ^ b", Xor(a, b)},
		{"a | b", Or(a, b)},
		{"a <=> b", Eq(a, b)},
		{"a != b", Xor(a, b)},
		{"a => b", Impl(a, b)},
		// Composiing ops
		{"a & b ^ c", Xor(And(a, b), c)},
		{"a & b | c", Or(And(a, b), c)},
		{"a & b <=> a & b", Eq(And(a, b), And(a, b))},
		{"a & b != a & b", Xor(And(a, b), And(a, b))},
		{"a & b => a & b", Impl(And(a, b), And(a, b))},
		{"a & (b | c)", And(a, Or(b, c))},

		// not case
		{"a not b c", Not(ID("a b c"))},
		{"a b not c", Not(ID("a b c"))},
		{"not a b c", Not(ID("a b c"))},
		{"not (a & b)", Not(And(a, b))},

		//{"a (b c)", ID("a b c")},
	}

	for _, td := range data {
		t.Run(td.src, func(t *testing.T) {
			l := newParser([]byte(td.src))
			got, err := l.parse(0)
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}
			eq := Simplify(Xor(td.want, got))
			t.Logf("got  %v \nwant %v \nequ %v", got, td.want, eq)
			if !eq.Is(false) {
				t.Errorf("parse(%q) error: got %v want %v", td.src, got, td.want)
			}
		})
	}
}

func TestSmullyan(t *testing.T) {
	// A Knight is someone that always tells the truth.
	// A Knaves is someone that always tells a lie.

	// The hero meets three guys 'a','b', 'c', each one can be a Knight or a Knaves.
	// Our hero asked 'a': "are 'b' and 'c' knights ?"
	// but 'a' also said that 'b' was a Knaves:

	pb, err := Parse(`
	  ( (b is knight & c is knight) <=> a is knight)
	& ( a is knight <=> b is not knight )
	`)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	res := Simplify(pb)
	want, _ := Parse("a is not knight & b is knight & c is not knight")

	cmp := Simplify(Xor(res, want))
	if !cmp.Is(false) {
		t.Errorf("invalid solution got %v want %v equ %v", res, want, cmp)
	}

	// Another example
	// https://en.wikipedia.org/wiki/Knights_and_Knaves#Examples
	// A says, "We are both knaves."

	pb, err = Parse("(a is not knight & b is not knight) <=> a is knight")
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	res = Simplify(pb)
	want, _ = Parse("a is not knight & b is knight")

	cmp = Simplify(Xor(res, want))
	if !cmp.Is(false) {
		t.Errorf("invalid solution got %v want %v equ %v", res, want, cmp)
	}
}
