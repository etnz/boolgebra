package boolgebra

import "testing"

type tokenized struct {
	pos int
	t   token
	lit string
}

func TestScan(t *testing.T) {
	data := []struct {
		src  string
		want []tokenized
	}{
		{"", nil},
		{"abc", []tokenized{{0, identifier, "abc"}}},
		{"a b", []tokenized{{0, identifier, "a"}, {2, identifier, "b"}}},

		{"a and b", []tokenized{{0, identifier, "a"}, {2, identifier, "and"}, {6, identifier, "b"}}},
		{"a And b", []tokenized{{0, identifier, "a"}, {2, and, "And"}, {6, identifier, "b"}}},
		{"a Or b", []tokenized{{0, identifier, "a"}, {2, or, "Or"}, {5, identifier, "b"}}},
		{"a Xor b", []tokenized{{0, identifier, "a"}, {2, xor, "Xor"}, {6, identifier, "b"}}},
		{"a not b", []tokenized{{0, identifier, "a"}, {2, not, "not"}, {6, identifier, "b"}}},
		{"a Not b", []tokenized{{0, identifier, "a"}, {2, not, "Not"}, {6, identifier, "b"}}},
		{"a != b", []tokenized{{0, identifier, "a"}, {2, neq, "!="}, {5, identifier, "b"}}},
		{"a => b", []tokenized{{0, identifier, "a"}, {2, impl, "=>"}, {5, identifier, "b"}}},
		{"a <=> b", []tokenized{{0, identifier, "a"}, {2, eq, "<=>"}, {6, identifier, "b"}}},
		{"a () b", []tokenized{{0, identifier, "a"}, {2, lparen, "("}, {3, rparen, ")"}, {5, identifier, "b"}}},
		{"a b.", []tokenized{{0, identifier, "a"}, {2, identifier, "b"}, {3, dot, "."}}},
	}

	for _, td := range data {
		t.Run(td.src, func(t *testing.T) {
			got := make([]tokenized, 0, len(td.want))
			s := newScanner([]byte(td.src))
			for {
				pos, k, lit := s.scan()
				if k == eof {
					break
				}
				got = append(got, tokenized{pos, k, lit})
			}

			// Compare results
			if len(td.want) != len(got) {
				t.Errorf("Error in scan(%q) size mismatch, got %v want %v", td.src, got, td.want)
				return
			}
			for i := range td.want {
				w, g := td.want[i], got[i]
				if w.pos != g.pos || w.lit != g.lit || w.t != g.t {
					t.Errorf("Error in scan(%q) content mismatch, got %v want %v", td.src, got, td.want)
					return
				}
			}
		})
	}
}
