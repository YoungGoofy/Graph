package graph

import "fmt"

type (
	Node  interface{}
	Nodes []Node
	Edge  [2]Node
	Edges []Edge
	Graph struct {
		Nodes []Node
		Edges []Edge
	}
)

var edge Edge

func (g *Graph) AddNode(n ...Node) {
	for _, item := range n {
		if g.checkDuplicates(item) {
			g.Nodes = append(g.Nodes, item)
		}
	}
}

func (g *Graph) AddNodeList(n ...Nodes) {
	for _, array := range n {
		for _, item := range array {
			if g.checkDuplicates(item) {
				g.Nodes = append(g.Nodes, item)
			}
		}
	}
}

func (g *Graph) AddEdge(node1, node2 Node) {
	if g.findNodes(node1, node2) {
		edge[0] = node1
		edge[1] = node2
		g.Edges = append(g.Edges, edge)
	}
}

func (g *Graph) AddEdgeList(e ...Edges) {
	for _, array := range e {
		for _, item := range array {
			if g.findNodes(item[0], item[1]) {
				edge[0] = item[0]
				edge[1] = item[1]
				g.Edges = append(g.Edges, edge)
			}

		}
	}
}

func (g *Graph) DeleteNode(n Node) {
	for index, item := range g.Nodes {
		if item == n {
			g.Nodes[index] = g.Nodes[len(g.Nodes)-1]
			g.Nodes[len(g.Nodes)-1] = ""
			g.Nodes = g.Nodes[:len(g.Nodes)-1]
		}
	}
}

func (g *Graph) DeleteEdge(e Edge) error {
	for index, item := range g.Edges {
		if item == e {
			g.Edges[index] = g.Edges[len(g.Edges)-1]
			g.Edges[len(g.Edges)-1] = Edge{}
			g.Edges = g.Edges[:len(g.Edges)-1]
			return nil
		}
	}
	return fmt.Errorf("Узла не существует")
}

func (g Graph) CountNodes() int {
	return len(g.Nodes)
}

func (g Graph) CountEdges() int {
	return len(g.Edges)
}

func (g *Graph) ClearAll() {
	g.Nodes = []Node{}
	g.Edges = []Edge{}
}

func (g Graph) findNodes(node1, node2 Node) bool {
	var count int
	for _, item := range g.Nodes {
		if item == node1 || item == node2 {
			count++
		}
	}
	if count == 2 {
		return true
	}
	return false
}

func (g Graph) checkDuplicates(num Node) bool {
	for _, item := range g.Nodes {
		if item == num {
			return false
		}
	}
	return true
}
