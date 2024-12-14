package boolgebra

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// This file contains the Pratt parser for boolgebra language.
//
// Pratt parsing was first described by Vaughan Pratt in the 1973 paper "Top Down Operator Precedence"
// Pratt, Vaughan. "Top Down Operator Precedence." Proceedings of the 1st Annual ACM SIGACT-SIGPLAN Symposium on Principles of Programming Languages (1973).
// https://web.archive.org/web/20151223215421/http://hall.org.ua/halls/wizzard/pdf/Vaughan.Pratt.TDOP.pdf
//
// The Pratt parsers starts with a scanner/lexer that detects tokens

// token identifies any token in the boolgebra language.
type token int

const (
	illegal token = iota
	eof           // end of file
	and           // 'And'
	or            // 'Or'
	xor           // 'Xor'
	not           // 'not' or 'Not'
	eq            // '<=>' logically Equivalent
	neq           // '!=' Not Equal
	impl          // '=>' Imply, or if..then statement
	lparen        // '('
	rparen        // ')'
	dot           // '.' closing a triple
	identifier
)

// A scanner holds the scanner internal state.
type scanner struct {
	source []byte // source.

	ch           rune // current character.
	offset, last int  // current offset in source.
}

func newScanner(src []byte) *scanner {
	s := &scanner{source: src}
	s.next()
	return s
}

// next read the next char in source.
func (s *scanner) next() {
	if s.last >= len(s.source) {
		s.ch = rune(-1)
		return
	}
	r, w := utf8.DecodeRune(s.source[s.last:])
	s.ch, s.offset, s.last = r, s.last, s.last+w
}

// basic unicode functions. We know they can be optimized, and to avoid depending on unicode package outside of this functions.

func lower(ch rune) rune            { return unicode.ToLower(ch) }
func isLetter(ch rune) bool         { return unicode.IsLetter(ch) }
func isDigit(ch rune) bool          { return unicode.IsDigit(ch) }
func isWhitespace(ch rune) bool     { return unicode.IsSpace(ch) }
func isIdentifierChar(ch rune) bool { return isLetter(ch) || isDigit(ch) }

func (s *scanner) skipf(f func(ch rune) bool) {
	for f(s.ch) {
		s.next()
	}
}

func (s *scanner) scanf(f func(ch rune) bool) (lit string) {
	offset, last := s.offset, s.last
	for f(s.ch) {
		last = s.last
		s.next()

	}
	return string(s.source[offset:last])
}

func (s *scanner) scan() (pos int, tok token, lit string) {
	s.skipf(isWhitespace)
	pos, ch := s.offset, s.ch

	switch {
	case ch == rune(-1):
		return pos, eof, ""
	case isLetter(ch):
		lit := s.scanf(isIdentifierChar)
		// before considering it an identifier check against reserved words
		switch lit {
		case "And":
			return pos, and, lit
		case "Or":
			return pos, or, lit
		case "Xor":
			return pos, xor, lit
		case "not", "Not": // not has two case form
			return pos, not, lit
		}
		// anything else is an identifier.
		return pos, identifier, lit
	default:
		// all cases based on ch value
		switch ch {
		case '(':
			s.next()
			return pos, lparen, "("
		case ')':
			s.next()
			return pos, rparen, ")"
		case '.':
			s.next()
			return pos, dot, "."

		// Dual char tokens.
		case '!': // Read the "!=" token
			pos := s.offset
			s.next()
			if s.ch == '=' {
				s.next()
				return pos, neq, "!="
			} else {
				return pos, illegal, fmt.Sprintf("expecting '=' found %q", s.ch)
			}
		case '=': // Read the "=>" token
			pos := s.offset
			s.next()
			if s.ch == '>' {
				s.next()
				return pos, impl, "=>"
			} else {
				return pos, illegal, fmt.Sprintf("expecting '>' found %q", s.ch)
			}
		// 3-char tokens.
		case '<': // Read the "<=>" token
			pos := s.offset
			s.next()
			if s.ch == '=' {
				s.next()
				if s.ch == '>' {
					s.next()
					return pos, eq, "<=>"
				} else {
					return pos, illegal, fmt.Sprintf("expecting '>' found %q", s.ch)
				}
			} else {
				return pos, illegal, fmt.Sprintf("expecting '=' found %q", s.ch)
			}

		}
		return
	}
}
