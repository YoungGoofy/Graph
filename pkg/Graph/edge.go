package graph

import (
	"fmt"
	"log"
)

type edge struct {
	start string
	end   string
}

func newEdge(start, end string) edge {
	return edge{start: start, end: end}
}

func (g *Graph) AddEdge(start, end string) error {
	if g.findNodes(start, end) {
		e := newEdge(start, end)
		if g.checkEdgeDuplicates(e) {
			g.edges = append(g.edges, e)
		} else {
			return fmt.Errorf("error: узел %v уже существует", e)
		}
	} else {
		return fmt.Errorf("error: какой-то из нод не существует")
	}
	return nil
}

func (g *Graph) AddEdges(edges ...Edge) []error {
	err := []error{}
	for index, edge := range edges {
		if g.findNodes(edge.Start, edge.End) {
			e := newEdge(edge.Start, edge.End)
			if g.checkEdgeDuplicates(e) {
				g.edges = append(g.edges, e)
			} else {
				err = append(err, fmt.Errorf("error: узел %v уже существует", e))
			}
		} else {
			err = append(err, fmt.Errorf("error: в узле под номером %d ошибка", index+1))
		}
	}
	if len(err) != 0 {
		return err
	}
	return nil
}

func (g *Graph) DeleteEdge(edge Edge) error {
	for i := 0; i < len(g.edges); i++ {
		if g.edges[i].start == edge.Start && g.edges[i].end == edge.End {
			g.edges = append(g.edges[i:], g.edges[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error: такого узла не существует")
}

func (g Graph) PrintEdges() {
	for index, edge := range g.edges {
		fmt.Printf("Узел %d = [%s %s]\n", index, edge.start, edge.end)
	}
}

func (g Graph) checkEdgeDuplicates(edge edge) bool {
	for _, item := range g.edges {
		if item == edge {
			return false
		}
	}
	return true
}

func (g Graph) findNodes(node1, node2 string) bool {
	var count int
	if node1 == node2 {
		log.Println("Нельзя объединить одну и ту же ноду")
		return false
	}
	for _, item := range g.nodes {
		if item.Name == node1 || item.Name == node2 {
			count++
		}
	}
	return count == 2
}

type Edge struct {
	Start string
	End   string
}
