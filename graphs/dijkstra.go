package graphs

import "math"

type DijkstraState struct {
	d map[*Vertex]float64
	π map[*Vertex]*Vertex
}

func Dijkstra(G *WeightedAdjacencyListGraph, s *Vertex) DijkstraState {
	d := make(map[*Vertex]float64, len(G.V))
	π := make(map[*Vertex]*Vertex, len(G.V))

	for _, v := range G.V {
		d[v] = math.MaxFloat64
		π[v] = nil
	}

	d[s] = 0
	Q := NewNodeMinPrioQueue(G, s)

	u := s
	for d[u] < math.MaxFloat64 && len(Q) > 0 {
		for _, v := range G.Adjacents(u) {
			if d[v] > d[u]+G.Weight(Edge{u, v}) {
				d[v] = d[u] + G.Weight(Edge{u, v})
				π[v] = u
				Q.Adjust(&d)
			} else {
				break
			}
		}
		u = Q.RemMin()
	}

	return DijkstraState{d: d, π: π}
}
