package graph

import (
	"github.com/YoungGoofy/Graph/pkg/Graph/edge"
	"github.com/YoungGoofy/Graph/pkg/Graph/node"
)

type Graph struct {
	Nodes []node.Node
	Edges []edge.Edge
}

func (g *Graph) AddNode(name string) {
	g.Nodes = append(g.Nodes, node.NewNode(name))
}

func (g *Graph) AddNodes(names ...string) {
	for _, name := range names {
		g.Nodes = append(g.Nodes, node.NewNode(name))
	}
}

func (g *Graph) DeleteNode() {

}
