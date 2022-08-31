package boolgebra

import (
	"fmt"
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
	Expr[T comparable] []Term[T]

	//Term is a product of identifiers or their negation. For instance
	// mintern  "AB'D" <=> "A or Not(B) or D" is coded as Term{ "A":true, "B":false, "D":true}
	//
	// it is conventional that https://en.wikipedia.org/wiki/Empty_product the empty Term is 1 the neutral for prod ( and for and too)
	//
	Term[T comparable] map[T]bool
)

// String return the literal representation (using primary functions) of the current expression.
func (x Expr[T]) String() string {
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
func (m Term[T]) String() string {
	if len(m) == 0 {
		return "Lit(true)"
	}

	terms := make([]T, 0, len(m))
	for k := range m {
		terms = append(terms, k)
	}
	var sTerms []string
	switch s := any(terms).(type) {
	case []string:
		sort.Strings(s)
		sTerms = s
	case []int:
		sort.Ints(s)
	default:
		sTerms = make([]string, len(terms))
		for i := range terms {
			sTerms[i] = fmt.Sprintf("%v", terms[i])
		}

		sort.Slice(sTerms, func(i, j int) bool {
			return sTerms[i] < sTerms[j]
		})
	}
	if sTerms != nil {
		sTerms = make([]string, len(terms))
		for i := range terms {
			sTerms[i] = fmt.Sprintf("%v", terms[i])
		}
	}

	for i, t := range terms {
		s := sTerms[i]
		if !m[t] {
			sTerms[i] = "Not(" + strconv.Quote(s) + ")"
		} else {
			sTerms[i] = strconv.Quote(s)
		}
	}
	if len(terms) == 1 {
		return sTerms[0]
	} else {
		return "And(" + strings.Join(sTerms, ", ") + ")"
	}
}

// NOT

func (x Expr[T]) Not() Expr[T] {
	factors := make([]Expr[T], 0, len(x))
	for _, e := range x {
		factors = append(factors, e.Not())
	}
	return And(factors...)

}
func (m Term[T]) Not() Expr[T] {
	res := make(Expr[T], 0, len(m))
	for k, v := range m {
		res = append(res, Term[T]{k: !v})
	}
	return res
}

// Is return true if this expression is equals to val
func (x Expr[T]) isLiteral(val bool) bool {
	if val {
		return len(x) == 1 && len(x[0]) == 0
	} else {
		return len(x) == 0
	}
}

// Is return true if this expression is equals to val
func (m Term[T]) isLiteral(val bool) bool {
	return val && len(m) == 0
}

// IDs return the set of ID in this expression
func (x Expr[T]) IDs() (ids map[T]struct{}) {
	ids = make(map[T]struct{})
	for _, m := range x {
		for k := range m {
			ids[k] = struct{}{}
		}
	}
	return
}

// IDs return the set of ID in this expression
func (m Term[T]) IDs() (ids map[T]struct{}) {
	ids = make(map[T]struct{})
	for k := range m {
		ids[k] = struct{}{}
	}
	return
}
