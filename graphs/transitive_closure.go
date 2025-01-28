package graphs

// TransitiveClosure computes the transitive closure of a directed graph
// represented by its adjacency matrix A. The transitive closure is a matrix T
// such that T[i][j] is true if there is a path from i to j in the graph.
// The algorithm runs in T (n) ∈ Θ(n3) time.
func TransitiveClosure(A, T [][]bool) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			T[i][j] = A[i][j]
		}
		T[i][i] = true
	}

	for k := 0; k < len(A); k++ {
		for i := 0; i < len(A); i++ {
			for j := 0; j < len(A); j++ {
				T[i][j] = T[i][j] || (T[i][k] && T[k][j])
			}
		}
	}
}
