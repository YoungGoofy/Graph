package node

import "fmt"

type Node interface{}
type Nodes []Node

var nodes Nodes

func AddNode(n ...Node) {
	for _, item := range n {
		if nodes.checkDuplicates(item) {
			nodes = append(nodes, item)
		}
	}
}

func AddList(n ...Nodes) {
	for _, array := range n {
		for _, item := range array {
			if nodes.checkDuplicates(item) {
				nodes = append(nodes, item)
			}
		}
	}
}

func PrintNode() {
	fmt.Printf("Nodes: %v", nodes)
}

func (n Nodes) checkDuplicates(num Node) bool {
	for _, item := range n {
		if item == num {
			return false
		}
	}
	return true
}
