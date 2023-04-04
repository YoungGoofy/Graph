package graph

import (
	"fmt"
)

type node struct {
	ID    int
	Name  string
	X, Y  int
	Edges []edge
}

type Node struct {
	ID    int
	Name  string
	X, Y  int
	Edges []edge
}

func newNode(id int, name string) node {
	return node{Name: name, ID: id}
}

func (g *Graph) AddNode(id int, name string) error {
	if g.checkNodeDuplicates(name) {
		g.nodes = append(g.nodes, newNode(id, name))
	} else {
		return fmt.Errorf("error: нода %s уже существует", name)
	}
	return nil
}

func (g *Graph) AddNodes(nodes ...Node) []error {
	err := []error{}
	for _, node := range nodes {
		if g.checkNodeDuplicates(node.Name) {
			g.nodes = append(g.nodes, newNode(node.ID, node.Name))
		} else {
			err = append(err, fmt.Errorf("error: нода %s уже существует", node.Name))
		}
	}
	if len(err) != 0 {
		return err
	}
	return nil
}

func (g *Graph) DeleteNode(id int) error {
	for i := 0; i < len(g.nodes); i++ {
		if g.nodes[i].ID == id {
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
