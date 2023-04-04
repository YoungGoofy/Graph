package graph

import "fmt"

type Graph struct {
	nodes []node
	edges []edge
}

func (g *Graph) Print() {
	fmt.Println("NODES")
	g.PrintNodes()
	fmt.Println("\nEDGES")
	g.PrintEdges()
}
