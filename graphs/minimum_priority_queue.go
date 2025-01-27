package graphs

import "sort"

type NodeMinPrioQueue []*Vertex

func NewNodeMinPrioQueue(G *WeightedAdjacencyListGraph, r *Vertex) NodeMinPrioQueue {
	Q := make(NodeMinPrioQueue, 0, len(G.V)-1)
	for _, v := range G.V {
		if v != r {
			Q = append(Q, v)
		}
	}

	return Q
}

func (Q *NodeMinPrioQueue) Adjust(c *map[*Vertex]float64) {
	sort.Slice(*Q, func(i, j int) bool {
		return (*c)[(*Q)[i]] < (*c)[(*Q)[j]]
	})

}

func (Q *NodeMinPrioQueue) RemMin() *Vertex {
	v := (*Q)[0]
	*Q = (*Q)[1:]
	return v
}

func (Q *NodeMinPrioQueue) Contains(v *Vertex) bool {
	for _, u := range *Q {
		if u == v {
			return true
		}
	}

	return false
}
