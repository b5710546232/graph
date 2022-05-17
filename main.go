package main

import "github.com/b5710546232/go-graph/graph"

func main() {
	g := graph.NewGraph()

	vertices := []string{"A", "B", "C", "D", "E"}
	for _, v := range vertices {
		g.AddVertex(v)
	}

	g.AddEdge("A", "B")
	g.AddEdge("A", "D")
	g.AddEdge("A", "E")
	g.AddEdge("B", "C")
	g.AddEdge("B", "E")
	g.AddEdge("D", "E")
	g.AddEdge("E", "C")

	g.PrintGraph()
}
