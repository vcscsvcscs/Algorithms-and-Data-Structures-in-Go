package graphs

import "sort"

type PrimStates struct {
	π map[*Vertex]*Vertex
	c map[*Vertex]float64
}

type PrimMinPrioQueue []*Vertex

func Prim(G *WeightedAdjacencyListGraph, r *Vertex) PrimStates {
	π := make(map[*Vertex]*Vertex, len(G.V))
	c := make(map[*Vertex]float64, len(G.V))

	for _, v := range G.V {
		v.d = -1
		π[v] = nil
	}

	c[r] = 0
	Q := minPrQ(G, r)

	u := r
	for len(Q) > 0 {
		for _, v := range G.Adjacents(u) {
			π[v] = u
			c[v] = G.Weight(Edge{u, v})
			Q.Adjust(&c)
		}
		u = Q.RemMin()
	}

	return PrimStates{π: π, c: c}

}

func minPrQ(G *WeightedAdjacencyListGraph, r *Vertex) PrimMinPrioQueue {
	Q := make(PrimMinPrioQueue, 0, len(G.V)-1)
	for _, v := range G.V {
		if v != r {
			Q = append(Q, v)
		}
	}

	return Q
}

func (Q *PrimMinPrioQueue) Adjust(c *map[*Vertex]float64) {
	sort.Slice(*Q, func(i, j int) bool {
		return (*c)[(*Q)[i]] < (*c)[(*Q)[j]]
	})

}

func (Q *PrimMinPrioQueue) RemMin() *Vertex {
	v := (*Q)[0]
	*Q = (*Q)[1:]
	return v
}
