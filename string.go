package boolgebra

import (
	"strconv"
	"strings"
)

// String format 'X' in Go.
//
// To copy paste the expression in Go source code, you just need to
//
//    import "boolgebra"
//    var (Or(And(A, B), And(A, Not(B)))
//    	And = boolgebra.And
//    	Or  = boolgebra.Or
//    	Not = boolgebra.Not
//    )
//
func String(X Expression) string {
	switch X.Type() {
	case TypeAnd:
		return "And(" + strings.Join(StringAll(X.Elems()), ", ") + ")"
	case TypeOr:
		return "Or(" + strings.Join(StringAll(X.Elems()), ", ") + ")"
	case TypeNot:
		return "Not(" + String(X.Elems()[0]) + ")"
	case TypeLiteral:
		if X.Val() {
			return "True"
		}
		return "False"

	case TypeIdentifier:
		return X.ID()
	default:
		return "<Invalid>"
	}
}

func nameOf(X Expression) string {
	switch X.Type() {
	case TypeAnd, TypeOr, TypeNot:
		all := make([]string, 0)
		all = append(all, X.Type().String())
		for _, x := range X.Elems() {
			all = append(all, nameOf(x))
		}
		return strings.Join(all, "_")
	case TypeLiteral:
		if X.Val() {
			return "_True"
		}
		return "_False"

	case TypeIdentifier:
		return strings.Replace(X.ID(), " ", "", -1)
	default:
		return "<Invalid>"
	}
}

// LiteralString return a full literal
func LiteralString(X Expression) string {
	switch X.Type() {
	case TypeAnd:
		return "And(" + strings.Join(LiteralStringAll(X.Elems()), ", ") + ")"
	case TypeOr:
		return "Or(" + strings.Join(LiteralStringAll(X.Elems()), ", ") + ")"
	case TypeNot:
		return "Not(" + LiteralString(X.Elems()[0]) + ")"
	case TypeLiteral:
		if X.Val() {
			return "True"
		}
		return "False"

	case TypeIdentifier:
		return "ID(" + strconv.Quote(X.ID()) + ")"
	default:
		return "<Invalid>"
	}
}

// Strings apply the function String elementwise.
func LiteralStringAll(n []Expression) []string {
	res := make([]string, len(n))
	for i, node := range n {
		if node == nil {
			res[i] = "<nil>"
		} else {
			res[i] = LiteralString(node)
		}
	}
	return res
}

// Strings apply the function String elementwise.
func StringAll(n []Expression) []string {
	res := make([]string, len(n))
	for i, node := range n {
		if node == nil {
			res[i] = "<nil>"
		} else {
			res[i] = String(node)
		}
	}
	return res
}
