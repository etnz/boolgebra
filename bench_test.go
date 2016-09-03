package boolgebra

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
	x := True
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify__False(b *testing.B) {
	x := False
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
	x := Not(True)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd__True_A(b *testing.B) {
	x := And(True, ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A__True(b *testing.B) {
	x := And(ID("A"), True)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd__False_A(b *testing.B) {
	x := And(False, ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_A__False(b *testing.B) {
	x := And(ID("A"), False)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr__True_A(b *testing.B) {
	x := Or(True, ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A__True(b *testing.B) {
	x := Or(ID("A"), True)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr__False_A(b *testing.B) {
	x := Or(False, ID("A"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_A__False(b *testing.B) {
	x := Or(ID("A"), False)
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
	x := Or(And(False, False), And(True, True))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__True__True_TypeAnd__False__False(b *testing.B) {
	x := Or(And(True, True), And(False, False))
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
	x := Or(And(False, Not(ID("A"))), And(Not(ID("A")), Not(ID("A"))), And(True, ID("A"), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A_TypeNot_A_TypeAnd__False_TypeNot_A_TypeAnd_A__True_A(b *testing.B) {
	x := Or(And(Not(ID("A")), Not(ID("A"))), And(False, Not(ID("A"))), And(ID("A"), True, ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__True__True_TypeAnd_TypeNot_A__True_TypeAnd__False_A__False(b *testing.B) {
	x := Or(And(True, True), And(Not(ID("A")), True), And(False, ID("A"), False))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A__True_TypeAnd__True__True_TypeAnd_A__False__False(b *testing.B) {
	x := Or(And(Not(ID("A")), True), And(True, True), And(ID("A"), False, False))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__False_TypeNot_A__False_TypeAnd__True__True_TypeAnd_A__True(b *testing.B) {
	x := Or(And(False, Not(ID("A")), False), And(True, True), And(ID("A"), True))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A__False__False_TypeAnd_A__True_TypeAnd__True__True(b *testing.B) {
	x := Or(And(Not(ID("A")), False, False), And(ID("A"), True), And(True, True))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd__True_TypeNot_A_TypeNot_A_TypeAnd__False_A_TypeAnd_A_A(b *testing.B) {
	x := Or(And(True, Not(ID("A")), Not(ID("A"))), And(False, ID("A")), And(ID("A"), ID("A")))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeOr_TypeAnd_TypeNot_A__True_TypeNot_A_TypeAnd_A_A_TypeAnd__False_A(b *testing.B) {
	x := Or(And(Not(ID("A")), True, Not(ID("A"))), And(ID("A"), ID("A")), And(False, ID("A")))
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
	x := Or(And(And(And(ID("B"), Not(ID("A"))), Not(ID("B"))), False), And(And(And(ID("A"), Not(ID("A"))), Not(ID("B"))), False), And(And(Not(ID("B")), Not(ID("A"))), True), And(ID("A"), True), And(ID("B"), True))
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
func BenchmarkSimplify_TypeAnd_TypeOr_TypeAnd_TypeNot_AisaSorcerer_TypeNot_BisaSorcerer_CisaSorcerer_TypeAnd_TypeNot_AisaSorcerer_BisaSorcerer_TypeNot_CisaSorcerer_TypeAnd_AisaSorcerer_TypeNot_BisaSorcerer_TypeNot_CisaSorcerer_TypeOr_TypeAnd_TypeNot_AisaKnight_TypeNot_AisaSorcerer_TypeAnd_AisaKnight_AisaSorcerer_TypeOr_TypeAnd_TypeNot_BisaKnight_TypeNot_BisaSorcerer_TypeAnd_BisaKnight_BisaSorcerer_TypeOr_TypeAnd_TypeNot_CisaKnight_TypeNot_TypeOr_TypeAnd_TypeNot_AisaKnight_TypeNot_BisaKnight_TypeNot_CisaKnight_TypeAnd_TypeNot_AisaKnight_TypeNot_BisaKnight_CisaKnight_TypeAnd_TypeNot_AisaKnight_BisaKnight_TypeNot_CisaKnight_TypeAnd_AisaKnight_TypeNot_BisaKnight_TypeNot_CisaKnight_TypeAnd_CisaKnight_TypeOr_TypeAnd_TypeNot_AisaKnight_TypeNot_BisaKnight_TypeNot_CisaKnight_TypeAnd_TypeNot_AisaKnight_TypeNot_BisaKnight_CisaKnight_TypeAnd_TypeNot_AisaKnight_BisaKnight_TypeNot_CisaKnight_TypeAnd_AisaKnight_TypeNot_BisaKnight_TypeNot_CisaKnight(b *testing.B) {
	x := And(Or(And(Not(ID("A is a Sorcerer")), Not(ID("B is a Sorcerer")), ID("C is a Sorcerer")), And(Not(ID("A is a Sorcerer")), ID("B is a Sorcerer"), Not(ID("C is a Sorcerer"))), And(ID("A is a Sorcerer"), Not(ID("B is a Sorcerer")), Not(ID("C is a Sorcerer")))), Or(And(Not(ID("A is a Knight")), Not(ID("A is a Sorcerer"))), And(ID("A is a Knight"), ID("A is a Sorcerer"))), Or(And(Not(ID("B is a Knight")), Not(ID("B is a Sorcerer"))), And(ID("B is a Knight"), ID("B is a Sorcerer"))), Or(And(Not(ID("C is a Knight")), Not(Or(And(Not(ID("A is a Knight")), Not(ID("B is a Knight")), Not(ID("C is a Knight"))), And(Not(ID("A is a Knight")), Not(ID("B is a Knight")), ID("C is a Knight")), And(Not(ID("A is a Knight")), ID("B is a Knight"), Not(ID("C is a Knight"))), And(ID("A is a Knight"), Not(ID("B is a Knight")), Not(ID("C is a Knight")))))), And(ID("C is a Knight"), Or(And(Not(ID("A is a Knight")), Not(ID("B is a Knight")), Not(ID("C is a Knight"))), And(Not(ID("A is a Knight")), Not(ID("B is a Knight")), ID("C is a Knight")), And(Not(ID("A is a Knight")), ID("B is a Knight"), Not(ID("C is a Knight"))), And(ID("A is a Knight"), Not(ID("B is a Knight")), Not(ID("C is a Knight")))))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_TypeAnd_TypeOr_TypeAnd_TypeNot_AisaSorcerer_BisaSorcerer_TypeAnd_AisaSorcerer_TypeNot_BisaSorcerer_TypeOr_TypeNot_TheSorcererisaKnight_TypeOr_TypeNot_AisaSorcerer_AisaKnight_TypeOr_TypeNot_TheSorcererisaKnight_TypeOr_TypeNot_BisaSorcerer_BisaKnight_TypeOr_TypeNot_TypeNot_TheSorcererisaKnight_TypeOr_TypeNot_AisaSorcerer_TypeNot_AisaKnight_TypeOr_TypeNot_TypeNot_TheSorcererisaKnight_TypeOr_TypeNot_BisaSorcerer_TypeNot_BisaKnight_TypeOr_TypeAnd_TypeNot_AisaKnight_TypeNot_TheSorcererisaKnight_TypeAnd_AisaKnight_TheSorcererisaKnight(b *testing.B) {
	x := And(And(Or(And(Not(ID("A is a Sorcerer")), ID("B is a Sorcerer")), And(ID("A is a Sorcerer"), Not(ID("B is a Sorcerer")))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("A is a Sorcerer")), ID("A is a Knight"))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("B is a Sorcerer")), ID("B is a Knight"))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("A is a Sorcerer")), Not(ID("A is a Knight")))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("B is a Sorcerer")), Not(ID("B is a Knight"))))), Or(And(Not(ID("A is a Knight")), Not(ID("The Sorcerer is a Knight"))), And(ID("A is a Knight"), ID("The Sorcerer is a Knight"))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_TypeAnd_TypeOr_TypeAnd_TypeNot_AisaSorcerer_BisaSorcerer_TypeAnd_AisaSorcerer_TypeNot_BisaSorcerer_TypeOr_TypeNot_TheSorcererisaKnight_TypeOr_TypeNot_AisaSorcerer_AisaKnight_TypeOr_TypeNot_TheSorcererisaKnight_TypeOr_TypeNot_BisaSorcerer_BisaKnight_TypeOr_TypeNot_TypeNot_TheSorcererisaKnight_TypeOr_TypeNot_AisaSorcerer_TypeNot_AisaKnight_TypeOr_TypeNot_TypeNot_TheSorcererisaKnight_TypeOr_TypeNot_BisaSorcerer_TypeNot_BisaKnight_TypeOr_TypeAnd_TypeNot_AisaKnight_TypeNot_TypeNot_TheSorcererisaKnight_TypeAnd_AisaKnight_TypeNot_TheSorcererisaKnight(b *testing.B) {
	x := And(And(Or(And(Not(ID("A is a Sorcerer")), ID("B is a Sorcerer")), And(ID("A is a Sorcerer"), Not(ID("B is a Sorcerer")))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("A is a Sorcerer")), ID("A is a Knight"))), Or(Not(ID("The Sorcerer is a Knight")), Or(Not(ID("B is a Sorcerer")), ID("B is a Knight"))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("A is a Sorcerer")), Not(ID("A is a Knight")))), Or(Not(Not(ID("The Sorcerer is a Knight"))), Or(Not(ID("B is a Sorcerer")), Not(ID("B is a Knight"))))), Or(And(Not(ID("A is a Knight")), Not(Not(ID("The Sorcerer is a Knight")))), And(ID("A is a Knight"), Not(ID("The Sorcerer is a Knight")))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd__True_TypeNot_AisaSorcerer_BisaSorcerer(b *testing.B) {
	x := And(True, Not(ID("A is a Sorcerer")), ID("B is a Sorcerer"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd__True_TypeAnd_BisaSorcerer_TypeNot_AisaSorcerer(b *testing.B) {
	x := And(True, And(ID("B is a Sorcerer"), Not(ID("A is a Sorcerer"))))
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
func BenchmarkSimplify_TypeAnd_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_TypeNot_Gstolethewatch_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_Gstolethewatch_TypeAnd_GisaKnight_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_Gstolethewatch(b *testing.B) {
	x := And(Or(And(Not(ID("G is a Knight")), Not(Or(And(Not(ID("G is a Knight")), Not(Not(ID("G stole the watch")))), And(ID("G is a Knight"), Not(ID("G stole the watch")))))), And(ID("G is a Knight"), Or(And(Not(ID("G is a Knight")), Not(Not(ID("G stole the watch")))), And(ID("G is a Knight"), Not(ID("G stole the watch")))))), Or(And(Not(ID("G is a Knight")), Not(Or(And(Not(ID("G is a Knight")), Not(ID("G stole the watch"))), And(ID("G is a Knight"), ID("G stole the watch"))))), And(ID("G is a Knight"), Or(And(Not(ID("G is a Knight")), Not(ID("G stole the watch"))), And(ID("G is a Knight"), ID("G stole the watch"))))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
func BenchmarkSimplify_TypeAnd_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_TypeNot_Gstolethewatch_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_TypeNot_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_Gstolethewatch_TypeAnd_GisaKnight_TypeNot_TypeOr_TypeAnd_TypeNot_GisaKnight_TypeNot_Gstolethewatch_TypeAnd_GisaKnight_Gstolethewatch(b *testing.B) {
	x := And(Or(And(Not(ID("G is a Knight")), Not(Or(And(Not(ID("G is a Knight")), Not(Not(ID("G stole the watch")))), And(ID("G is a Knight"), Not(ID("G stole the watch")))))), And(ID("G is a Knight"), Or(And(Not(ID("G is a Knight")), Not(Not(ID("G stole the watch")))), And(ID("G is a Knight"), Not(ID("G stole the watch")))))), Or(And(Not(ID("G is a Knight")), Not(Not(Or(And(Not(ID("G is a Knight")), Not(ID("G stole the watch"))), And(ID("G is a Knight"), ID("G stole the watch")))))), And(ID("G is a Knight"), Not(Or(And(Not(ID("G is a Knight")), Not(ID("G stole the watch"))), And(ID("G is a Knight"), ID("G stole the watch")))))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simplify(x)
	}
}
