package graphs

import "sort"

type KruskalStates struct {
	π map[*Vertex]*Vertex
	s map[*Vertex]int
}

// Θ(1)
func (ks *KruskalStates) MakeSet(v *Vertex) {
	ks.π[v] = v
	ks.s[v] = 1
}

// O(log n)
func (ks *KruskalStates) FindSet(v *Vertex) *Vertex {
	if ks.π[v] != v {
		ks.π[v] = ks.FindSet(ks.π[v])
	}

	return ks.π[v]
}

// Θ(1)
func (ks *KruskalStates) Union(x *Vertex, y *Vertex) {
	if ks.s[x] < ks.s[y] {
		ks.π[x] = y
		ks.s[y] += ks.s[x]
	} else {
		ks.π[y] = x
		ks.s[x] += ks.s[y]
	}
}

func Kruskal(G *WeightedAdjacencyListGraph, A []Edge) {
	ks := KruskalStates{π: make(map[*Vertex]*Vertex, len(G.V)), s: make(map[*Vertex]int, len(G.V))}
	for _, v := range G.V {
		ks.MakeSet(v)
	}
	A = make([]Edge, 0)
	k := len(G.V)
	Q := minPriorityQueue(&A, func(e Edge) float64 { return G.Weight(e) })
	for k > 1 && len(Q) > 0 {
		e := <-Q
		x := ks.FindSet(e.U)
		y := ks.FindSet(e.V)
		if x != y {
			
	}

}

func minPriorityQueue(E *[]Edge, Weight func(Edge) float64) chan Edge {
	Q := make(chan Edge, len(*E))
	//  min sort on queue
	sort.Slice(*E, func(i, j int) bool { return Weight((*E)[i]) < Weight((*E)[j]) })
	for _, e := range *E {
		Q <- e
	}
	return Q
}
