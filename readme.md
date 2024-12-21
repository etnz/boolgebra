[![GoDoc](https://godoc.org/github.com/etnz/boolgebra?status.svg)](https://godoc.org/github.com/etnz/boolgebra)


Package boolgebra provide basic boolean algebra operations.

It is possible to compare, expand, reduce, parse boolean expressions.

It makes it possible to resolve logic puzzles, like Smullyan's, or logic grid.

Because I cannot solve those puzzles without trying to get the computer doing it for me.

Let's solve this Smullyan's [problem](https://en.wikipedia.org/wiki/Knights_and_Knaves#Both_knaves):

Alice and Bob are residents of the island of knights and knaves, where knights always tell the truth and 
knaves always tell a lie.

Alice says, "We are both knaves”.

Can you tell who is a knight, and who is knave?

In the island of knights and knaves, when you know for sure that "Alice says it is raining"
you cannot tell for sure if it rains, but you can be certain that if Alice is a knight, then she is telling the truth
then it is raining, therefore `Alice is knight => it is raining`. 
Also, if it is raining, then Alice told the truth, then she must be a knight. `it is raining => Alice is a knight`. 
Therefore, when you know for sure that "Alice says it is raining" you can write for sure that `Alice is a knight <=> it is raining`.

Solving Smullyan's puzzles with boolgebra is all about writing *correctly* what you know for certain in the puzzle.

back to our problem, knowing for sure that Alice said "We are both knaves”, you can write, for sure 
that `Alice is a knight <=> Alice is not a knight & Bob is not a knight`

using boolgebra you can compute the solution by just doing:

```go
	problem, _ := Parse(`Alice is a Knight <=> Alice is not a Knight & Bob is not a Knight`)
	fmt.Println(Simplify(problem))
	// Output: Alice is not a Knight & Bob is a Knight
```