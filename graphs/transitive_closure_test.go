package graphs

import (
	"reflect"
	"testing"
)

func TestTransitiveClosure(t *testing.T) {
	tests := []struct {
		name string
		A    [][]bool
		want [][]bool
	}{
		{
			name: "simple graph",
			A: [][]bool{
				{false, true, false},
				{false, false, true},
				{false, false, false},
			},
			want: [][]bool{
				{true, true, true},
				{false, true, true},
				{false, false, true},
			},
		},
		{
			name: "disconnected graph",
			A: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
			want: [][]bool{
				{true, false, false},
				{false, true, false},
				{false, false, true},
			},
		},
		{
			name: "fully connected graph",
			A: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, true, true},
			},
			want: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, true, true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T := make([][]bool, len(tt.A))
			for i := range T {
				T[i] = make([]bool, len(tt.A))
			}
			TransitiveClosure(tt.A, T)
			if !reflect.DeepEqual(T, tt.want) {
				t.Errorf("TransitiveClosure() = %v, want %v", T, tt.want)
			}
		})
	}
}
