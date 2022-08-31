package grid

import (
	"fmt"

	. "github.com/etnz/boolgebra" //
	"github.com/etnz/permute"
)

// creates an ID that means property name = value
// just for test, not safe for any kind of injection
func P(name, value string) Expr { return ID(name + "=" + value) }

// Values is a list all possible values without repetition
//
// a Group is a subset of Values, Groups is the set of all defined Group, N is card(Groups)
//
// Let's define 'R' an transitive and symetric relation in Values noted `\forall x,y \in Values xRy`
//
//  1. `\forall g,h \in Groups² |g| = |h| \and g \inter h = \phi`
//  2. `\forall G \in Groups, \forall v \notin G \exists! w in G vRw`
//
// groups are defined by to position in the list
func Rules(N int, values ...string) Expr {

	if len(values)%N != 0 {
		panic(fmt.Sprintf("inconsistent number of values %d with number of groups %d (not divisible)", len(values), N))
	}
	// value set seem to be in good shape, let's generate all the rules
	M := len(values) / N

	{ // the following code is in a a block 'cause I don't keep anyting from here the index is purely local, and temporary
		index := make(map[string]struct{}) // index values (that shall be unique)
		for _, v := range values {
			index[v] = struct{}{}
		}
		if len(values) != len(index) {
			panic("values have repetitions")
		}
	}

	// solutions for the game can be written as relations to the first property
	//
	//
	//   prop0 , prop1, propi...
	//     val0,  valj,  valk...
	//    ....
	//
	//  prop0 is always values in their original order, then each column, can be sorted randomly => any permutation will do (there are M! of such permutations)
	//  so we have to sweep every permutation of columns from 1 to N-1
	//
	// therefore the number of solutions is (M!)^(N-1)

	// the first goal is to generate sol the vector of permutations, initialised with the identity permutation
	sol := make([][]int, N) // all including column 0 even if we don't sweep it
	for i := range sol {
		sol[i] = make([]int, M)
		for j := range sol[i] {
			sol[i][j] = j
		}
	}

	// next computes the next solution
	next := func() bool {
		c := 1 // always start incrementing skipping the first column
		for c < len(sol) && !permute.LexNext(sol[c]) {
			c++ // current column was incremented, but reached the end, I need to increment the next one, hence the c++
		}
		return c < len(sol)
	}

	// find the offset in the column 0 that is related to the value i.
	// in other words, if "Paul" (position 3) is "5yo" (position 12) then whois(12) return 3. Who is "5yo"? "Paul"
	whois := func(i int) (i0 int) {
		// i/M is the column in the solution of the ith value
		// i%M is the value offset in the column ( like 0 for the first value).
		// so simply search in the column for it. It gives us the row that is related to i.
		// the column 0 value is exactly this row, so we just return it
		for i0 = 0; sol[i/M][i0] != i%M; i0++ {
		}
		return
	}

	//rules := make([]map[string]bool, 0) // the game rules will be an Or() of all possible solutions
	rules := Lit(false)
	for ok := true; ok; ok = next() { // loop over all columns x all permutations, and break when done

		var tb TermBuilder // build a big AND expression

		// for the given solution, scan all possible ID ( "Paul is 6yo") wether it is true or not
		for i, v := range values {
			for j, w := range values {
				if i != j {
					tb.And(v+"="+w, whois(i) == whois(j))
				}
			}
		}

		// tb contains the ONE solution expressed as boolean
		rules = Or(rules, tb.Build()) // append it
	}
	// now I need to convert it to the boolgebra

	return rules
}

func Solve(nbproperties int, values []string, hints ...Expr) Expr {

	rules := []Expr{}
	rules = append(rules, Rules(nbproperties, values...))
	rules = append(rules, hints...)
	return Simplify(And(rules...))
}
