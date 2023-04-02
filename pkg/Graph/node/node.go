package node

type Node struct {
	Name string
	// X, Y int
}

func NewNode(name string) Node {
	return Node{Name: name}
}
