package boole

// subAnd return the conjunction of the subset of exprs pointed by p.
func subAnd(p []int, exprs []Expr) Expr {
	res := make([]Expr, 0, len(p))
	for _, i := range p {
		res = append(res, exprs[i])
	}
	return And(res...)
}

// subid return the identity permutation
func subid(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = i
	}
	return res
}

func complement(p []int, n int) (c []int) {
	i := 0 // the absolute index, that will sweep all values from 0 to n
	for _, pi := range p {
		for ; i < pi; i++ {
			// the next item to belong to p is pi, i was the last one, so all between belong to the complement
			c = append(c, i)
		}
		i++
	}
	// do til the end
	for ; i < n; i++ {
		// the next item to belong to p is pi, i was the last one, so all between belong to the complement
		c = append(c, i)
	}
	return
}

// nextsubset updates 'p' a 'subset' and return true until it has gone through all
// the possible subsets.
//
// 'p' is a slice of indices from the original list of items (of length 'n')
//
// The first subset to use is [ 0, 1, 2, ...].
func nextsubset(p []int, n int) bool {
	j, k := 0, len(p)

	for ; j < k && p[j] == j; j++ {
	}

	if (k-j)%2 == 0 {
		if j == 0 {
			p[0]--
		} else {
			iset(p, j-2, j-1, j)
		}

	} else {
		pj, pj1 := n+1, n+1
		switch j {
		case k:
		case k - 1:
			pj = p[j]
		default:
			pj = p[j]
			pj1 = p[j+1]
		}

		if pj1 != pj+1 {
			if pj+1 >= n { // detects the end of the revolving door
				p[k-1] = k - 1
				return false
			}
			iset(p, j-1, pj, pj+1)
		} else {
			iset(p, j, j, pj)
		}
	}
	return true
}

//iset is used to do p[i], p[i+1] = vi, vj but skip all the out of bounds cases
func iset(p []int, i, vi, vj int) {
	//there are obviously
	switch {
	case i == -1: //obviously i is outside
		p[0] = vj
	case i < len(p)-1: // i and j are still inside
		// only one is deleted
		p[i], p[i+1] = vi, vj
	case i < len(p):
		p[i] = vi
	}
}
