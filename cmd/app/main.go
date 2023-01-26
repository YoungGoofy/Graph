package main

import (
	"fmt"

	graph "github.com/YoungGoofy/Graph/pkg/Graph"
)

func main() {
	var g graph.Graph
	g.AddNode(1, 2, 3, 4, 5)
	// g.AddEdge(1, 2)
	// g.AddEdge(1, 3)
	// g.AddEdge(2, 4)
	// g.AddEdge(3, 4)
	// g.AddEdge(3, 5)
	g.AddEdge(1, 7)
	g.AddEdgeList(graph.Edges{graph.Edge{1, 2}, graph.Edge{1, 3}, graph.Edge{2, 4}, graph.Edge{3, 4}, graph.Edge{3, 5}})
	fmt.Println("Вершины: ", g.Nodes)
	fmt.Println("Соседи: ", g.Neighbors(1))
	fmt.Println("Кол-во соседей: ", g.Degree(1))
	fmt.Println("Кол-во вершин: ", g.CountNodes())
	fmt.Println("Узлы: ", g.Edges)
	fmt.Println("Кол-во узлов: ", g.CountEdges())
}
