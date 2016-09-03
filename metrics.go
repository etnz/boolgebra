package boolgebra

import "fmt"

// goal is to measure any given expression

type metrics struct {
	Or, And, Not, Lit, Id uint64
}

func (m *metrics) Add(n metrics) {
	m.Or += n.Or
	m.And += n.And
	m.Not += n.Not
	m.Lit += n.Lit
	m.Id += n.Id
}
func (m metrics) String() string {
	return fmt.Sprintf("Or:%v, And:%v, Not:%v, Id:%v, Lit:%v",
		m.Or,
		m.And,
		m.Not,
		m.Id,
		m.Lit,
	)
}

func metricsOf(X Expression) (m metrics) {
	switch X.Type() {

	case TypeAnd:
		m.And++
		for _, x := range X.Elems() {
			m.Add(metricsOf(x))
		}

	case TypeOr:
		m.Or++
		for _, x := range X.Elems() {
			m.Add(metricsOf(x))
		}

	case TypeNot:
		m.Not++
		m.Add(metricsOf(X.Elems()[0]))

	case TypeLiteral:
		m.Lit++

	case TypeIdentifier:
		m.Id++
	}
	return
}
