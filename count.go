package boolgebra

// combine terms into a big And() based on bits in i:
//
//    if jth bit is 1, terms[j] is used
//    if jth bit is 0, Not(terms[j]) is used
//
func combine(i uint64, terms ...Expression) Expression {

	L := len(terms)
	mask := uint64(1) << uint(L)
	//we start on the first item of "terms" that is more natural to be
	// the heavy bit, so we start with that bit
	sub := make([]Expression, len(terms))
	for j, t := range terms {
		mask >>= 1          // shift right the mask of one bit
		if mask&i != mask { // there is a 0 in the jth pos of i
			t = Not(t)
		}
		sub[j] = t
	}
	return And(sub...)
}

// anyOf return an Or() of all terms combined with 'i' if 'i' is accepted
//
// this is the common function for stuff like "at least one is true", "at most",etc
func anyOf(accept func(i uint64) bool, terms ...Expression) Expression {
	var i, L, combinations uint64
	L = uint64(len(terms))
	combinations = 1 << L

	subs := make([]Expression, 0, L)
	for i = 0; i < combinations; i++ {
		if accept(i) {
			subs = append(subs, combine(i, terms...))
		}
	}
	return Or(subs...)
}

// AtMost retuns an expression that is true if and only if at most 'i' terms are True.
func AtMost(i int, terms ...Expression) Expression {
	return anyOf(func(j uint64) bool {
		return popcount(j) <= byte(i)
	}, terms...)

}

// AtLeast retuns an expression that is true if and only if at least 'i' terms are True.
func AtLeast(i int, terms ...Expression) Expression {
	return anyOf(func(j uint64) bool {
		return popcount(j) >= byte(i)
	}, terms...)

}

// Exactly retuns an expression that is true if and only if exactly 'i' terms are True.
func Exactly(i int, terms ...Expression) Expression {
	return anyOf(func(j uint64) bool {
		return popcount(j) == byte(i)
	}, terms...)

}

// one important tool is to count bit == 1 in a number
// for efficiency purpose we don't want to scan every single bit, instead we want to scan each bytes, and have
// remember the bit count for each possible bytes (there are 256 only !)

// bytePopcounts act as a dict for bytes -> popcount
// used to count population of larger ints
var bytePopcounts = [1 << 8]byte{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7, 4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8}

// to compute the magic const above, you can use the following code,
// func init() {
// 	for i := range bytePopcounts {
// 		var n byte
// 		// brute force popcount each byte in a
// 		for x := i; x != 0; x >>= 1 {
// 			if x&1 != 0 {
// 				n++
// 			}
// 		}
// 		bytePopcounts[i] = n
// 	}
// }

// popcount any uint64: return the number of one set
func popcount(x uint64) (n byte) {
	return bytePopcounts[byte(x>>(0*8))] +
		bytePopcounts[byte(x>>(1*8))] +
		bytePopcounts[byte(x>>(2*8))] +
		bytePopcounts[byte(x>>(3*8))] +
		bytePopcounts[byte(x>>(4*8))] +
		bytePopcounts[byte(x>>(5*8))] +
		bytePopcounts[byte(x>>(6*8))] +
		bytePopcounts[byte(x>>(7*8))]
}
