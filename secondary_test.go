package boolgebra

import "testing"

func Test_truthTables2(t *testing.T) {

	T := Lit(true)
	F := Lit(false)

	//XOr
	truthTester(t, "Xor(F,F)", Xor(F, F), false)
	truthTester(t, "Xor(F,T)", Xor(F, T), true)
	truthTester(t, "Xor(T,F)", Xor(T, F), true)
	truthTester(t, "Xor(T,T)", Xor(T, T), false)

}
