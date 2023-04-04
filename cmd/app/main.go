package main

import (
	"image/png"
	"log"
	"os"

	g "github.com/YoungGoofy/Graph/pkg/Graph"
)

func main() {

	var graph g.Graph
	if err := graph.AddNodes(
		g.Node{ID: 1, Name: "1"},
		g.Node{ID: 2, Name: "2"},
		g.Node{ID: 3, Name: "3"},
		g.Node{ID: 4, Name: "4"},
		g.Node{ID: 5, Name: "5"},
	); err != nil {
		for _, e := range err {
			log.Println(e)
		}
	}

	if err := graph.AddEdges(
		g.Edge{StartID: 1, EndID: 2},
		g.Edge{StartID: 1, EndID: 3},
		g.Edge{StartID: 1, EndID: 4},
		g.Edge{StartID: 1, EndID: 5},
	); err != nil {
		for _, e := range err {
			log.Println(e)
		}
	}
	graph.Print()
	img := graph.DrawGraph()
	f, err := os.Create("graph.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}
