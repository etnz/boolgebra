package boole

import "testing"

func BenchmarkSimplify_Bigone(b *testing.B) {
	x := And(And(And(Or(And(Not(ID("A is a Sorcerer")), ID("B is a Sorcerer")), And(ID("A is a Sorcerer"), Not(ID("B is a Sorcerer")))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("A is a Sorcerer")), ID("A is a Knight"))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("B is a Sorcerer")), ID("B is a Knight"))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("A is a Sorcerer")), Not(ID("A is a Knight")))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("B is a Sorcerer")), Not(ID("B is a Knight"))))), Or(And(Not(ID("A is a Knight")), Not(ID("The Sorcerer is a Knight"))), And(ID("A is a Knight"), ID("The Sorcerer is a Knight")))), And(And(Or(And(Not(ID("A is a Sorcerer")), ID("B is a Sorcerer")), And(ID("A is a Sorcerer"), Not(ID("B is a Sorcerer")))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("A is a Sorcerer")), ID("A is a Knight"))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("B is a Sorcerer")), ID("B is a Knight"))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("A is a Sorcerer")), Not(ID("A is a Knight")))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("B is a Sorcerer")), Not(ID("B is a Knight"))))), Or(And(Not(ID("A is a Knight")), Not(Not(ID("The Sorcerer is a Knight")))), And(ID("A is a Knight"), Not(ID("The Sorcerer is a Knight"))))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}

// the following benchmrk were generated.
func BenchmarkSimplify__True(b *testing.B) {
	x := Lit(true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify__False(b *testing.B) {
	x := Lit(false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_A(b *testing.B) {
	x := ID("A")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeNot_A(b *testing.B) {
	x := Not(ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A_B(b *testing.B) {
	x := And(ID("A"), ID("B"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A_B(b *testing.B) {
	x := Or(ID("A"), ID("B"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeNot__True(b *testing.B) {
	x := Not(Lit(true))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd__True_A(b *testing.B) {
	x := And(Lit(true), ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A__True(b *testing.B) {
	x := And(ID("A"), Lit(true))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd__False_A(b *testing.B) {
	x := And(Lit(false), ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A__False(b *testing.B) {
	x := And(ID("A"), Lit(false))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr__True_A(b *testing.B) {
	x := Or(Lit(true), ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A__True(b *testing.B) {
	x := Or(ID("A"), Lit(true))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr__False_A(b *testing.B) {
	x := Or(Lit(false), ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A__False(b *testing.B) {
	x := Or(ID("A"), Lit(false))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A_B_A(b *testing.B) {
	x := Or(ID("A"), ID("B"), ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A_A_A_B(b *testing.B) {
	x := Or(ID("A"), ID("A"), ID("A"), ID("B"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A_B_B(b *testing.B) {
	x := Or(ID("A"), ID("B"), ID("B"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A_B_B_B(b *testing.B) {
	x := Or(ID("A"), ID("B"), ID("B"), ID("B"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A_A_A(b *testing.B) {
	x := Or(ID("A"), ID("A"), ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A_B_A(b *testing.B) {
	x := And(ID("A"), ID("B"), ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A_A_A_B(b *testing.B) {
	x := And(ID("A"), ID("A"), ID("A"), ID("B"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A_B_B(b *testing.B) {
	x := And(ID("A"), ID("B"), ID("B"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A_B_B_B(b *testing.B) {
	x := And(ID("A"), ID("B"), ID("B"), ID("B"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A_A_A(b *testing.B) {
	x := And(ID("A"), ID("A"), ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_A_B_TypeAnd_A_TypeNot_B(b *testing.B) {
	x := Or(And(ID("A"), ID("B")), And(ID("A"), Not(ID("B"))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_B_TypeNot_A_A_B(b *testing.B) {
	x := Or(And(Not(ID("B")), Not(ID("A"))), ID("A"), ID("B"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__False__False_TypeAnd__True__True(b *testing.B) {
	x := Or(And(Lit(false), Lit(false)), And(Lit(true), Lit(true)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__True__True_TypeAnd__False__False(b *testing.B) {
	x := Or(And(Lit(true), Lit(true)), And(Lit(false), Lit(false)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_A_A(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A"))), And(ID("A"), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_A_A_TypeAnd_TypeNot_A_TypeNot_A(b *testing.B) {
	x := Or(And(ID("A"), ID("A")), And(Not(ID("A")), Not(ID("A"))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_B_TypeAnd_TypeNot_B_TypeNot_A_TypeAnd_TypeNot_B_TypeNot_B_TypeAnd_A_B_A_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("B"))), And(Not(ID("B")), Not(ID("A"))), And(Not(ID("B")), Not(ID("B"))), And(ID("A"), ID("B"), ID("A"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_B_TypeNot_A_TypeNot_B_TypeAnd_A_A_TypeAnd_A_B_TypeAnd_B_A_TypeAnd_B_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("B")), Not(ID("A")), Not(ID("B"))), And(ID("A"), ID("A")), And(ID("A"), ID("B")), And(ID("B"), ID("A")), And(ID("B"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__False_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd__True_A_A(b *testing.B) {
	x := Or(And(Lit(false), Not(ID("A"))), And(Not(ID("A")), Not(ID("A"))), And(Lit(true), ID("A"), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd__False_TypeNot_A_TypeAnd_A__True_A(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A"))), And(Lit(false), Not(ID("A"))), And(ID("A"), Lit(true), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__True__True_TypeAnd_TypeNot_A__True_TypeAnd__False_A__False(b *testing.B) {
	x := Or(And(Lit(true), Lit(true)), And(Not(ID("A")), Lit(true)), And(Lit(false), ID("A"), Lit(false)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A__True_TypeAnd__True__True_TypeAnd_A__False__False(b *testing.B) {
	x := Or(And(Not(ID("A")), Lit(true)), And(Lit(true), Lit(true)), And(ID("A"), Lit(false), Lit(false)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__False_TypeNot_A__False_TypeAnd__True__True_TypeAnd_A__True(b *testing.B) {
	x := Or(And(Lit(false), Not(ID("A")), Lit(false)), And(Lit(true), Lit(true)), And(ID("A"), Lit(true)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A__False__False_TypeAnd_A__True_TypeAnd__True__True(b *testing.B) {
	x := Or(And(Not(ID("A")), Lit(false), Lit(false)), And(ID("A"), Lit(true)), And(Lit(true), Lit(true)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__True_TypeNot_A_TypeNot_A_TypeAnd__False_A_TypeAnd_A_A(b *testing.B) {
	x := Or(And(Lit(true), Not(ID("A")), Not(ID("A"))), And(Lit(false), ID("A")), And(ID("A"), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A__True_TypeNot_A_TypeAnd_A_A_TypeAnd__False_A(b *testing.B) {
	x := Or(And(Not(ID("A")), Lit(true), Not(ID("A"))), And(ID("A"), ID("A")), And(Lit(false), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_B_TypeNot_A_TypeNot_A_TypeNot_B_TypeAnd_A_A_TypeAnd_A_B_TypeAnd_B_A_TypeAnd_B_B_TypeAnd_A_A_TypeAnd_A_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("B")), Not(ID("A")), Not(ID("A")), Not(ID("B"))), And(ID("A"), ID("A")), And(ID("A"), ID("B")), And(ID("B"), ID("A")), And(ID("B"), ID("B")), And(ID("A"), ID("A")), And(ID("A"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeNot_A_TypeNot_B_TypeNot_A_TypeNot_B_TypeAnd_A_A_TypeAnd_A_B_TypeAnd_A_A_TypeAnd_A_B_TypeAnd_A_A_TypeAnd_A_B_TypeAnd_B_A_TypeAnd_B_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A")), Not(ID("A")), Not(ID("B")), Not(ID("A")), Not(ID("B"))), And(ID("A"), ID("A")), And(ID("A"), ID("B")), And(ID("A"), ID("A")), And(ID("A"), ID("B")), And(ID("A"), ID("A")), And(ID("A"), ID("B")), And(ID("B"), ID("A")), And(ID("B"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_B_TypeNot_B_TypeNot_A_TypeNot_B_TypeAnd_A_A_TypeAnd_A_B_TypeAnd_B_A_TypeAnd_B_B_TypeAnd_B_A_TypeAnd_B_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("B")), Not(ID("B")), Not(ID("A")), Not(ID("B"))), And(ID("A"), ID("A")), And(ID("A"), ID("B")), And(ID("B"), ID("A")), And(ID("B"), ID("B")), And(ID("B"), ID("A")), And(ID("B"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_B_TypeNot_B_TypeNot_B_TypeNot_A_TypeNot_B_TypeAnd_A_A_TypeAnd_A_B_TypeAnd_B_A_TypeAnd_B_B_TypeAnd_B_A_TypeAnd_B_B_TypeAnd_B_A_TypeAnd_B_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("B")), Not(ID("B")), Not(ID("B")), Not(ID("A")), Not(ID("B"))), And(ID("A"), ID("A")), And(ID("A"), ID("B")), And(ID("B"), ID("A")), And(ID("B"), ID("B")), And(ID("B"), ID("A")), And(ID("B"), ID("B")), And(ID("B"), ID("A")), And(ID("B"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeNot_A_TypeNot_A_TypeAnd_A_A_TypeAnd_A_A_TypeAnd_A_A(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A")), Not(ID("A")), Not(ID("A"))), And(ID("A"), ID("A")), And(ID("A"), ID("A")), And(ID("A"), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_B_TypeAnd_TypeNot_B_TypeNot_A_TypeAnd_TypeNot_B_TypeNot_B_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_B_TypeAnd_A_B_A_A_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("B"))), And(Not(ID("B")), Not(ID("A"))), And(Not(ID("B")), Not(ID("B"))), And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("B"))), And(ID("A"), ID("B"), ID("A"), ID("A"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_B_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_B_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_B_TypeAnd_TypeNot_B_TypeNot_A_TypeAnd_TypeNot_B_TypeNot_B_TypeAnd_A_A_A_B_A_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("B"))), And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("B"))), And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("B"))), And(Not(ID("B")), Not(ID("A"))), And(Not(ID("B")), Not(ID("B"))), And(ID("A"), ID("A"), ID("A"), ID("B"), ID("A"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_B_TypeAnd_TypeNot_B_TypeNot_A_TypeAnd_TypeNot_B_TypeNot_B_TypeAnd_TypeNot_B_TypeNot_A_TypeAnd_TypeNot_B_TypeNot_B_TypeAnd_A_B_B_A_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("B"))), And(Not(ID("B")), Not(ID("A"))), And(Not(ID("B")), Not(ID("B"))), And(Not(ID("B")), Not(ID("A"))), And(Not(ID("B")), Not(ID("B"))), And(ID("A"), ID("B"), ID("B"), ID("A"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_B_TypeAnd_TypeNot_B_TypeNot_A_TypeAnd_TypeNot_B_TypeNot_B_TypeAnd_TypeNot_B_TypeNot_A_TypeAnd_TypeNot_B_TypeNot_B_TypeAnd_TypeNot_B_TypeNot_A_TypeAnd_TypeNot_B_TypeNot_B_TypeAnd_A_B_B_B_A_B(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("B"))), And(Not(ID("B")), Not(ID("A"))), And(Not(ID("B")), Not(ID("B"))), And(Not(ID("B")), Not(ID("A"))), And(Not(ID("B")), Not(ID("B"))), And(Not(ID("B")), Not(ID("A"))), And(Not(ID("B")), Not(ID("B"))), And(ID("A"), ID("B"), ID("B"), ID("B"), ID("A"), ID("B")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd_A_A_A_A(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("A"))), And(Not(ID("A")), Not(ID("A"))), And(ID("A"), ID("A"), ID("A"), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeAnd_TypeNot_A_TypeNot_A_TypeNot_A_TypeAnd_TypeAnd_TypeNot_A_B_TypeNot_A_TypeAnd_TypeAnd_TypeNot_B_TypeNot_A_TypeNot_A_TypeAnd_TypeAnd_TypeNot_B_B_TypeNot_A_TypeAnd_TypeAnd_A_B_A_TypeAnd_TypeAnd_A_TypeNot_B_A(b *testing.B) {
	x := Or(And(And(Not(ID("A")), Not(ID("A"))), Not(ID("A"))), And(And(Not(ID("A")), ID("B")), Not(ID("A"))), And(And(Not(ID("B")), Not(ID("A"))), Not(ID("A"))), And(And(Not(ID("B")), ID("B")), Not(ID("A"))), And(And(ID("A"), ID("B")), ID("A")), And(And(ID("A"), Not(ID("B"))), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeAnd_TypeAnd_B_TypeNot_A_TypeNot_B__False_TypeAnd_TypeAnd_TypeAnd_A_TypeNot_A_TypeNot_B__False_TypeAnd_TypeAnd_TypeNot_B_TypeNot_A__True_TypeAnd_A__True_TypeAnd_B__True(b *testing.B) {
	x := Or(And(And(And(ID("B"), Not(ID("A"))), Not(ID("B"))), Lit(false)), And(And(And(ID("A"), Not(ID("A"))), Not(ID("B"))), Lit(false)), And(And(Not(ID("B")), Not(ID("A"))), Lit(true)), And(ID("A"), Lit(true)), And(ID("B"), Lit(true)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_TypeNot_B_TypeNot_A(b *testing.B) {
	x := And(Not(ID("B")), Not(ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_B(b *testing.B) {
	x := ID("B")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_TypeOr_TypeAnd_TypeNot_AisaKnight_TypeNot_TypeAnd_BisaKnight_CisaKnight_TypeAnd_AisaKnight_TypeAnd_BisaKnight_CisaKnight_TypeOr_TypeAnd_TypeNot_AisaKnight_TypeNot_TypeNot_BisaKnight_TypeAnd_AisaKnight_TypeNot_BisaKnight(b *testing.B) {
	x := And(Or(And(Not(ID("A is a Knight")), Not(And(ID("B is a Knight"), ID("C is a Knight")))), And(ID("A is a Knight"), And(ID("B is a Knight"), ID("C is a Knight")))), Or(And(Not(ID("A is a Knight")), Not(Not(ID("B is a Knight")))), And(ID("A is a Knight"), Not(ID("B is a Knight")))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}

// this one generated an "identifier too long" in go test 1.1
func BenchmarkSimplify_IdentifiertooLong(b *testing.B) {
	x := And(Or(And(Not(ID("A is a Sorcerer")), Not(ID("B is a Sorcerer")), ID("C is a Sorcerer")), And(Not(ID("A is a Sorcerer")), ID("B is a Sorcerer"), Not(ID("C is a Sorcerer"))), And(ID("A is a Sorcerer"), Not(ID("B is a Sorcerer")), Not(ID("C is a Sorcerer")))), Or(And(Not(ID("A is a Knight")), Not(ID("A is a Sorcerer"))), And(ID("A is a Knight"), ID("A is a Sorcerer"))), Or(And(Not(ID("B is a Knight")), Not(ID("B is a Sorcerer"))), And(ID("B is a Knight"), ID("B is a Sorcerer"))), Or(And(Not(ID("C is a Knight")), Not(Or(And(Not(ID("A is a Knight")), Not(ID("B is a Knight")), Not(ID("C is a Knight"))), And(Not(ID("A is a Knight")), Not(ID("B is a Knight")), ID("C is a Knight")), And(Not(ID("A is a Knight")), ID("B is a Knight"), Not(ID("C is a Knight"))), And(ID("A is a Knight"), Not(ID("B is a Knight")), Not(ID("C is a Knight")))))), And(ID("C is a Knight"), Or(And(Not(ID("A is a Knight")), Not(ID("B is a Knight")), Not(ID("C is a Knight"))), And(Not(ID("A is a Knight")), Not(ID("B is a Knight")), ID("C is a Knight")), And(Not(ID("A is a Knight")), ID("B is a Knight"), Not(ID("C is a Knight"))), And(ID("A is a Knight"), Not(ID("B is a Knight")), Not(ID("C is a Knight")))))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}

// this one generated an "identifier too long" in go test 1.1
func BenchmarkSimplify_TooLong2(b *testing.B) {
	x := And(And(Or(And(Not(ID("A is a Sorcerer")), ID("B is a Sorcerer")), And(ID("A is a Sorcerer"), Not(ID("B is a Sorcerer")))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("A is a Sorcerer")), ID("A is a Knight"))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("B is a Sorcerer")), ID("B is a Knight"))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("A is a Sorcerer")), Not(ID("A is a Knight")))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("B is a Sorcerer")), Not(ID("B is a Knight"))))), Or(And(Not(ID("A is a Knight")), Not(ID("The Sorcerer is a Knight"))), And(ID("A is a Knight"), ID("The Sorcerer is a Knight"))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}

// this one generated an "identifier too long" in go test 1.1
func BenchmarkSimplify_TooLong3(b *testing.B) {
	x := And(And(Or(And(Not(ID("A is a Sorcerer")), ID("B is a Sorcerer")), And(ID("A is a Sorcerer"), Not(ID("B is a Sorcerer")))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("A is a Sorcerer")), ID("A is a Knight"))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("B is a Sorcerer")), ID("B is a Knight"))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("A is a Sorcerer")), Not(ID("A is a Knight")))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("B is a Sorcerer")), Not(ID("B is a Knight"))))), Or(And(Not(ID("A is a Knight")), Not(Not(ID("The Sorcerer is a Knight")))), And(ID("A is a Knight"), Not(ID("The Sorcerer is a Knight")))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd__True_TypeNot_AisaSorcerer_BisaSorcerer(b *testing.B) {
	x := And(Lit(true), Not(ID("A is a Sorcerer")), ID("B is a Sorcerer"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd__True_TypeAnd_BisaSorcerer_TypeNot_AisaSorcerer(b *testing.B) {
	x := And(Lit(true), And(ID("B is a Sorcerer"), Not(ID("A is a Sorcerer"))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_TypeOr_TypeAnd_TypeNot_BisaKnight_TypeNot_TypeOr_TypeAnd_TypeNot_CisaKnight_TypeNot_TypeAnd_CSpoke_CLied_TypeAnd_CisaKnight_TypeAnd_CSpoke_CLied_TypeAnd_BisaKnight_TypeOr_TypeAnd_TypeNot_CisaKnight_TypeNot_TypeAnd_CSpoke_CLied_TypeAnd_CisaKnight_TypeAnd_CSpoke_CLied_TypeOr_TypeAnd_TypeNot_CLied_TypeNot_TypeNot_CisaKnight_TypeAnd_CLied_TypeNot_CisaKnight(b *testing.B) {
	x := And(Or(And(Not(ID("B is a Knight")), Not(Or(And(Not(ID("C is a Knight")), Not(And(ID("C Spoke"), ID("C Lied")))), And(ID("C is a Knight"), And(ID("C Spoke"), ID("C Lied")))))), And(ID("B is a Knight"), Or(And(Not(ID("C is a Knight")), Not(And(ID("C Spoke"), ID("C Lied")))), And(ID("C is a Knight"), And(ID("C Spoke"), ID("C Lied")))))), Or(And(Not(ID("C Lied")), Not(Not(ID("C is a Knight")))), And(ID("C Lied"), Not(ID("C is a Knight")))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}

// this one generated an "identifier too long" in go test 1.1
func BenchmarkSimplify_TooLong4(b *testing.B) {
	x := And(Or(And(Not(ID("G is a Knight")), Not(Or(And(Not(ID("G is a Knight")), Not(Not(ID("G stole the watch")))), And(ID("G is a Knight"), Not(ID("G stole the watch")))))), And(ID("G is a Knight"), Or(And(Not(ID("G is a Knight")), Not(Not(ID("G stole the watch")))), And(ID("G is a Knight"), Not(ID("G stole the watch")))))), Or(And(Not(ID("G is a Knight")), Not(Or(And(Not(ID("G is a Knight")), Not(ID("G stole the watch"))), And(ID("G is a Knight"), ID("G stole the watch"))))), And(ID("G is a Knight"), Or(And(Not(ID("G is a Knight")), Not(ID("G stole the watch"))), And(ID("G is a Knight"), ID("G stole the watch"))))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}

// this one generated an "identifier too long" in go test 1.1
func BenchmarkSimplify_TooLong5(b *testing.B) {
	x := And(Or(And(Not(ID("G is a Knight")), Not(Or(And(Not(ID("G is a Knight")), Not(Not(ID("G stole the watch")))), And(ID("G is a Knight"), Not(ID("G stole the watch")))))), And(ID("G is a Knight"), Or(And(Not(ID("G is a Knight")), Not(Not(ID("G stole the watch")))), And(ID("G is a Knight"), Not(ID("G stole the watch")))))), Or(And(Not(ID("G is a Knight")), Not(Not(Or(And(Not(ID("G is a Knight")), Not(ID("G stole the watch"))), And(ID("G is a Knight"), ID("G stole the watch")))))), And(ID("G is a Knight"), Not(Or(And(Not(ID("G is a Knight")), Not(ID("G stole the watch"))), And(ID("G is a Knight"), ID("G stole the watch")))))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
