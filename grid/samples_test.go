package grid

import (
	"fmt"

	. "github.com/etnz/boolgebra"
)

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

	result := Solve(4, []string{Philippe, Viviane, Mathilde, Anne,
		Math, French, Physics, Sport,
		R1, R3, R5, R6,
		A14, A15, A17, A19,
	},

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

func ExampleArt_Contest() {
	// from https://www.brainzilla.com/logic/logic-grid/art-contest/
	// Four artists exhibited their arts using very particular techniques at the national art contest.
	// Knowing that the surfaces are displayed from the softest to the hardest, which tool did Monica use?

	aerograph, brush, spatula, sponge := "aerograph", "brush", "spatula", "sponge"
	Andrew, Frank, Monica, Sandra := "Andrew", "Frank", "Monica", "Sandra"
	acrylic, oil, tempera, watercolor := "acrylic", "oil", "tempera", "watercolor"
	paper, cardboard, canvas, wood := "paper", "cardboard", "canvas", "wood"

	result := Solve(4, []string{
		aerograph, brush, spatula, sponge,
		Andrew, Frank, Monica, Sandra,
		acrylic, oil, tempera, watercolor,
		paper, cardboard, canvas, wood},

		// 1. Exactly one of the artists has the same initial in its name and its tool.
		Xor(P(Andrew, aerograph), Or(P(Sandra, spatula), P(Sandra, sponge))),

		// 2. A man used acrylics, a woman painted on paper.
		Xor(P(Andrew, acrylic), P(Frank, acrylic)),
		Xor(P(Monica, paper), P(Sandra, paper)),

		// 3. Oil colors were used on a less resistant surface than Frank's choice. Frank didn't use the sponge.
		Impl(P(Frank, wood), Or(P(oil, canvas), P(oil, cardboard), P(oil, paper))),
		Impl(P(Frank, canvas), Or(P(oil, cardboard), P(oil, paper))),
		Impl(P(Frank, cardboard), P(oil, paper)),
		Not(P(Frank, sponge)),

		// 4. Among the artist who used the tempera and the artist who painted on wood, one is Monica and the other used the brush.
		Or(And(P(Monica, tempera), P(brush, wood)), And(P(Monica, wood), P(brush, tempera))),
		Not(P(tempera, wood)),

		// 5. The sponge was used with acrylic or on cardboard.
		Or(P(sponge, acrylic), P(sponge, cardboard)),

		// 6. Andrew either used the spatula or painted on canvas.
		Or(P(Andrew, spatula), P(Andrew, canvas)),

		// 7. Watercolor was used on a harder surface than tempera.
		Impl(P(tempera, paper), Or(P(watercolor, cardboard), P(watercolor, canvas), P(watercolor, wood))),
		Impl(P(tempera, cardboard), Or(P(watercolor, canvas), P(watercolor, wood))),
		Impl(P(tempera, canvas), P(watercolor, wood)),
		Not(P(tempera, wood)),

		// 8. The aerograph is related to two items that both have the same initial.
		Or(
			And(P(aerograph, Andrew), P(aerograph, acrylic)),
			And(P(aerograph, watercolor), P(aerograph, wood)),
		),
	)

	answer, _ := Factor(Or(
		And(P(Monica, aerograph), P(Monica, brush), P(Monica, spatula), P(Monica, sponge)),
		result,
	))
	fmt.Printf("The answer is %q\n", answer)
	// Output: The answer is "Monica=aerograph"
}
