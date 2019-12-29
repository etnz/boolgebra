package boolgebra

// Quine-McCluskey is an algorithm to simplify a sum of prod.

//reduce combine together all minterms of x into prime implicants
func reduce(x expression) expression {
	// to reduce we need to cluster minterms of x into number of non neg ID

	var cluster [][]minterm // index minterm by their 1s. store them in a slice
	var primes []minterm    // also keep primes all together

	// fill the first cluster
	cluster = make([][]minterm, 1+len(x.IDs()))
	for _, m := range x {
		ones := positives(m)
		cluster[ones] = append(cluster[ones], m)
	}

	emptycluster := false // combinations in the current cluster are moved to the next cluster. we repeat until the "next" cluster is empty
	for !emptycluster {
		// the next cluster will become the current one soon, so we already set the bool to true, because we start with an empty one
		emptycluster = true // we start with an empty next one, let see if it get filled
		next := make([][]minterm, len(cluster))

		// attempt all possible combinations.
		// a minterm with n Positive IDs, can only be combined with another one with n or n+1 ( or n-1 but combination is symetric so we don't care)
		// that is the reason why we built the cluster in the first place
		// so we loop over all the minterms couples that *could* be mixed
		for i, ms := range cluster {
			// loop over all cluster minterms
			for mi, m := range ms {
				//m is left candidate
				misprime := true // by default m is assumed to be prime, but this will be tested

				// check m against peers in the same cluster
				for mj, r := range ms {
					if mj != mi {
						// except self combination check it
						if c, ok := combine(m, r); ok {
							misprime = false // it has been combined, so it's not a prime anymore
							// append it in the next cluster
							emptycluster = false         // mark the next as non empty
							next[i] = append(next[i], c) // we know that the number of 1s cannot be reduced
							// proof: every ID can have only three possible values 1, 0, or _ (don't care)
							// to be combined they have to bo identical on all bits but one.
							// number of ones on those "identical" bits are equals obviously, therefore
							// the number of 1 on the bit that is different must be equal too, so this cannot be a 1.
							// so this can only b a 0 and a _ reduced to be _, so the total amount of 1 is unchanged

						}
					}
				}

				// check against minterms in the next cluster
				if i+1 < len(cluster) {

					for _, r := range cluster[i+1] {
						if c, ok := combine(m, r); ok {
							misprime = false // it has been combined, so it's not a prime anymore
							// append it in the next cluster
							emptycluster = false         // mark the next as non empty
							next[i] = append(next[i], c) // we know that the number of 1s cannot be reduced
							// proof: every ID can have only three possible values 1, 0, or _ (don't care)
							// to be combined they have to bo identical on all bits but one.
							// number of ones on those "identical" bits are equals obviously
							// r has an extra 1 compared to m ( by definition of the cluster[i+1]
							// therefore on this bit, r has a 1, and m has either a 0 or a _, so on this bit, c is now _
							// and has the exact same amount of 1s as m
						}
					}
				}

				// check against minterms in the previous cluster. those checks have already been made, but this is *just* to
				// find out if m is still a prime. Indeed if it has been combined with a lower minterm, it is not a prime one
				if misprime && i-1 >= 0 {
					// so far m has been tested against self, and next in cluster, but it might have be used previously ( i.e cluster - 1
					// therefore we are testing backward to see if it was used previously
					for _, r := range cluster[i-1] {
						if _, ok := combine(m, r); ok {
							misprime = false // it has been combined previously, so it's not a prime anymore
							// no other side effect, because they have already been made when testing the
						}
					}
				}

				if misprime {

					primes = appendunique(primes, m)
				}
			}

		}
		// done with filling the next cluster
		cluster = next
	}
	//done, we now have all the prime implicant
	return expression(primes)
}

// appenunique behave like 'append' except for items in 'terms' that are present in 'set': they
// are not appended in this case.
func appendunique(set []minterm, terms ...minterm) []minterm {
termsloop:
	for _, m := range terms {
		for _, x := range set {
			if equals(x, m) {
				// noting to append
				continue termsloop
			}
		}
		set = append(set, m)
	}
	return set
}

// combine computes c, if possible, that combines x and y
//
// x and y must be identical but on exactly one identifier.
//
// the combined is then then intersection of x and y.
func combine(x, y minterm) (c minterm, ok bool) {
	// alg: find out the one and only one difference between x,y
	// so scan for differences and count.
	var d string // the identifier that is different (if diffs == 1))
	diffs := 0   // number of differences
	for k, v := range x {
		w, exists := y[k]
		if !exists || v != w {
			// this one is different, store it
			d = k
			diffs++
			if diffs > 1 {
				return c, false
			}
		}
	}
	// same goes backward too ( to check for missing in x only

	for k := range y {
		_, exists := x[k]
		if !exists {
			// this one is different, store it
			d = k
			diffs++
			if diffs > 1 {
				return c, false
			}
		}
	}
	if diffs != 1 {
		return c, false
	}
	// we hit the one difference !
	// build c accordingly then
	// x and y are guaranteed to be identical but on 'd'
	// so copy x but 'd'
	c = make(minterm)
	for k, v := range x {
		if k != d {
			c[k] = v
		}
	}
	return c, true

}

// equals return true if and only if m and n are both minterm, then they are semantically equals
func equals(m, n minterm) bool {
	if len(m) != len(n) {
		return false
	}
	for k, v := range m {
		if w, exists := n[k]; !exists || v != w {
			return false
		}
	}
	return true
}

// positives returns the number of positive identifiers
func positives(m minterm) int {
	count := 0
	for _, v := range m {
		if v {
			count++
		}
	}
	return count
}

//inter computes the intersection of x inter  y
func inter(x, y minterm) minterm {
	res := make(minterm)
	for k, v := range x {
		if w, exists := y[k]; exists && v == w {
			res[k] = v
		}
	}
	return res

}

//div computes x/y i.e z so that And(z,y) = x
// can be seen as x removed from items in y
func div(x, y minterm) minterm {
	res := make(minterm)
	for k, v := range x {
		if w, exists := y[k]; !exists || v != w {
			res[k] = v
		}
	}
	return res
}
