package boolgebra

import (
	"sort"
	"strings"
)

type (
	// expression is boolean algebra expression as a sum of prod of minterms.
	// As such it is a slice of minterms. It must be considered as a set
	//
	// an empty expression is always false ( and this is the definition of false
	//
	expression []minterm

	//minterm is a product of identifier or their negation. For instance
	// mintern  "AB'D" <=> "A or Not(B) or D" is coded as minterm{ "A":true, "B":false, "D":true}
	//
	// it is conventional that https://en.wikipedia.org/wiki/Empty_product the empty minterm is 1 the neutral for prod ( and for and too)
	//
	minterm map[string]bool

	// Expr is the interface that all elements of a boolean algebra share
	Expr interface {
		String() string
		// return the negation of receiver
		Not() Expr
		// Is return true if the Expr is literally equals to value
		Is(val bool) bool
		Terms() int
		Term(i int) Expr
		IDs() (ids map[string]struct{})
	}
)

// String return the literal representation (using primary functions) of the current expression.
func (x expression) String() string {
	if len(x) == 0 {
		return "false"
	}
	if len(x) == 1 {
		return x[0].String()
	}

	var terms []string
	for _, k := range x {
		terms = append(terms, k.String())
	}

	return strings.Join(terms, " | ")
}

// String return the literal representation (using primary functions) of the current minterm
func (m minterm) String() string {
	if len(m) == 0 {
		return "true"
	}

	var terms []string
	for k := range m {
		terms = append(terms, k)
	}
	sort.Strings(terms)
	for i, t := range terms {
		if !m[t] {
			identifiers := strings.Split(t, " ")
			if len(identifiers) < 3 { // short id, use 'not id'
				identifiers = append([]string{"not"}, identifiers...)
			} else { // put a not in third position (typical for 'a is not x')
				identifiers = append(identifiers, "not") // Ensure capacity.
				copy(identifiers[3:], identifiers[2:])   // Shift index up.
				identifiers[2] = "not"                   // Insert item.
			}
			terms[i] = strings.Join(identifiers, " ")
		} else {
			terms[i] = t
		}
	}
	if len(terms) == 1 {
		return terms[0]
	} else {
		return strings.Join(terms, " & ")
	}
}

// NOT

func (x expression) Not() Expr {
	factors := make([]Expr, 0, len(x))
	for _, e := range x {
		factors = append(factors, e.Not())
	}
	return And(factors...)

}
func (m minterm) Not() Expr {
	res := make(expression, 0, len(m))
	for k, v := range m {
		res = append(res, minterm{string(k): !v})
	}
	return res
}

// Is return true if this expression is equals to val
func (x expression) Is(val bool) bool {
	if val {
		return len(x) == 1 && len(x[0]) == 0
	} else {
		return len(x) == 0
	}
}

// Is return true if this expression is equals to val
func (m minterm) Is(val bool) bool {
	return val && len(m) == 0
}

// Terms retuns the number of terms in this expression
func (x expression) Terms() int { return len(x) }

// Terms retuns the number of terms in this expression
func (m minterm) Terms() int { return 1 }

// Term retuns the ith terms. Panic if out of bounds ( negative, or >= Terms())
func (x expression) Term(i int) Expr {
	if i < 0 || i >= x.Terms() {
		panic("Term is not defined for this index value")
	}
	return x[i]
}

// Term retuns the ith terms. Panic if out of bounds ( negative, or >= Terms())
func (m minterm) Term(i int) Expr {
	if i != 0 {
		panic("Term is not defined for this index value")
	}
	return m
}

// IDs return the set of ID in this expression
func (x expression) IDs() (ids map[string]struct{}) {
	ids = make(map[string]struct{})
	for _, m := range x {
		for k := range m {
			ids[k] = struct{}{}
		}
	}
	return
}

// IDs return the set of ID in this expression
func (m minterm) IDs() (ids map[string]struct{}) {
	ids = make(map[string]struct{})
	for k := range m {
		ids[k] = struct{}{}
	}
	return
}
