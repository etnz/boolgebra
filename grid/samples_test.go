package grid

import (
	"fmt"
	"sort"

	. "github.com/etnz/boolgebra"
)

// TODO this is copy/paste from MY own other code, use a better approach (module when improving logic graph solving)

// LexNext finds the next permutation in lexicographical order.
//
// return false if it has gone back to the identity permutation.
//
// inspired from Narayana Pandita in https://en.wikipedia.org/wiki/Permutation
func LexNext(p []int) bool {

	// the principle for lehmer code is to convert an factoradic number into its permutation:
	//
	// each factoradic digit stand for the selection of an index, among all possible.
	//
	// Excerpt from wikipedia:
	//
	// The process may become clearer with a longer example. For example, here is how the digits in the factoradic 4041000! (equal to 298210) pick out the digits in (4,0,6,2,1,3,5), the 2982nd permutation of the numbers 0 through 6.
	//
	//                                  4041000! → (4,0,6,2,1,3,5)
	// factoradic:  4          0                        4        1          0          0        0!
	//              |          |                        |        |          |          |        |
	//     (0,1,2,3,4,5,6) -> (0,1,2,3,5,6) -> (1,2,3,5,6) -> (1,2,3,5) -> (1,3,5) -> (3,5) -> (5)
	//              |          |                        |        |          |          |        |
	// permutation:(4,         0,                       6,       2,         1,         3,       5)
	//
	//
	// One property is that the last index is always 0 (there is no choice left)
	//
	// If I want to compute the next one (the goal here), I need to compute the next factoradic number by applying two simple rules:
	//
	// if the current digit is the last of its position (0! 1! 2! 3! etc), then I set it to 0, and increment the next digit.
	//
	// else simply increment it
	//
	// In the example it would be :
	//
	// digit0 = 0 last possible one, move to next
	//
	// digit1 = 3 it's the index 0 of (3,5), so increment the index (1) and apply digit1 := 5
	//
	//
	// We can find out that a digit can be incremented iif it is *not* the max of the previous digits values.
	//
	// In the example previous digits values are 5,3 , and 3 is not the max, it can be simply incremented

	// the first number that is not the "max" of the previous values, is the one that can be "incremented"

	i := len(p) - 2 // loop from last but one digit
	// the last one is always 0, and can never be incremented
	for i >= 0 && p[i] > p[i+1] {
		i--
	}
	// this test seems odd, but it works by recursion:
	// recursive hypothethis: p[i+1] is the max of the remaining digits
	// therefore  p[i]> p[i+1] =>  p[i] is also the max of the remaining digits
	//

	// Therefore now 'i' is either -1 meaning that we have reached the end, the next is the first in fact
	if i < 0 {
		sort.Ints(p) //we start again
		return false
	}

	// we keep the actual index at i
	val := p[i]     // this is the one that will be "incremented" ( next value in possible ones)
	vals := p[i+1:] // subslice of remaining values (one is greater that val, rembember!)
	sort.Ints(vals) // sort them, 1- to search among them, but also because this is how they will be added

	pos := sort.SearchInts(vals, val) // find where I would insert val (this is the next one)
	p[i], vals[pos] = vals[pos], val  // permute val, with the next one (found above)

	// and now copy vals to the rest of p
	copy(p[i+1:], vals)
	return true
}

// creates an ID that means property name = value
// just for test, not safe for any kind of injection
func P(name, value string) Expr { return ID(name + "=" + value) }

// Values is a list all possible values without repetition
//
// a Group is a subset of Values, Groups is the set of all defined Group, N is card(Groups)
//
// Let's define 'R' an transitive and symetric relation in Values noted `\forall x,y \in Values xRy`
//
//     1. `\forall g,h \in Groups² |g| = |h| \and g \inter h = \phi`
//     2. `\forall G \in Groups, \forall v \notin G \exists! w in G vRw`
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
		for c < len(sol) && !LexNext(sol[c]) {
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

// logic grid game
func ExampleSimplify_logic3x1() {
	// let's define three properties
	//"name"
	Paul, Luc, Fernand := "Paul", "Luc", "Fernand"
	//"game"
	A, B, C := "A", "B", "C"

	// "age"
	Age5, Age7, Age9 := "5", "7", "9"

	// generate all the logical relations due to game rules
	rules := Rules(3, /*groups*/
		Paul, Luc, Fernand,
		A, B, C,
		Age5, Age7, Age9)

	//log.Printf("rules: len=%v ids=%v\n%v", rules.Terms(), len(rules.IDs()), rules)

	// hints

	// Fernand is older than Luc
	Hint1 := And(
		Not(P(Luc, Age9)),
		Not(P(Fernand, Age5)),
		Impl(P(Luc, Age5), Or(P(Fernand, Age7), P(Fernand, Age9))),
		Impl(P(Luc, Age7), P(Fernand, Age9)),
		//Impl(P(Luc, Age9), Lit(false)),
	)

	// The one that plays C is 7 years old.
	Hint2 := P(C, Age7)

	// Paul is not the youngest, he plays A
	Hint3 := And(
		Not(P(Paul, Age5)),
		P(Paul, A),
	)

	result := Simplify(And(rules, Hint1, Hint2, Hint3))

	if result.Terms() > 1 {
		fmt.Printf("There are %d solutions, that's too many\n", result.Terms())
		fmt.Println(Factor(result))
	} else {
		fmt.Printf("There is %d solution.\n", result.Terms())
	}
	//Output:
	// There is 1 solution.
}

// logic grid game
func ExampleSimplify_logic4x1() {

	Philippe, Viviane, Mathilde, Anne := "nPhilippe", "nViviane", "nMathilde", "nAnne"
	Math, French, Physics, Sport := "mMath", "mFrench", "mPhysics", "mSport"
	R1, R3, R5, R6 := "R1", "R3", "R5", "R6"
	A14, A15, A17, A19 := "A14", "A15", "A17", "A19"

	// generate all the logical relations due to game rules
	rules := Rules(4,
		Philippe, Viviane, Mathilde, Anne,
		Math, French, Physics, Sport,
		R1, R3, R5, R6,
		A14, A15, A17, A19,
	)

	// hints

	Hints := And(
		// - L'élève qui réussit en maths a 17 de moyenne, n'est pas premier et s'entend bien avec Anne.
		P(Math, A17),
		Not(P(Math, R1)),
		Not(P(Anne, Math)),

		// - L'élève qui réussit en Sciences Physiques n'est pas Philippe et n'a ni la plus haute moyenne ni la plus basse moyenne.
		Not(P(Philippe, Physics)),
		Not(P(Physics, A19)),
		Not(P(Physics, A14)),

		// - Mathilde réussit bien en Français mais elle n'est pas dans les trois premiers de sa classe.
		P(Mathilde, French),
		Not(P(Mathilde, R1)),
		Not(P(Mathilde, R3)),

		// - Philippe a moins de 16 de moyenne dans sa matière ce qui le met 6ème de sa classe.
		Not(P(Philippe, A17)),
		Not(P(Philippe, A19)),

		//Or(P(Philippe, A14), P(Philippe, A15)),

		P(Philippe, R6),
	)

	result := Simplify(And(rules, Hints))

	if result.Terms() > 1 {
		fmt.Printf("There are %d solutions, that's too many\n", result.Terms())
		deduction, rem := Factor(result)
		fmt.Println(deduction)
		fmt.Println(rem)
	} else {
		fmt.Printf("There is %d solution.\n", result.Terms())
	}
	//Output:
	// There is 1 solution.
}
