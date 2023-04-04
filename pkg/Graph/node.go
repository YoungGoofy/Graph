package graph

import (
	"fmt"
)

type node struct {
	Name string
	X, Y int
}

func newNode(name string) node {
	return node{Name: name}
}

func (g *Graph) AddNode(name string) error {
	if g.checkNodeDuplicates(name) {
		g.nodes = append(g.nodes, newNode(name))
	} else {
		return fmt.Errorf("error: нода %s уже существует", name)
	}
	return nil
}

func (g *Graph) AddNodes(names ...string) []error {
	err := []error{}
	for _, name := range names {
		if g.checkNodeDuplicates(name) {
			g.nodes = append(g.nodes, newNode(name))
		} else {
			err = append(err, fmt.Errorf("error: нода %s уже существует", name))
		}
	}
	if len(err) != 0 {
		return err
	}
	return nil
}

func (g *Graph) DeleteNode(name string) error {
	for i := 0; i < len(g.nodes); i++ {
		if g.nodes[i].Name == name {
			g.nodes = append(g.nodes[:i], g.nodes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("error: такой ноды нет")
}

func (g *Graph) UpdateNode(oldNode, newNode string) error {
	for i := 0; i < len(g.nodes); i++ {
		if g.nodes[i].Name == oldNode {
			if g.checkNodeDuplicates(newNode) {
				g.nodes[i].Name = newNode
			} else {
				return fmt.Errorf("error: нода %s уже существует", newNode)
			}
			return nil
		}
	}
	return fmt.Errorf("error: выбранной ноды не существует")
}

func (g Graph) PrintNodes() {
	for index, node := range g.nodes {
		fmt.Printf("Node %d = %s\n", index+1, node.Name)
	}
}

func (g Graph) checkNodeDuplicates(node string) bool {
	for _, item := range g.nodes {
		if item.Name == node {
			return false
		}
	}
	return true
}
