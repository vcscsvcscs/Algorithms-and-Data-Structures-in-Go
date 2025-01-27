package graphs

type PrimStates struct {
	π map[*Vertex]*Vertex
	c map[*Vertex]float64
}

func Prim(G *WeightedAdjacencyListGraph, r *Vertex) PrimStates {
	π := make(map[*Vertex]*Vertex, len(G.V))
	c := make(map[*Vertex]float64, len(G.V))

	for _, v := range G.V {
		v.d = -1
		π[v] = nil
	}

	c[r] = 0
	Q := NewNodeMinPrioQueue(G, r)

	u := r
	for len(Q) > 0 {
		for _, v := range G.Adjacents(u) {
			if Q.Contains(v) && G.Weight(Edge{u, v}) < c[v] {
				π[v] = u
				c[v] = G.Weight(Edge{u, v})
				Q.Adjust(&c)
			} else {
				break
			}
		}
		u = Q.RemMin()
	}

	return PrimStates{π: π, c: c}

}
