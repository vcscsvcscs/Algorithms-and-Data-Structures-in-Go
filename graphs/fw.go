package graphs

import "math"

// FloydWarshall implements the Floyd-Warshall algorithm for all-pairs shortest paths.
// returns the index of a vertex that is part of a negative cycle, if any.
// M T (n) ∈ Θ(n3) and mT (n) ∈ Θ(n2)
func FloydWarshall(A, D [][]float64, π [][]int) int {
	N := len(A)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			D[i][j] = A[i][j]
			if i != j && A[i][j] < math.MaxFloat64 {
				π[i][j] = i
			} else {
				π[i][j] = 0
			}
		}
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if D[i][k]+D[k][j] < D[i][j] {
					D[i][j] = D[i][k] + D[k][j]
					π[i][j] = π[k][j]
					if i == j && D[i][j] < 0 {
						return i
					}
				}
			}
		}
	}

	return 0
}
