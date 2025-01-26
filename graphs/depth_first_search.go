package graphs

func DepthFirstSearch(G *AdjacencyListGraph) {
	for _, u := range G.V {
		u.color = "white"
	}
	time := 0
	for _, r := range G.V {
		if r.color == "white" {
			r.π = nil
			DepthFirstSearchVisit(G, r, &time)
		}
	}
}

func DepthFirstSearchVisit(G *AdjacencyListGraph, u *Vertex, time *int) {
	*time++
	u.d = *time
	u.color = "gray"
	for _, v := range G.Adjacents(u) {
		if v.color == "white" {
			v.π = u
			DepthFirstSearchVisit(G, v, time)
		} else if v.color == "gray" {
			backEdge(u, v)
		}
	}
	*time++
	u.f = *time

}

func backEdge(u *Vertex, v *Vertex) {
	// Do something with the back edges
}
