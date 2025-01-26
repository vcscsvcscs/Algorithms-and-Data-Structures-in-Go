package graphs

func BreathFirstSearch(G *AdjacencyListGraph, Start *Vertex) {
	for _, v := range G.V {
		v.π = nil
		v.d = -1
		v.color = "white"
	}
	Start.d = 0
	Start.color = "gray"
	Q := make(chan *Vertex, len(G.V))
	Q <- Start
	for len(Q) > 0 {
		u := <-Q
		for _, v := range G.Adjacents(u) {
			if v.d == -1 {
				v.d = u.d + 1
				v.π = u
				v.color = "gray"
				Q <- v
			} else {
				continue
			}
		}
		u.color = "black"
	}
}
