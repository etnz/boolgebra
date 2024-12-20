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
	tkIllegal    tokenKind = iota
	tkEOF                  // end of file
	tkAnd                  // '&'
	tkLongAnd              // 'and'
	tkOr                   // '|'
	tkXor                  // '^'
	tkNot                  // 'not'
	tkEq                   // '<=>' logically Equivalent
	tkNeq                  // '!=' Not Equal
	tkImpl                 // '=>' Imply, or if..then statement
	tkLParen               // '('
	tkRParen               // ')'
	tkIdentifier           //  letter, (letter | digit)*
	tkFalse                // 'false'
	tkTrue                 // 'true'
	tkReduce               // 'Reduce'
	tkAscertain            // 'Ascertain'

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
	tkIllegal:    {0, "illegal"},
	tkEOF:        {0, "eof"},
	tkLParen:     {1, "lparen"},
	tkRParen:     {1, "rparen"},
	tkReduce:     {2, "Reduce"},
	tkAscertain:  {2, "Ascertain"},
	tkLongAnd:    {3, "and"},
	tkEq:         {5, "eq"},
	tkNeq:        {5, "neq"},
	tkImpl:       {5, "impl"},
	tkOr:         {6, "or"},
	tkXor:        {7, "xor"},
	tkAnd:        {8, "&"},
	tkNot:        {9, "not"},
	tkIdentifier: {11, "identifier"},
	tkFalse:      {12, "false"},
	tkTrue:       {12, "true"},
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
	return fmt.Sprintf("%s:%v %q", t.kind, t.pos, t.lit)
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
		return token{pos, tkEOF, ""}
	case isLetter(ch):
		lit := p.scanf(isIdentifierChar)
		// before considering it an identifier check against reserved words
		switch lit {
		case tkReduce.String():
			return token{pos, tkReduce, lit}
		case tkAscertain.String():
			return token{pos, tkAscertain, lit}
		case tkLongAnd.String():
			return token{pos, tkLongAnd, lit}
		case tkNot.String():
			return token{pos, tkNot, lit}
		case tkTrue.String():
			return token{pos, tkTrue, lit}
		case tkFalse.String():
			return token{pos, tkFalse, lit}
		}
		// anything else is an identifier.
		return token{pos, tkIdentifier, lit}
	default:
		// all cases based on ch value
		switch ch {
		case '(':
			p.nextChar()
			return token{pos, tkLParen, "("}
		case ')':
			p.nextChar()
			return token{pos, tkRParen, ")"}
		case '&':
			p.nextChar()
			return token{pos, tkAnd, "&"}
		case '|':
			p.nextChar()
			return token{pos, tkOr, "|"}
		case '^':
			p.nextChar()
			return token{pos, tkXor, "^"}

		// Dual char tokens.
		case '!': // Read the "!=" token
			pos := p.beg
			p.nextChar()
			if p.ch == '=' {
				p.nextChar()
				return token{pos, tkNeq, "!="}
			} else {
				return token{pos, tkIllegal, fmt.Sprintf("expecting '=' found %q", p.ch)}
			}
		case '=': // Read the "=>" token
			pos := p.beg
			p.nextChar()
			if p.ch == '>' {
				p.nextChar()
				return token{pos, tkImpl, "=>"}
			} else {
				return token{pos, tkIllegal, fmt.Sprintf("expecting '>' found %q", p.ch)}
			}
		// 3-char tokens.
		case '<': // Read the "<=>" token
			pos := p.beg
			p.nextChar()
			if p.ch == '=' {
				p.nextChar()
				if p.ch == '>' {
					p.nextChar()
					return token{pos, tkEq, "<=>"}
				} else {
					return token{pos, tkIllegal, fmt.Sprintf("expecting '>' found %q", p.ch)}
				}
			} else {
				return token{pos, tkIllegal, fmt.Sprintf("expecting '=' found %q", p.ch)}
			}

		}
		return token{pos, tkIllegal, fmt.Sprintf("unknown character %q", p.ch)}
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
	case tkTrue:
		return Lit(true), nil

	case tkFalse:
		return Lit(false), nil

	case tkReduce:
		expr, err := l.parse(tk.precedence())
		if err != nil {
			return nil, err
		}
		return Simplify(expr), nil

	case tkAscertain:
		expr, err := l.parse(tk.precedence())
		if err != nil {
			return nil, err
		}
		expr, _ = Factor(expr)
		return expr, nil

	case tkIdentifier:
		return ID(tk.lit), nil

	case tkNot:
		expr, err := l.parse(tk.precedence())
		if err != nil {
			return nil, err
		}
		return Not(expr), nil

	case tkLParen:
		expr, err := l.parse(tk.precedence())
		if err != nil {
			return nil, err
		}
		//and consume the rparen
		if r := l.next(); r.kind != tkRParen {
			return nil, fmt.Errorf("invalid token %s expecting %q", r, tkRParen)
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

	case tkIdentifier:
		// left MUST have been an ID
		id, val, ok := readID(left)
		if !ok {
			return nil, fmt.Errorf("invalid token %s expected an ID expression got %v", tk, left)
		}
		// preserve the current val of this ID. so that any 'not' token in the sentence of identifier switches the value.
		// TODO: decide whether or not we should have a fixed size of identifiers to make an ID.

		// TODO: we are formating ID as a space separated list of identifier. This should be better formalized.
		return minterm{id + " " + tk.lit: val}, nil

	case tkNot:
		// left MUST have been an ID
		id, val, ok := readID(left)
		if !ok {
			return nil, fmt.Errorf("invalid %s expected an ID expression got %v", tk, left)
		}
		return minterm{id: !val}, nil

	case tkOr, tkAnd, tkXor, tkEq, tkNeq, tkImpl, tkLongAnd:
		right, err := p.parse(tk.precedence())
		if err != nil {
			return nil, err
		}
		var res Expr
		switch tk.kind {
		case tkOr:
			res = Or(left, right)
		case tkAnd, tkLongAnd:
			res = And(left, right)
		case tkXor:
			res = Xor(left, right)
		case tkEq:
			res = Eq(left, right)
		case tkNeq:
			res = Neq(left, right)
		case tkImpl:
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
// Variables: any usual identifier `<letter> (digit|letter)*`
//
// an ID is made of one or more Variables.
//
// 'not' can be inserted in a list of variables to form the negation of an ID:
// e.g. `John is not a Knight` is equivalent to `not John is a Knight` or `John not is a Knight`.
//
// Note that the 'not' can be placed anywhere in the ID.
// 'not' can also be used as a prefix operator `not a & b`.
//
// The following operators in increasing binding power (precedence)
//
//   - 'Reduce' reduce expression to a minimal form using [prime implicant]
//   - Ascertain factor an expression x in two implicant expressions a,b so that `x <=> a and b` and return 'a', that can be considered the
//     part of 'x' that is certain.
//   - 'and' long and, same operation as '&' but with low binding power.
//   - '<=>' logically equivalent
//   - '!='  not logically equivalent, same operation as '^' but with low binding power
//   - '=>' imply.
//   - '|' or
//   - '^' xor
//   - '&' and
//   - 'not' not
//
// So that `a | b & c` is equivalent to `a | (b&c)`
//
// Note: '!=' and '^' are the same operation, just with different precedence.
//
// [prime implicant]: https://en.wikipedia.org/wiki/Prime_implicant
func Parse(src string) (Expr, error) {
	return newParser([]byte(src)).parse(0)
}
