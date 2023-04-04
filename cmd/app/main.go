package main

import (
	"log"

	g "github.com/YoungGoofy/Graph/pkg/Graph"
)

func main() {

	var graph g.Graph
	err := graph.AddNodes("1", "2", "3", "4", "5", "6")
	if err != nil {
		for _, e := range err {
			log.Println(e)
		}
	}

	err = graph.AddEdges(
		g.Edge{Start: "1", End: "2"},
		g.Edge{Start: "1", End: "3"},
		g.Edge{Start: "1", End: "4"},
		g.Edge{Start: "1", End: "5"},
		g.Edge{Start: "1", End: "6"},
	)
	if err != nil {
		for _, e := range err {
			log.Println(e)
		}
	}
	graph.Print()
}
