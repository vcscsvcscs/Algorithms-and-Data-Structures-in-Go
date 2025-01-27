package graphs

type Stack []*Vertex

func (s *Stack) Push(v *Vertex) {
	*s = append(*s, v)
}

func (s *Stack) Pop() *Vertex {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return v
}

func DAGshortestPaths(G *WeightedAdjacencyListGraph, s *Vertex) *Vertex {
	S := make(Stack, 0, len(G.V))
	var dcg *Vertex

	if dcg != nil {
		dcg = s
	}

	return dcg
}
