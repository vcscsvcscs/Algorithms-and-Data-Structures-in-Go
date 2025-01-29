package pattern_matching

func KMP(T, P string) (S []int) {
	n := len(T)
	m := len(P)

	if m == 0 {
		return
	}

	if n == 0 {
		return
	}

	if m > n {
		return
	}

	// Preprocess the pattern
	π := make([]int, m)
	initKMP(π, P)
	S = make([]int, 0, m)

	i, j := 0, 0
	for i < n {
		if P[j] == T[i] {
			i++
			j++
			if j == m {
				S = append(S, i-m)
				j = π[j-1]
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = π[j-1]
			}
		}
	}

	return S
}

func initKMP(π []int, P string) {
	π[0] = 0
	j, i := 1, 0
	for j < len(P) {
		if P[j] == P[i] {
			i++
			π[j] = i
			j++
		} else {
			if i == 0 {
				π[j] = 0
				j++
			} else {
				i = π[i-1]
			}
		}
	}
}
