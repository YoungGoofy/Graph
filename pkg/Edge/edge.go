package edge

import (
	"fmt"

	node "github.com/YoungGoofy/Graph/pkg/Node"
)

type Edge [2]node.Node
type Edges []Edge

var (
	edge  Edge
	edges []Edge
)

func AddEdge(node1, node2 node.Node) error {
	nodes := node.ReturnNodeList()
	if nodes.FindNodes(node1, node2) {
		edge[0] = node1
		edge[1] = node2
		edges = append(edges, edge)
		return nil
	}
	return fmt.Errorf(`
		Вершины не сущетвует...
		Для проверки вершин добавьте команду fmt.Println(node.ReturnNodeList())
		Чтоб добавить вершину добавьте команду node.AddNode() или node.AddList()
		`)
}

func DeleteEdge(e Edge) error {
	for index, item := range edges {
		if item == e {
			edges[index] = edges[len(edges)-1]
			edges[len(edges)-1] = Edge{}
			edges = edges[:len(edges)-1]
			return nil
		}
	}
	return fmt.Errorf("Связи не существует...")
}

func ReturnEdgeList() Edges {
	return edges
}
