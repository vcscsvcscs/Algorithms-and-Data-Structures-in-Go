package graphs

import (
	"math"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	tests := []struct {
		name     string
		A        [][]float64
		expected [][]float64
		π        [][]int
	}{
		{
			name: "small graph",
			A: [][]float64{
				{0, 3, 8},
				{math.MaxFloat64, 0, 2},
				{math.MaxFloat64, 4, 0},
			},
			expected: [][]float64{
				{0, 3, 5},
				{math.MaxFloat64, 0, 2},
				{math.MaxFloat64, 4, 0},
			},
			π: [][]int{
				{0, 0, 1},
				{0, 1, 1},
				{0, 1, 2},
			},
		},
		{
			name: "simple graph",
			A: [][]float64{
				{0, 3, math.MaxFloat64, 7},
				{8, 0, 2, math.MaxFloat64},
				{5, math.MaxFloat64, 0, 1},
				{2, math.MaxFloat64, math.MaxFloat64, 0},
			},
			expected: [][]float64{
				{0, 3, 5, 6},
				{5, 0, 2, 3},
				{3, 6, 0, 1},
				{2, 5, 7, 0},
			},
			π: [][]int{
				{0, 0, 1, 0},
				{3, 1, 1, 2},
				{3, 0, 2, 2},
				{3, 0, 1, 3},
			},
		},
		{
			name: "graph with negative cycle",
			A: [][]float64{
				{0, 1, math.MaxFloat64},
				{math.MaxFloat64, 0, -1},
				{-1, math.MaxFloat64, 0},
			},
			expected: [][]float64{
				{0, 1, 0},
				{math.MaxFloat64, 0, -1},
				{-1, 0, -1},
			},
			π: [][]int{
				{0, 0, 1},
				{2, 1, 1},
				{2, 0, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			N := len(tt.A)
			D := make([][]float64, N)
			π := make([][]int, N)
			for i := range D {
				D[i] = make([]float64, N)
				π[i] = make([]int, N)
			}

			negativeCycleVertex := FloydWarshall(tt.A, D, π)

			for i := range D {
				for j := range D[i] {
					if D[i][j] != tt.expected[i][j] {
						t.Errorf("D[%d][%d] = %v; want %v", i, j, D[i][j], tt.expected[i][j])
					}
				}
			}

			if tt.name == "graph with negative cycle" && negativeCycleVertex == 0 {
				t.Errorf("Expected a negative cycle but got none")
			}
		})
	}
}
