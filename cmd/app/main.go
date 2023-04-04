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
		g.Node{ID: 6, Name: "6"},
		g.Node{ID: 7, Name: "7"},
		g.Node{ID: 8, Name: "8"},
		g.Node{ID: 9, Name: "9"},
		g.Node{ID: 10, Name: "10"},
		g.Node{ID: 11, Name: "11"},
		g.Node{ID: 12, Name: "12"},
		g.Node{ID: 13, Name: "13"},
		g.Node{ID: 14, Name: "14"},
		g.Node{ID: 15, Name: "15"},
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
		g.Edge{StartID: 1, EndID: 6},
		g.Edge{StartID: 2, EndID: 7},
		g.Edge{StartID: 3, EndID: 8},
		g.Edge{StartID: 3, EndID: 9},
		g.Edge{StartID: 4, EndID: 10},
		g.Edge{StartID: 4, EndID: 11},
		g.Edge{StartID: 5, EndID: 12},
		g.Edge{StartID: 5, EndID: 13},
		g.Edge{StartID: 6, EndID: 14},
		g.Edge{StartID: 6, EndID: 15},
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
