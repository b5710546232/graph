package graph

import "fmt"

type Graph struct {
	adj map[string][]string
}

func NewGraph() *Graph {
	g := &Graph{}
	g.adj = make(map[string][]string)
	return g
}

func (g *Graph) AddVertex(v string) {
	g.adj[v] = make([]string, 0)
}

func (g *Graph) AddEdge(u, v string) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) PrintGraph() {
	for v, list := range g.adj {
		str := ""
		for _, vv := range list {
			str += fmt.Sprintf("-%s", vv)
		}
		fmt.Printf("%s-->%s", v, str)
		fmt.Printf("\n")
	}
}
