package boolgebra

import (
	"sort"
	"strconv"
	"strings"
)

type (
	// Expr is boolean algebra Expr as a sum of prod of minterms.
	// As such it is a slice of minterms. It must be considered as a set
	//
	// an empty Expr is always false ( and this is the definition of false
	//
	Expr []Term

	//Term is a product of identifiers or their negation. For instance
	// mintern  "AB'D" <=> "A or Not(B) or D" is coded as Term{ "A":true, "B":false, "D":true}
	//
	// it is conventional that https://en.wikipedia.org/wiki/Empty_product the empty Term is 1 the neutral for prod ( and for and too)
	//
	Term map[string]bool
)

// String return the literal representation (using primary functions) of the current expression.
func (x Expr) String() string {
	if len(x) == 0 {
		return "Lit(false)"
	}
	if len(x) == 1 {
		return x[0].String()
	}

	var terms []string
	for _, k := range x {
		terms = append(terms, k.String())
	}

	if len(terms) > 3 {
		return "Or(\n    " + strings.Join(terms, ",\n    ") + "\n)"
	}
	return "Or(" + strings.Join(terms, ", ") + ")"
}

// String return the literal representation (using primary functions) of the current minterm
func (m Term) String() string {
	if len(m) == 0 {
		return "Lit(true)"
	}

	var terms []string
	for k := range m {
		terms = append(terms, k)
	}
	sort.Strings(terms)
	for i, t := range terms {
		if !m[t] {
			terms[i] = "Not(" + strconv.Quote(t) + ")"
		} else {
			terms[i] = strconv.Quote(t)
		}
	}
	if len(terms) == 1 {
		return terms[0]
	} else {
		return "And(" + strings.Join(terms, ", ") + ")"
	}
}

// NOT

func (x Expr) Not() Expr {
	factors := make([]Expr, 0, len(x))
	for _, e := range x {
		factors = append(factors, e.Not())
	}
	return And(factors...)

}
func (m Term) Not() Expr {
	res := make(Expr, 0, len(m))
	for k, v := range m {
		res = append(res, Term{string(k): !v})
	}
	return res
}

// Is return true if this expression is equals to val
func (x Expr) isLiteral(val bool) bool {
	if val {
		return len(x) == 1 && len(x[0]) == 0
	} else {
		return len(x) == 0
	}
}

// Is return true if this expression is equals to val
func (m Term) isLiteral(val bool) bool {
	return val && len(m) == 0
}

// IDs return the set of ID in this expression
func (x Expr) IDs() (ids map[string]struct{}) {
	ids = make(map[string]struct{})
	for _, m := range x {
		for k := range m {
			ids[k] = struct{}{}
		}
	}
	return
}

// IDs return the set of ID in this expression
func (m Term) IDs() (ids map[string]struct{}) {
	ids = make(map[string]struct{})
	for k := range m {
		ids[k] = struct{}{}
	}
	return
}
