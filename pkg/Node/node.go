package node

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

func DeleteNode(n Node) {
	for index, item := range nodes {
		if item == n {
			nodes[index] = nodes[len(nodes)-1]
			nodes[len(nodes)-1] = ""
			nodes = nodes[:len(nodes)-1]
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

func ReturnNodeList() Nodes {
	return nodes
}

func (n Nodes) checkDuplicates(num Node) bool {
	for _, item := range n {
		if item == num {
			return false
		}
	}
	return true
}

func (n Nodes) FindNodes(node1, node2 Node) bool {
	var count int
	for _, item := range n {
		if item == node1 || item == node2 {
			count++
		}
	}
	if count == 2 {
		return true
	}
	return false
}
