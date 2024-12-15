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

// tokenKind identifies any tokenKind in the boolgebra language.
type tokenKind int

const (
	illegal tokenKind = iota
	eof               // end of file
	and               // '&'
	or                // '|'
	xor               // '^'
	not               // 'not'
	eq                // '<=>' logically Equivalent
	neq               // '!=' Not Equal
	impl              // '=>' Imply, or if..then statement
	lparen            // '('
	rparen            // ')'
	identifier
	litFalse
	litTrue
	lastKind // must always be the latest kind
)

// tokens maps tokenKind to a struct with various mapping information.
var tokens [lastKind]struct {
	prec int
	name string
} = [...]struct {
	prec int
	name string
}{
	//token		prec, name
	illegal:    {0, "illegal"},
	eof:        {0, "eof"},
	lparen:     {1, "lparen"},
	rparen:     {1, "rparen"},
	eq:         {5, "eq"},
	neq:        {5, "neq"},
	impl:       {5, "impl"},
	or:         {6, "or"},
	xor:        {7, "xor"},
	and:        {8, "and"},
	not:        {9, "not"},
	identifier: {11, "identifier"},
	litFalse:   {12, "false"},
	litTrue:    {12, "true"},
}

// String returns the tokenKind's name
func (t tokenKind) String() string { return tokens[t].name }

// token is what the tokenizer has read: the pos in the source, the tokenKind and its literal value.
type token struct {
	pos  int
	kind tokenKind
	lit  string
}

// String returns a debug view for token.
func (t token) String() string {
	return fmt.Sprintf("%s(%v) %q", t.kind, t.pos, t.lit)
}

// precedence returns the token precedence for this token.
func (t token) precedence() int { return tokens[t.kind].prec }

// basic unicode functions. We know they can be optimized, and to avoid depending on unicode package outside of this functions.

func isLetter(ch rune) bool         { return unicode.IsLetter(ch) }
func isDigit(ch rune) bool          { return unicode.IsDigit(ch) }
func isWhitespace(ch rune) bool     { return unicode.IsSpace(ch) }
func isIdentifierChar(ch rune) bool { return isLetter(ch) || isDigit(ch) }

// A parser holds the Pratt parser&parser internal state.
type parser struct {
	source []byte // source.

	ch       rune // current character.
	beg, end int  // current offsets in source.

}

func newParser(src []byte) *parser {
	p := &parser{source: src}
	p.nextChar()
	return p
}

// nextChar read the nextChar char in source.
func (p *parser) nextChar() {
	if p.end >= len(p.source) {
		p.ch = rune(-1) // by convention, eof is represented by run(-1)
		return
	}
	// read next unicode rune in the source.
	r, w := utf8.DecodeRune(p.source[p.end:])
	p.ch, p.beg, p.end = r, p.end, p.end+w
}

// skipf is a utility function that consume characters as long as they match 'f'.
func (p *parser) skipf(f func(ch rune) bool) {
	for f(p.ch) {
		p.nextChar()
	}
}

// scanf is a utility function that consume chararacters
func (p *parser) scanf(f func(ch rune) bool) (lit string) {
	offset, last := p.beg, p.end
	for f(p.ch) {
		last = p.end
		p.nextChar()
	}
	return string(p.source[offset:last])
}

// next returns the next token, and move the scanner to the next token.
func (p *parser) next() token {
	p.skipf(isWhitespace)
	pos, ch := p.beg, p.ch

	switch {
	case ch == rune(-1):
		return token{pos, eof, ""}
	case isLetter(ch):
		lit := p.scanf(isIdentifierChar)
		// before considering it an identifier check against reserved words
		switch lit {
		case "not":
			return token{pos, not, lit}
		case "true":
			return token{pos, litTrue, lit}
		case "false":
			return token{pos, litFalse, lit}
		}
		// anything else is an identifier.
		return token{pos, identifier, lit}
	default:
		// all cases based on ch value
		switch ch {
		case '(':
			p.nextChar()
			return token{pos, lparen, "("}
		case ')':
			p.nextChar()
			return token{pos, rparen, ")"}
		case '&':
			p.nextChar()
			return token{pos, and, "&"}
		case '|':
			p.nextChar()
			return token{pos, or, "|"}
		case '^':
			p.nextChar()
			return token{pos, xor, "^"}

		// Dual char tokens.
		case '!': // Read the "!=" token
			pos := p.beg
			p.nextChar()
			if p.ch == '=' {
				p.nextChar()
				return token{pos, neq, "!="}
			} else {
				return token{pos, illegal, fmt.Sprintf("expecting '=' found %q", p.ch)}
			}
		case '=': // Read the "=>" token
			pos := p.beg
			p.nextChar()
			if p.ch == '>' {
				p.nextChar()
				return token{pos, impl, "=>"}
			} else {
				return token{pos, illegal, fmt.Sprintf("expecting '>' found %q", p.ch)}
			}
		// 3-char tokens.
		case '<': // Read the "<=>" token
			pos := p.beg
			p.nextChar()
			if p.ch == '=' {
				p.nextChar()
				if p.ch == '>' {
					p.nextChar()
					return token{pos, eq, "<=>"}
				} else {
					return token{pos, illegal, fmt.Sprintf("expecting '>' found %q", p.ch)}
				}
			} else {
				return token{pos, illegal, fmt.Sprintf("expecting '=' found %q", p.ch)}
			}

		}
		return token{pos, illegal, fmt.Sprintf("unknown character %q", p.ch)}
	}
}

// peek returns the next token but doesn't consume it.
func (p *parser) peek() (tk token) {
	// Record the lexer state.
	ch, off, last := p.ch, p.beg, p.end
	tk = p.next()
	// Restore the lexer state.
	p.ch, p.beg, p.end = ch, off, last
	return tk
}

// parse a stream of tokens from a given sub expression precedence.
// To start parsing from scrach, rbp should be 0.
func (p *parser) parse(rbp int) (Expr, error) {
	//Start with the head recursion function akak nul denotation.
	left, err := p.head(p.next()) // create the subexpression
	if err != nil {
		return nil, err
	}

	// While next token precedence is greater than the subexpression one, tail recurse.
	for p.peek().precedence() > rbp {
		left, err = p.tail(left, p.next()) // tail recurse into the expression.
		if err != nil {
			return nil, err
		}
	}
	return left, nil
}

// head starts recursion based on 'tk'.
func (l *parser) head(tk token) (Expr, error) {
	switch tk.kind {
	case litTrue:
		return Lit(true), nil
	case litFalse:
		return Lit(false), nil

	case identifier:
		return ID(tk.lit), nil
	case not:
		expr, err := l.parse(tk.precedence())
		if err != nil {
			return nil, err
		}
		return Not(expr), nil

	case lparen:
		expr, err := l.parse(tk.precedence())
		if err != nil {
			return nil, err
		}
		//and consume the rparen
		if r := l.next(); r.kind != rparen {
			return nil, fmt.Errorf("invalid token %s expecting %q", r, rparen)
		}
		return expr, nil
	default:
		return nil, fmt.Errorf("invalid token %s in this position", tk)
	}
}

// readID is a utility to cast left as an ID with it's current value.
func readID(left Expr) (id string, val, ok bool) {
	var m minterm
	m, ok = left.(minterm)
	if !ok {
		return "", false, false
	}
	for k, val := range m {
		return k, val, true
	}
	return "", false, false

}

// tail continues the recursion for an existing expression.
func (p *parser) tail(left Expr, tk token) (Expr, error) {
	switch tk.kind {

	case identifier:
		// left MUST have been an ID
		id, val, ok := readID(left)
		if !ok {
			return nil, fmt.Errorf("invalid token %s expected an ID expression got %v", tk, left)
		}
		// preserve the current val of this ID. so that any 'not' token in the sentence of identifier switches the value.
		// TODO: decide whether or not we should have a fixed size of identifiers to make an ID.

		// TODO: we are formating ID as a space separated list of identifier. This should be better formalized.
		return minterm{id + " " + tk.lit: val}, nil

	case not:
		// left MUST have been an ID
		id, val, ok := readID(left)
		if !ok {
			return nil, fmt.Errorf("invalid %s expected an ID expression got %v", tk, left)
		}
		return minterm{id: !val}, nil

	case or, and, xor, eq, neq, impl:
		right, err := p.parse(tk.precedence())
		if err != nil {
			return nil, err
		}
		var res Expr
		switch tk.kind {
		case or:
			res = Or(left, right)
		case and:
			res = And(left, right)
		case xor:
			res = Xor(left, right)
		case eq:
			res = Eq(left, right)
		case neq:
			res = Neq(left, right)
		case impl:
			res = Impl(left, right)
		default:
			panic("unexpected kind")
		}
		return res, nil
	default:
		return nil, fmt.Errorf("invalid token %s after %v", tk, left)
	}
}

// Parse a boolean expression in src, and returns its Expr.
//
// Literal: `true` or `false` are the two possible literal.
//
// Variable: any usual identifier `<letter> (digit|letter)*`
//
// an ID is made of one or more Variable.
//
// 'not' can be inserted in a list of variable to form the negation of an ID:
// e.g. `John is not Knight` is equivalent to Not(ID("John is Knight")).
// Note that the 'not' can be placed anywhere in the ID, every occurrence toggles the boolean value. So that
// e.g. in `not John is not Knight` the 'not' effect are canceled.
// 'not' can be used as a prefix operator `not (a&b)` is equivalent to `Not(And(a,b))`
//
// The following binary operators in increasing binding power (precedence)
//
//	'<=>' logically equivalent
//	'!='  not logically equivalent
//	'=>' imply.
//	'|' or
//	'^' xor
//	'&' and
//
// So that `a | b & c` is equivalent to `a | (b&c)`
//
// Note as a mnemonic that the shorter the operator the highest precedence it has. For that matter, '!=' and '^' are
// the same operation, just at different priority level.
func Parse(src string) (Expr, error) {
	return newParser([]byte(src)).parse(0)
}
