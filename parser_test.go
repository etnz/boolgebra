package boolgebra

import "testing"

func TestScan(t *testing.T) {
	data := []struct {
		src  string
		want []token
	}{
		{"", nil},
		{"abc", []token{{0, tkIdentifier, "abc"}}},
		{"a b", []token{{0, tkIdentifier, "a"}, {2, tkIdentifier, "b"}}},

		{"a and b", []token{{0, tkIdentifier, "a"}, {2, tkLongAnd, "and"}, {6, tkIdentifier, "b"}}},
		{"a & b", []token{{0, tkIdentifier, "a"}, {2, tkAnd, "&"}, {4, tkIdentifier, "b"}}},
		{"a | b", []token{{0, tkIdentifier, "a"}, {2, tkOr, "|"}, {4, tkIdentifier, "b"}}},
		{"a ^ b", []token{{0, tkIdentifier, "a"}, {2, tkXor, "^"}, {4, tkIdentifier, "b"}}},
		{"a not b", []token{{0, tkIdentifier, "a"}, {2, tkNot, "not"}, {6, tkIdentifier, "b"}}},
		{"not a b", []token{{0, tkNot, "not"}, {4, tkIdentifier, "a"}, {6, tkIdentifier, "b"}}},
		{"a != b", []token{{0, tkIdentifier, "a"}, {2, tkNeq, "!="}, {5, tkIdentifier, "b"}}},
		{"a => b", []token{{0, tkIdentifier, "a"}, {2, tkImpl, "=>"}, {5, tkIdentifier, "b"}}},
		{"a <=> b", []token{{0, tkIdentifier, "a"}, {2, tkEq, "<=>"}, {6, tkIdentifier, "b"}}},
		{"a () b", []token{{0, tkIdentifier, "a"}, {2, tkLParen, "("}, {3, tkRParen, ")"}, {5, tkIdentifier, "b"}}},
		//lit
		{"true", []token{{0, tkTrue, "true"}}},
		{"false", []token{{0, tkFalse, "false"}}},
	}

	for _, td := range data {
		t.Run(td.src, func(t *testing.T) {
			got := make([]token, 0, len(td.want))
			s := newParser([]byte(td.src))
			for {
				next := s.peek()
				tk := s.next()
				if tk.kind == tkEOF {
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
		{"not a & b", And(Not(a), b)},
		{"a & (b | c)", And(a, Or(b, c))},

		// command call (ID is the first builtin function)
		{"Reduce a => a", Lit(true)},
		{"Reduce (a => a)", Lit(true)},
		{"(Reduce a => a) & b", b},
		{"Reduce a => b and b => a", Eq(a,b) },
		

		{"Ascertain a | b", Lit(true)},
		{"Ascertain a&b | a&c", a},
		//{"Exactly 1 in  a&b , a&c", a},

		// lit
		{"true", Lit(true)},
		{"false", Lit(false)},
		{"a & false", And(a, Lit(false))},
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

func TestFormat(t *testing.T) {
	data := []struct {
		src string
	}{
		{"a b c"},
		{"a b c | c d e"},
		{"a & b"},
		{"a ^ b"},
		{"a | b"},
		{"a <=> b"},
		{"a != b"},
		{"a => b"},
		// Composing ops
		{"a & b ^ c"},
		{"a & b | c"},
		{"a & b <=> a & b"},
		{"a & b != a & b"},
		{"a & b => a & b"},
		{"a & (b | c)"},

		// not case
		{"a not b c"},
		{"a b not c"},
		{"not a b c"},
		{"not (a & b)"},

		// lit
		{"true"},
		{"false"},
		{"a & false"},
	}

	for _, td := range data {
		t.Run(td.src, func(t *testing.T) {

			want, err := Parse(td.src)
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}

			got, err := Parse(want.String())
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}
			if got.String() != want.String() {
				t.Errorf("String(%q) error: got %v want %v", td.src, got, want)
			}
		})
	}

}
