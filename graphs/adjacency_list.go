package graphs

type Vertex struct {
	Name  int
	Value any
	d     int
	π     *Vertex
	f     int
	color string
}

type Edge struct {
	V *Vertex
	U *Vertex
}

type WeightedEdge struct {
	U *Vertex
	V *Vertex
	W float64
}

type AdjacencyListEdge struct {
	V    *Vertex
	U    *AdjacencyListEdge // Next
	Type string             // classifies the edge
}

type AdjacencyListWeightedEdge struct {
	V    *Vertex
	W    float64
	U    *AdjacencyListWeightedEdge // Next
	Type string                     // classifies the edge
}

type AdjacencyListGraph struct {
	V []*Vertex
	a map[*Vertex]*AdjacencyListEdge // Adjacency List
}

type WeightedAdjacencyListGraph struct {
	V []*Vertex                              // Vertices
	a map[*Vertex]*AdjacencyListWeightedEdge // Adjacency List
}

func (G *AdjacencyListGraph) Adjacents(v *Vertex) []*Vertex { // Adjacents
	adjacents := make([]*Vertex, 0)
	u := G.a[v]
	for u != nil {
		adjacents = append(adjacents, u.V)
		u = u.U
	}

	return adjacents
}

func (G *WeightedAdjacencyListGraph) Adjacents(v *Vertex) []*Vertex { // Adjacents
	adjacents := make([]*Vertex, 0)
	u := G.a[v]
	for u != nil {
		adjacents = append(adjacents, u.V)
		u = u.U
	}

	return adjacents
}

func (G *WeightedAdjacencyListGraph) Weight(E Edge) float64 {
	u := G.a[E.U]
	for u != nil {
		if u.V == E.V {
			return u.W
		}
		u = u.U
	}

	return 0
}
