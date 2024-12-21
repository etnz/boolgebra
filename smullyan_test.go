package boolgebra

import (
	"fmt"
)

// smullyan_test.go holds examples taken from Raymond Smullyan "Satan, Cantor and Infinity"
// In the book,
// a Knight is someone that always tells the truth,
// a Knaves is someone that always tells a lie.
//
// In all the Smullyan examples it is important to understand that the logical translation of a sentence like
//  'A say P' is that `A is a Knight <=> P`.
//

// ExampleSimplify_BothKnaves solves the [Both Knaves] problem.
// [Both Knaves]: https://en.wikipedia.org/wiki/Knights_and_Knaves#Both_knaves
func ExampleSimplify_BothKnaves() {
	// Alice says, "We are both knaves”.
	solution := `Reduce Alice is a Knight <=> Alice is not a Knight & Bob is not a Knight`
	expression, _ := Parse(solution)
	fmt.Println(expression)
	// Output: Alice is not a Knight & Bob is a Knight
}

// ExampleSimplify_SameOfDifferentKinds solves the [Same of Different Kinds] problem.
// [Same of Different Kinds]: https://en.wikipedia.org/wiki/Knights_and_Knaves#Same_or_different_kinds
func ExampleSimplify_SameOfDifferentKinds() {
	// Alice says, "We are the same kind," but Bob says, "We are of different kinds."
	solution := `Reduce 
		    Alice is a Knight <=> not (Alice is a Knight ^ Bob is a Knight) 
		and Bob is a Knight <=> Alice is a Knight ^ Bob is a Knight`
	expression, _ := Parse(solution)
	fmt.Println(expression)
	// Output: Alice is not a Knight & Bob is a Knight
}

// ExampleSimplify_GoodmanVariant solves the [Goodman Variant] problem.
// [Goodman Variant]: https://en.wikipedia.org/wiki/Knights_and_Knaves#Goodman's_1931_variant
func ExampleSimplify_GoodmanVariant() {
	// In goodman Variant, Knights and Krave are replaced by Nobles and Hunters, we translated it to unify.
	// Alice says either "I am a Knight" or "I am a Knave", we don't yet know which.
	// Then Bob, in reply to a query, says "Alice said, 'I am a Knave'".
	// After that, Bob says "Carol is a Knave".
	// Then, Carol says "Alice is a Knight"

	// Noone can ever say "I am a Knave", so the first hint is useless, but let's code it anyway.

	solution := `Reduce 
	    Alice is a Knight <=> Alice is not a Knight | Alice is a Knight
		and Bob is a Knight <=> (Alice is a Knight <=> Alice is not a Knight) 
		and Bob is a Knight <=> Carol is not a Knight
		and Carol is a Knight <=> Alice is a Knight
		`

	expression, _ := Parse(solution)
	fmt.Println(expression)

	// Output: Alice is a Knight & Bob is not a Knight & Carol is a Knight
}

func ExampleSimplify_ThreeKnaves() {

	// Our hero asked Alice: "are Bob and Carol knights?"
	// Alice answered "yes", but Alice also said that Bob was a Knaves

	solution := `Reduce 
	    Alice is a Knight <=> Bob is a Knight & Carol is a Knight
	and Alice is a Knight <=> Bob is not a Knight
	`
	expression, _ := Parse(solution)
	fmt.Println(expression)

	//Output: Alice is not a Knight & Bob is a Knight & Carol is not a Knight
}

func ExampleSimplify_smullyan2() {
	// this example cannot use the bool language, yet, as support for Exactly 1 is not available yet.

	// A Knight is someone that always tells the truth.
	// A Knaves is someone that always tells a lie.

	// The hero meets three guys 'a','b', 'c', each one can be a Knight or a Knaves
	A_is_a_knight := ID("A is a Knight")
	B_is_a_knight := ID("B is a Knight")
	C_is_a_knight := ID("C is a Knight")

	// the goal is to find out who is a Sorcerer

	A_is_Sorcerer := ID("A is a Sorcerer")
	B_is_Sorcerer := ID("B is a Sorcerer")
	C_is_Sorcerer := ID("C is a Sorcerer")

	// But only one can be a Sorcerer
	Fact1 := Exactly(1, A_is_Sorcerer, B_is_Sorcerer, C_is_Sorcerer)

	// The hero asked 'a': are you a Sorcerer ? 'a' answered, yes
	// as always with knights and knaves, we can say that:
	Fact2 := Eq(A_is_a_knight, A_is_Sorcerer)

	// He asked the same question to 'b'
	Fact3 := Eq(B_is_a_knight, B_is_Sorcerer)

	// But 'c', said that, at most one of them was a Knight !!
	Fact4 := Eq(C_is_a_knight, AtMost(1, A_is_a_knight, B_is_a_knight, C_is_a_knight))

	// The answer is straightforward.
	fmt.Println(Simplify(And(Fact1, Fact2, Fact3, Fact4)))

	//Output:
	// A is not a Knight & A is not a Sorcerer & B is not a Knight & B is not a Sorcerer & C is a Knight & C is a Sorcerer
}

func ExampleSimplify_smullyan3() {

	// The hero finds two guies 'a' and 'b'
	A_is_a_knight := ID("A is a Knight")
	B_is_a_knight := ID("B is a Knight")
	// exactly one of them is a Sorcerer
	A_is_Sorcerer := ID("A is a Sorcerer")
	B_is_Sorcerer := ID("B is a Sorcerer")
	Fact1 := Exactly(1, A_is_Sorcerer, B_is_Sorcerer)

	// he asks 'a': is the sorcerer a knight?

	Sorcerer_is_a_knight := ID("The Sorcerer is a Knight")

	// Knowing that the sorcerer is a knight means that beeing a sorcerer => beeing a knight.
	Fact2 := Impl(Sorcerer_is_a_knight, Impl(A_is_Sorcerer, A_is_a_knight))
	Fact3 := Impl(Sorcerer_is_a_knight, Impl(B_is_Sorcerer, B_is_a_knight))
	Fact4 := Impl(Not(Sorcerer_is_a_knight), Impl(A_is_Sorcerer, Not(A_is_a_knight)))
	Fact5 := Impl(Not(Sorcerer_is_a_knight), Impl(B_is_Sorcerer, Not(B_is_a_knight)))

	Facts := And(Fact1, Fact2, Fact3, Fact4, Fact5)

	// if A answered yes
	Hypothesis1 := Eq(A_is_a_knight, Sorcerer_is_a_knight)
	// what can be deduced is what is always true, i.e y, so that x = y AND something else
	Deduction1, _ := Factor(And(Facts, Hypothesis1))

	// if A answered no
	Hypothesis2 := Eq(A_is_a_knight, Not(Sorcerer_is_a_knight))
	Deduction2, _ := Factor(And(Facts, Hypothesis2))
	// If nothing can be deduced, then Deduction is "True" (the least thing that is always true)
	//So we can do

	fmt.Println(Simplify(And(Deduction1, Deduction2)))
	//Output:
	// A is not a Sorcerer & B is a Sorcerer
}

func ExampleSimplify_smullyan4() {
	// A Knight is someone that always tells the truth.
	// A Knaves is someone that always tells a lie.

	// there are two persons accused 'b' and 'c'.
	B_is_a_knight := ID("B is a Knight")
	C_is_a_knight := ID("C is a Knight")

	// 'b' pretend that 'c' said "I lied, yesterday"
	// there are two statement in this sentences
	// I spoke yesterday, and it was a lie
	//
	C_Spoke := ID("C Spoke")
	C_Lied := ID("C Lied")
	Story := And(C_Spoke, C_Lied)

	Fact1 := Eq(B_is_a_knight, Eq(C_is_a_knight, Story))

	Fact2 := Eq(C_Lied, Not(C_is_a_knight))
	deduction, _ := Factor(And(Fact1, Fact2))
	fmt.Println(deduction)
	// The only thing that is certain is "True". There is not relevant information here

	//Output:
	// true
}
