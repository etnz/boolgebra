//Package boolgebra provide basic boolean algebra operation and simplification.
//
// Simplifications
//
// A few laws rule the boolean algebra. see https://en.wikipedia.org/wiki/Boolean_algebra#Laws
//
// Among them, simplification laws are those, by definition, that reduce the complexity of
// the gobal expression. For instance
//    And(True, x) -> x
// Left-hand side complexity (3 maybe) is greater than the right-hand side (1).
//
// The goal of boolgebra is to provide the best possible simplification algorithm, because
// boolean expression simplification outsmarts most of us humans ! See Examples from Mullyan's book
//
// Literal simplification Rules:
//
//      And(True, x) -> x             Identity
//      And(False,x) -> False         Annihilator
//      Or(True, x)  -> True          Annihilator
//      Or(False,x)  -> x             Identity
//
//
// Negation Rules:
//
//      Not(Not(x))  -> x
//      Not(False)   -> True
//      Not(True)    -> False
//
//
// Complementation Rules:
//
//      And(x, Not(x))-> False        Complementation
//      Or(x, Not(x)) -> True         Complementation
//
// Idempotence Rules:
//
//      And(x, x)-> x
//      Or(x, x) -> x
//
//
package boolgebra

import "fmt"

const (
	True  literalExpr = true  // Simplest "True" expression
	False literalExpr = false // Simplest "False" expression
)

const (
	// all possible Expression types
	//
	// The order of this constant defines the natural order for Type
	TypeInvalid Type = iota // n.b: zero value for Type is invalid
	TypeLiteral
	TypeIdentifier
	TypeOr
	TypeAnd
	TypeNot
)

// Expression is a boolean arithmetic expression tree
//
// It is defined by its root element, which can be:
//
//    - an Identifier: a string
//    - a Literal: a bool
//    - a binary Operator: Or, And
//    - the negation operator: Not
//
// Nothing more.
//
type Expression interface {
	// Type returns the current type (never TypeInvalid)
	Type() Type

	// Elems returns direct children, in particular, for type:
	//     Or, And   : returns the list of children expressions
	//     Not       : returns a slice containing only the negated expression
	//     by default: returns nil
	Elems() []Expression

	// Id returns the identifier for an expression, in particular, for type:
	//     Identifier: returns its name
	//     by default: returns empty string
	ID() string

	// Val returns the bool value of an expression, in particular, for type:
	//     Literal   : returns its boolean value
	//     by default: returns false
	Val() bool
}

// Type is one of the possible Node type
type Type uint8

// String returns a literal representation of the type
func (t Type) String() string {
	switch t {
	case TypeInvalid:
		return "TypeInvalid"
	case TypeLiteral:
		return "TypeLiteral"
	case TypeIdentifier:
		return "TypeIdentifier"
	case TypeOr:
		return "TypeOr"
	case TypeAnd:
		return "TypeAnd"
	case TypeNot:
		return "TypeNot"
	default:
		return fmt.Sprintf("Type(%v)", uint8(t))
	}
}

// here are all the nodes type.
// there are not too many of them

type (
	andExpr        []Expression
	orExpr         []Expression
	notExpr        [1]Expression
	literalExpr    bool
	identifierExpr string
)

//Fluent constructors

// ID returns the expression made of a single Identifier
func ID(a string) Expression { return identifierExpr(a) }

// Lit returns the expression made of a single Literal with value 'a'
func Lit(a bool) Expression { return literalExpr(a) }

// Or returns an expression that is logically the Or of all expressions
//
// if called with:
//    0 argument: returns False expression ( Identify for Or)
//    1 argument: returns the argument itself
//    any other : build an actual TypeOr expression
func Or(expressions ...Expression) Expression {
	switch len(expressions) {
	case 0:
		return False
	case 1:
		return expressions[0]
	default:
		return orExpr(expressions[:])
	}
}

// And returns an expression that is logically the And of all expressions
//
// if called with:
//    0 argument: returns True expression ( Identify for And)
//    1 argument: returns the argument itself
//    any other : build an actual TypeAnd expression
func And(expressions ...Expression) Expression {
	switch len(expressions) {
	case 0:
		return True
	case 1:
		return expressions[0]
	default:
		return andExpr(expressions[:])
	}
}

// Not return an Expression equivalent to 'Not(n)'
func Not(n Expression) Expression { return notExpr{n} }

// all types Elems methods

func (n andExpr) Elems() []Expression        { return n[:] }
func (n orExpr) Elems() []Expression         { return n[:] }
func (n notExpr) Elems() []Expression        { return n[:] }
func (n literalExpr) Elems() []Expression    { return nil }
func (n identifierExpr) Elems() []Expression { return nil }

// all types Type methods

func (n andExpr) Type() Type        { return TypeAnd }
func (n orExpr) Type() Type         { return TypeOr }
func (n notExpr) Type() Type        { return TypeNot }
func (n literalExpr) Type() Type    { return TypeLiteral }
func (n identifierExpr) Type() Type { return TypeIdentifier }

// all types Val methods

func (n andExpr) Val() bool        { return false }
func (n orExpr) Val() bool         { return false }
func (n notExpr) Val() bool        { return false }
func (n literalExpr) Val() bool    { return bool(n) }
func (n identifierExpr) Val() bool { return false }

// all types ID methods

func (n andExpr) ID() string        { return "" }
func (n orExpr) ID() string         { return "" }
func (n notExpr) ID() string        { return "" }
func (n literalExpr) ID() string    { return "" }
func (n identifierExpr) ID() string { return string(n) }

func (n andExpr) String() string        { return String(n) }
func (n orExpr) String() string         { return String(n) }
func (n notExpr) String() string        { return String(n) }
func (n literalExpr) String() string    { return String(n) }
func (n identifierExpr) String() string { return String(n) }
