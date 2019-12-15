package boole

import "fmt"

//testcases from Raymond Smullyan "Satan, Cantor and Infinity"

func ExampleSimplify_smullyan1() {
	// A Knight is someone that always tells the truth.
	// A Knaves is someone that always tells a lie.

	// The hero meets three guys 'a','b', 'c', each one can be a Knight or a Knaves
	// let's define three properties
	A_is_a_knight := ID("A is a Knight")
	B_is_a_knight := ID("B is a Knight")
	C_is_a_knight := ID("C is a Knight")

	// Our hero asked 'a': "are 'b' and 'c' knights ?"
	Q1 := And(B_is_a_knight, C_is_a_knight)

	//'a' answered yes.

	// like always with Knights and Knaves,
	// if 'a' is a knight, then Q1 is true obviouly,
	// if not 'a' is a knight then Q1 is not true
	//
	//     A_is_a_knight     Q1
	//     true              true
	//     false             false
	//
	// therefore A_is_a_knight  ==   Q1
	Fact1 := Eq(A_is_a_knight, Q1)

	// but 'a' also said that 'b' was a Knaves:
	Fact2 := Eq(A_is_a_knight, Not(B_is_a_knight))

	fmt.Println(Simplify(And(Fact1, Fact2)))
	//Output:
	// !"A is a Knight" & "B is a Knight" & !"C is a Knight"

}

func ExampleSimplify_smullyan2() {
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
	// !"A is a Knight" & !"A is a Sorcerer" & !"B is a Knight" & !"B is a Sorcerer" & "C is a Knight" & "C is a Sorcerer"

}

func ExampleSimplify_smullyan3() {
	// A Knight is someone that always tells the truth.
	// A Knaves is someone that always tells a lie.

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
	Deduction1 := Factor(And(Facts, Hypothesis1))

	// if A answered no
	Hypothesis2 := Eq(A_is_a_knight, Not(Sorcerer_is_a_knight))
	Deduction2 := Factor(And(Facts, Hypothesis2))
	// If nothing can be deduced, then Deduction is "True" (the least thing that is always true)
	//So we can do

	fmt.Println(Simplify(And(Deduction1, Deduction2)))
	//Output:
	// !"A is a Sorcerer" & "B is a Sorcerer"
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
	fmt.Println(Factor(And(Fact1, Fact2)))
	// The only thing that is certain is "True". There is not relevant information here

	//Output:
	//True
}

func ExampleSimplify_smullyan5() {
	// A Knight is someone that always tells the truth.
	// A Knaves is someone that always tells a lie.

	// 'g' is accused by 'c' of stealing a watch
	G_is_a_knight := ID("G is a Knight")
	S := ID("G stole the watch")

	// 'g' said that he pretended that he didn't steal the watch
	Fact1 := Eq(G_is_a_knight, Eq(G_is_a_knight, Not(S)))

	// 'c' asks a second question:
	// did you pretend that you have stolen the watch ?
	Q := Eq(G_is_a_knight, S)
	Hypothesis1 := Eq(G_is_a_knight, Q)
	Hypothesis2 := Eq(G_is_a_knight, Not(Q))
	//
	Deduction1 := Simplify(And(Fact1, Hypothesis1))
	Deduction2 := Simplify(And(Fact1, Hypothesis2))

	fmt.Println(Deduction1)
	fmt.Println(Deduction2)

	//Output:
	// False
	// !"G stole the watch"
}
