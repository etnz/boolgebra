package boolgebra

// Factor computes 'y' so that
//
//    X = And(Y, z)
//
// y beeing as big as possible
func Factor(X Expression) (Y Expression) {
	X = Simplify(X)
	if X.Type() != TypeOr {
		return X
	}

	// get all elems of X
	subs := X.Elems()

	// M is all the monomials of X, without repetition
	M := monomials(X)
	// we are going to look for each monomial if it is in all subs
	var result Expression = True

monomialsLoop:
	for _, m := range M {

		for _, sub := range subs {
			if !contains(m, sub.Elems()) {
				continue monomialsLoop // exit this monomial handling
			}
		}
		// m, is present in all subs, add it to the result
		result = iand(result, m)
	}
	return Simplify(result)
}
