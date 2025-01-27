package graphs

import "math"

type QueueBasedBelmanFordResult struct {
	s *Vertex
	D map[*Vertex]float64
	π map[*Vertex]*Vertex
}

/*
Negative cycle detection is not implemented in this version of the algorithm.
*/
func QueueBasedBelmanFordWithoutNegativeCycleFinding(G *WeightedAdjacencyListGraph, s *Vertex) QueueBasedBelmanFordResult {
	d := make(map[*Vertex]float64, len(G.V))
	π := make(map[*Vertex]*Vertex, len(G.V))
	inQ := make(map[*Vertex]bool, len(G.V))

	for _, v := range G.V {
		d[v] = math.MaxFloat64
		π[v] = nil
		inQ[v] = false
	}

	d[s] = 0
	Q := make(chan *Vertex, len(G.V))
	Q <- s
	inQ[s] = true

	for len(Q) > 0 {
		u := <-Q
		inQ[u] = false
		for _, v := range G.Adjacents(u) {
			if d[v] > d[u]+G.Weight(Edge{u, v}) {
				d[v] = d[u] + G.Weight(Edge{u, v})
				π[v] = u
				if !inQ[v] {
					Q <- v
					inQ[v] = true
				}
			}
		}
	}

	return QueueBasedBelmanFordResult{D: d, π: π, s: s}
}
