package boolgebra

// Equals return true if x and y are Equals.
//
// This is a [syntactic equality](https://en.wikipedia.org/wiki/Symbolic_computation#Equality).
//
// Different types always return false, otherwise
//
//     Literals : their bool Val is compared
//     Ids      : their string ID is compared
//     Nots     : negated expressions are compared
//     Ands, Ors: Elems() are compared like sets (no order)
func Equals(x, y Expression) bool {
	//protect against nil:
	if x == nil && y == nil {
		//both are nil
		return true
	}
	if x == nil || y == nil {
		// only one is nil (both cannot be nil, it has been tested above)
		return false
	}

	if x.Type() != y.Type() {
		// compare actual types, to be able to deep compare after
		return false
	}

	// now x and y share the same type:
	switch x.Type() {
	case TypeLiteral:
		return x.Val() == y.Val()

	case TypeIdentifier:
		return String(x) == String(y)

	default:
		n, m := x.Elems(), y.Elems()

		if len(n) != len(m) {
			// if they can have the same length
			// they must be different!
			return false
		}
		//
		for i := range n {
			found := -1
			for j := range m {
				if Equals(n[i], m[j]) {
					found = j
					break
				}
			}
			if found < 0 {
				return false
			}
			m = remove(m, found)
		}
		return true
	}
}
