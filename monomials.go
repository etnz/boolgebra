package boolgebra

// extract all monomials of an expression
//
// monomials are Literals, Identifiers, and Not( <monome>)
//
// returns a slice of all the monomials, without repetition
func monomials(X Expression) []Expression {
	// M will begin to contain all occurences of monomials, no order and duplication
	if X == nil {
		return nil
	}
	M := make([]Expression, 0)
	switch X.Type() {
	case TypeAnd, TypeOr:
		for _, sub := range X.Elems() {
			for _, y := range monomials(sub) {
				if !contains(y, M) {
					M = append(M, y)

				}
			}
		}
	default: // it's already monomial
		if !contains(X, M) {
			M = append(M, X)

		}

	}

	return M
}

func contains(x Expression, X []Expression) bool {
	for _, y := range X {
		if Equals(x, y) {
			return true
		}
	}
	return false
}
