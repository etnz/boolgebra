package boole

//simplify reduces 'x' using Quine-McCluskey algorithm
//func simplify(x expression) expression{
// QM algo goes: two phases
// - reduce minterms for instance ABCD or ABC'D => ABD, the expression is not yet minimal, but it has some good enough properties
// - minimize the minterms, some are not not prime, Petrick's method, makes it possible to find the minimum by creating another expression to simply reduce ( as defined above)

//	red := reduce(x) // apply the first level of reduction

// each minterm in red is given an ID
// then Petricks builds a new expression P out of it that just need to be reduced
// then pick P-minterms that have the lowest number of ID.
// they are all good candidates, but we can optimize one step further, once expanded to red-minterms, use the one with the lowest number of red-ID
//}

//reduce combine together all minterms of x into prime implicants
//
// this is not the final simplification since a subset of the prime implicants is enough to generate an equivalent expression. This final stage is not part of this function,
// therefore you'll get only prime implicants, but not a minimal subset of those.
func reduce(x expression) expression {
	// to reduce we need to cluster minterms of x into number of non neg ID

	var cluster [][]minterm // index minterm by their 1s. store them in a slice
	var primes []minterm    // also keep primes all together

	// fill the first cluster
	cluster = make([][]minterm, 1+len(x.IDs()))
	for _, m := range x {
		ones := m.PosLen()
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

func appendunique(set []minterm, m minterm) []minterm {
	for _, x := range set {
		if x.Equals(m) { // nothing to append
			return set
		}
	}
	return append(set, m)

}
