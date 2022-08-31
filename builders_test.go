package boolgebra

import (
	"strconv"
	"testing"
)

// BenchmarkAnd using And as a accumulator function is counter productive.
// like in for `{ res = And(res, x[i]) }`
//
// This benchmark is to show that, in comparision with BenchmarkTermBuilder
func BenchmarkAnd(b *testing.B) {

	res := Lit[string](true) // neutral for And
	// prepare 1000 different ids
	M := 1000
	var ids []string
	for i := 0; i < M; i++ {
		ids = append(ids, "parameter_"+strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = And(res, ID(ids[i%M]))
	}
}

// BenchmarkTermBuilder does the same as BenchmarkAnd but using the TermBuilder.
//
// TermBuilder is much more faster than using And over and over. Indeed, the API uses
// immutable object that require a lot of memory allocation. The builder save those allocations
func BenchmarkTermBuilder(b *testing.B) {

	// prepare 1000 different ids
	M := 1000
	var ids []string
	for i := 0; i < M; i++ {
		ids = append(ids, "parameter_"+strconv.Itoa(i))
	}
	var t TermBuilder[string]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.And(ids[i%M], true)
	}
}
