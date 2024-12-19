package grid

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	. "github.com/etnz/boolgebra" //
	"github.com/etnz/permute"
)

// Grid holds the data about the NxM logic grid.
type Grid struct {
	m, n int
	ids  []string
}

// create a grid with n groups
func New(n int, names ...string) (g *Grid, err error) {

	if len(names)%n != 0 {
		return nil, fmt.Errorf("inconsistent number of names %d with number of groups %d (not divisible)", len(names), n)
	}
	// value set seem to be in good shape, let's generate all the rules
	m := len(names) / n

	// Check for unicity.
	{
		index := make(map[string]struct{}) // index values (that shall be unique)
		var dups []string
		for _, v := range names {
			if _, exists := index[v]; exists {
				dups = append(dups, v)
			}
			index[v] = struct{}{}
		}
		if len(names) != len(index) {
			return nil, fmt.Errorf("names are not unique: %v", dups)
		}
	}
	return &Grid{
		m:   m,
		n:   n,
		ids: names,
	}, nil
}

// Rules returns the logical constraints from the game rules.
func (g *Grid) name(i, j int) string { return g.ids[i*g.n+j] }

// Rules returns the logical constraints from the game rules.
func (g *Grid) Rules() Expr {
	N, M := g.n, g.m

	// solutions for the game can be written as relations to the first property
	//
	//
	//   prop0 , prop1, propi...
	//     val0,  valj,  valk...
	//    ....
	//
	//  prop0 is always values in their original order, then for each column every permutation is a different solution.
	// therefore the number of solutions is (M!)^(N-1)

	// a solution is then made of NxM index to take values from values
	solution := make([][]int, N) // all including column 0 even if we don't sweep it
	for i := range solution {
		solution[i] = make([]int, M)
		for j := range solution[i] {
			solution[i][j] = j
		}
	}

	// next function computes the next solution by permuting a column until completion, and then the other etc.
	next := func() bool {
		c := 1 // always start incrementing skipping the first column
		for c < len(solution) && !permute.LexNext(solution[c]) {
			c++ // current column was incremented, but reached the end, I need to increment the next one, hence the c++
		}
		return c < len(solution)
	}

	// find the offset in the column 0 that is related to the value i.
	// in other words, if "Paul" (position 3) is "5yo" (position 12) then whois(12) return 3. Who is "5yo"? "Paul"
	whois := func(i int) (i0 int) {
		// i/M is the column in the solution of the ith value
		// i%M is the value offset in the column ( like 0 for the first value).
		// so simply search in the column for it. It gives us the row that is related to i.
		// the column 0 value is exactly this row, so we just return it
		for i0 = 0; solution[i/M][i0] != i%M; i0++ {
		}
		return
	}

	//rules := make([]map[string]bool, 0) // the game rules will be an Or() of all possible solutions
	rules := Lit(false)
	for ok := true; ok; ok = next() { // loop over all columns x all permutations, and break when done

		var tb TermBuilder // build a big AND expression

		// for the given solution, scan all possible ID ( "Paul is 6yo") wether it is true or not
		for i, v := range g.ids {
			for j, w := range g.ids {
				if i != j {
					tb.And(v+" is "+w, whois(i) == whois(j))
				}
			}
		}

		// tb contains the ONE solution expressed as boolean
		rules = Or(rules, tb.Build()) // append it
	}
	return rules
}

func (g *Grid) Solve(hints ...string) (Expr, error) {
	clock := time.Now()
	r := g.Rules()
	log.Printf("solving rules-minterms=%v  in=%v", r.Terms(), time.Since(clock))

	// Parses the hints, and perform some sanity checks on each of them.
	all := r.IDs()
	var hintList []Expr
	clock = time.Now()
	for _, h := range hints {
		x, err := Parse(h)
		if err != nil {
			return nil, fmt.Errorf("invalid hint %q. Parse error: %v", h, err)
		}
		// Check the IDs introduced by this hint.
		for k := range x.IDs() {
			// They must not introduce an unknown relation. This is most likely a typo.
			if _, exists := all[k]; !exists {
				log.Printf("warning hint=%q new ID=%q", h, k)
			}
			words := strings.Split(k, " ")
			if len(words) != 3 {
				log.Printf("warning hint=%q invalid-length ID=%q", h, k)
			}
			if words[1] != "is" {
				log.Printf("warning hint=%q invalid-is ID=%q", h, k)
			}
		}
		hintList = append(hintList, x)
	}

	h := And(hintList...)
	log.Printf("solving hints=%v hints-minterms=%v in=%v", len(hintList), h.Terms(), time.Since(clock))

	clock = time.Now()
	// We figured out (without understanding why) that And(rules, hints...) faster than And(rules, And(hints...) )
	// also order matters. ( And(rules, hints...) faster than And(hints..., rules)
	//
	x := Simplify(And(append([]Expr{r}, hintList...)...))
	log.Printf("solving solutions=%v in=%v", x.Terms(), time.Since(clock))
	return x, nil

}

// Summary returns only the true significant relations if result has only 1 Term.
func (g *Grid) Summary(result Expr) Expr {
	if result.Terms() != 1 {
		return nil
	}
	var solutions []string
	for id := range result.Term(0).IDs() {
		x, positive := result.ID(id)
		if x != nil && positive {
			solutions = append(solutions, id)
		}
	}
	sort.Strings(solutions)

	var tb TermBuilder

	for _, name := range g.ids[:g.m] {
		for _, sol := range solutions {
			if strings.HasPrefix(sol, name) {
				tb.And(sol, true)
			}
		}
	}
	return tb.Build()
}

// Filter
func (g *Grid) Filter(result Expr, id1, id2 string) Expr {
	if result.Terms() != 1 {
		return nil
	}
	t := result.Term(0)
	for id := range t.IDs() {
		if strings.HasPrefix(id, id1) && strings.HasSuffix(id, id2) || strings.HasPrefix(id, id2) && strings.HasSuffix(id, id1) {
			x, _ := t.ID(id)
			return x
		}
	}
	return nil
}
