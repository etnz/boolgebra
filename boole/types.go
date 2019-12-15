package boole

import (
	"sort"
	"strconv"
	"strings"
)

/*
to implement Quine-McCluskey algo we need several stages:

- Minterm type
- extracting from the AST, the list (set?) of minterm that represent the sum of product equivalent
- simplify this


*/

//minterm is product of identifier or their negation. For instance
// mintern  "AB'D" is coded as minterm{ "A":true, "B":false, "D":true}
//
// it is conventional that https://en.wikipedia.org/wiki/Empty_product the empty minterm is 1 the neutral for prod ( and for and too)
//
type minterm map[string]bool

func (m minterm) String() string {
	if len(m) == 0 {
		return "True"
	}
	var terms []string
	for k := range m {
		terms = append(terms, k)
	}
	sort.Strings(terms)
	for i, t := range terms {
		if !m[t] {
			terms[i] = "!" + strconv.Quote(t)
		} else {
			terms[i] = strconv.Quote(t)
		}
	}
	return strings.Join(terms, " & ")

}

func (m minterm) Equals(n minterm) bool {
	if m.Len() != n.Len() {
		return false
	}
	for k, v := range m {
		if w, exists := n[k]; !exists || v != w {
			return false
		}
	}
	return true
}

//Len returns the number of identifier in this minterm
func (m minterm) Len() int { return len(m) }

// PosLen returns the number identifier not negated in this minterm
func (m minterm) PosLen() int {
	count := 0
	for _, v := range m {
		if v {
			count++
		}
	}
	return count
}

//NegLen returns the number of identifier negated in this minterm
func (m minterm) NegLen() int { return m.Len() - m.PosLen() }

//inter computes the intersection of bot x, and y
func inter(x, y minterm) minterm {
	res := make(minterm)
	for k, v := range x {
		if w, exists := y[k]; exists && v == w {
			res[k] = v
		}
	}
	return res

}

// combine computes c if possible that combines x and y ( to be defined)
func combine(x, y minterm) (c minterm, ok bool) {
	// alg: find out the one and only one difference between x,y
	// so scan for differences and count.
	var d string // the identifier that is different (if diffs == 1))
	diffs := 0   // number of differences
	for k, v := range x {
		w, exists := y[k]
		if !exists || v != w {
			// this one is different, store it
			d = k
			diffs++
			if diffs > 1 {
				return c, false
			}
		}
	}
	// same goes backward too ( to check for missing in x only

	for k, _ := range y {
		_, exists := x[k]
		if !exists {
			// this one is different, store it
			d = k
			diffs++
			if diffs > 1 {
				return c, false
			}
		}
	}
	if diffs != 1 {
		return c, false
	}
	// we hit the one difference !
	// build c accordingly then
	// x and y are guaranteed to be identical but on 'd'
	// so copy x but 'd'
	c = make(minterm)
	for k, v := range x {
		if k != d {
			c[k] = v
		}
	}
	return c, true

}

// mnot return the negation of the minterm
func mnot(x minterm) (y expression) {

	for id, exponent := range x {
		y = append(y, minterm{id: !exponent})
	}
	return
}

//mand returns x and y.
//
// if x and y hold the same ID, but their exponent is different, we are then doing "a and a'" which is always false
func mand(x, y minterm) (z expression) {
	m := make(minterm)

	// copy x terms, if needed
	for id, exp := range x {
		if yexp, exists := y[id]; exists {
			if yexp != exp {
				// we are doing a and  a' this is always false, this is equivalent to an empty mintern
				return // empty expression (that is false)
			} // else we just skip this one, it will be copied when cloning y
		} else {
			// append it

			m[id] = exp
		}

	}

	for id, exp := range y {
		m[id] = exp
	}
	return expression{m} // a singleton
}

// expression is boolean algebra expression as a sum of prod of minterms. As such it is a slice of minterms. It must be considered as a set ( at least, non-ordered list)
//
// an empty expression is always false ( and this is the definition of false
//
type expression []minterm

//we implement the basic operation on expression to be able to convert the AST into an expression, more suitable for simplification

func (x expression) String() string {
	if len(x) == 0 {
		return "False"
	}
	var terms []string
	for _, k := range x {
		terms = append(terms, k.String())
	}
	return strings.Join(terms, " | ")

}

// IDs return the set of unique ID in this expression
func (x expression) IDs() (ids map[string]struct{}) {
	ids = make(map[string]struct{})
	for _, m := range x {
		for k := range m {
			ids[k] = struct{}{}
		}
	}
	return
}

func id(id string) expression {
	return []minterm{minterm{id: true}}
}

func isLit(x expression, val bool) bool {
	if val {
		return len(x) == 1 && len(x[0]) == 0
	} else {
		return len(x) == 0
	}
}

func not(x expression) (y expression) {

	// create an accumulator that is true
	y = lit(true) // this is TRUE

	for _, m := range x {
		y = and(y, mnot(m))
	}
	return

}

func or(x, y expression) (z expression) {
	//and is simple, just append the lists
	//handle a few special cases
	if isLit(x, true) || isLit(y, true) {
		return lit(true)
	}
	// now append x, and then y only if they are not lit false
	if !isLit(x, false) {
		z = append(z, x...)
	}
	if !isLit(y, false) {
		z = append(z, y...)
	}
	return
}

func and(x, y expression) (z expression) {
	// this one is painful, I need to do all the distribution

	if isLit(x, false) || isLit(y, false) {
		return lit(false)
	}
	if isLit(x, true) {
		return y
	}
	if isLit(y, true) {
		return x
	}
	// general case
	for _, m := range x {
		for _, n := range y {
			z = append(z, mand(m, n)...)
		}
	}
	return
}

func lit(val bool) expression {
	if val {
		return expression{minterm{}}
	} else {
		return expression{}
	}
}
