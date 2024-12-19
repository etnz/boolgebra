package grid

import (
	"fmt"

	"github.com/etnz/boolgebra"
)

// logic grid game
func ExampleSimplify_logic3x1() {
	// let's define three properties
	g, _ := New(3,
		"Paul", "Luc", "Fernand",
		"A", "B", "C",
		"age5", "age7", "age9",
	)
	solution, _ := g.Solve(

		// 1. Fernand is older than Luc
		`Luc is not age9 & Fernand is not age5`,
		`Luc is age5 => Fernand is age7 | Fernand is age9`,
		`Luc is age7 => Fernand is age9`,

		// 2. The one that plays C is 7 years old.
		`C is age7`,

		// 3. Paul is not the youngest, he plays A
		`Paul is not age5 & Paul is A`,
	)

	fmt.Println(g.Summary(solution))
	//Output:
	// Fernand is C & Fernand is age7 & Luc is B & Luc is age5 & Paul is A & Paul is age9
}

// logic grid game
func ExampleSimplify_logic4x1() {

	g, _ := New(4,
		"Philippe", "Viviane", "Mathilde", "Anne",
		"Math", "French", "Physics", "Sport",
		"R1", "R3", "R5", "R6",
		"A14", "A15", "A17", "A19",
	)

	result, _ := g.Solve(

		// 1. the student that succeeded in Math has 17, is not the first, and doesn't get along with Anne.
		`Math is A17 & Math is not R1 & Anne is not Math`,

		// 2. the student that succeeded in Physics is not Philippe, and has neither the highers mark, nor the lowest.
		`Philippe is not Physics & Physics is not A19 & Physics is not A14`,

		// 3. Mathilde has succeeded in French but she is not in the three best ranked.
		`Mathilde is French`,
		`Mathilde is not R1 & Mathilde is not R3`,

		// 4. Philippe has marked less than 16 and he is ranked 6th.
		`Philippe is not A17 & Philippe is not A19`,
		`Philippe is R6`,
	)

	fmt.Println(g.Summary(result))
	//Output:
	// Anne is A15 & Anne is Physics & Anne is R1 & Mathilde is A19 & Mathilde is French & Mathilde is R5 & Philippe is A14 & Philippe is R6 & Philippe is Sport & Viviane is A17 & Viviane is Math & Viviane is R3
}

func ExampleArt_Contest() {
	// from https://www.brainzilla.com/logic/logic-grid/art-contest/
	// Four artists exhibited their arts using very particular techniques at the national art contest.
	// Knowing that the surfaces are displayed from the softest to the hardest, which tool did Monica use?

	g, err := New(4,
		"aerograph", "brush", "spatula", "sponge",
		"Andrew", "Frank", "Monica", "Sandra",
		"acrylic", "oil", "tempera", "watercolor",
		"paper", "cardboard", "canvas", "wood",
	)
	if err != nil {
		panic(err)
	}

	result, err := g.Solve(
		// 1. Exactly one of the artists has the same initial in its name and its tool.
		`Andrew is aerograph ^ (Sandra is spatula | Sandra is sponge)`,

		// 2. A man used acrylics, a woman painted on paper.
		`Andrew is acrylic | Frank is acrylic`,
		`Monica is paper | Sandra is paper`,

		// 3. Oil colors were used on a less resistant surface than Frank's choice. Frank didn't use the sponge.
		`Frank is wood => oil is canvas | oil is cardboard | oil is paper`,
		`Frank is canvas => oil is cardboard | oil is paper`,
		`Frank is cardboard => oil is paper`,
		`Frank is not sponge`,

		// 4. Among the artist who used the tempera and the artist who painted on wood, one is Monica and the other used the brush.
		`Monica is tempera & brush is wood | Monica is wood & brush is tempera`,
		`Monica is not brush`,
		`tempera is not wood`,

		// 5. The sponge was used with acrylic or on cardboard.
		`sponge is acrylic | sponge is cardboard`,

		// 6. Andrew either used the spatula or painted on canvas.
		`Andrew is spatula ^ Andrew is canvas`,

		// 7. Watercolor was used on a harder surface than tempera.
		`tempera is paper => watercolor is cardboard | watercolor is canvas | watercolor is wood`,
		`tempera is cardboard => watercolor is canvas | watercolor is wood`,
		`tempera is canvas => watercolor is wood`,
		`tempera is not wood`,

		// 8. The aerograph is related to two items that both have the same initial.
		`Andrew is aerograph & aerograph is acrylic | aerograph is watercolor & aerograph is wood`,
	)
	if err != nil {
		panic(err)
	}

	// Which tool did Monica use?
	x, _ := boolgebra.Parse(`Monica is aerograph & Monica is brush & Monica is spatula & Monica is sponge`)
	result, _ = boolgebra.Factor(boolgebra.Or(x, result))

	fmt.Println(result)
	// Output:
	// Monica is aerograph
}

func ExampleChoice_University() {
	// from https://www.brainzilla.com/logic/logic-grid/a-choice-of-university/
	// Four students, from different countries, chose to study in the United Kingdom.
	// Each one chose a subject and a city (where the university is situated).
	// Find out the nationality of the history student.

	g, err := New(4, "Carol", "Elisa", "Lucas", "Oliver",
		"Cambridge", "Edinburgh", "London", "Oxford",
		"Australia", "Canada", "SouthAfrica", "USA",
		"Architecture", "History", "Law", "Medicine")
	if err != nil {
		panic(err)
	}

	result, err := g.Solve(
		// 1. Exactly one boy and one girl chose a university in a city with the same initial of their names.
		`Lucas is London ^ Oliver is Oxford`,
		`Carol is Cambridge ^ Elisa is Edinburgh`,

		// 2. A boy is from Australia, the other studies Architecture.
		`Lucas is Australia | Oliver is Australia`,
		`Lucas is Architecture | Oliver is Architecture`,
		`Australia is not Architecture`,

		// 3. A girl goes to Cambridge, the other studies Medicine.
		`Carol is Cambridge | Elisa is Cambridge`,
		`Carol is Medicine | Elisa is Medicine`,
		`Cambridge is not Medicine`,

		// 4. Oliver studies Law or is from the USA. He is not from South Africa.
		`Oliver is Law | Oliver is USA`,
		`Oliver is not SouthAfrica`,

		// 5. The student from Canada is either a historian or will go to Oxford.
		`Canada is History ^ Oxford is Canada`,

		// 6. The student from the USA will go to Edinburgh or will study Medicine.
		`Edinburgh is USA | USA is Medicine`,

		// 7. Lucas is not from the USA and will not study History.
		`Lucas is not USA & Lucas is not History`,

		// 8. The student from South Africa is going to Edinburgh or will study Law.
		`Edinburgh is SouthAfrica | SouthAfrica is Law`,

		// 9. The Canadian is not studying Law.
		`Canada is not Law`,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(g.Summary(result))
	// Output:
	// Carol is Cambridge & Carol is Canada & Carol is History & Elisa is London & Elisa is Medicine & Elisa is USA & Lucas is Architecture & Lucas is Edinburgh & Lucas is SouthAfrica & Oliver is Australia & Oliver is Law & Oliver is Oxford
}
