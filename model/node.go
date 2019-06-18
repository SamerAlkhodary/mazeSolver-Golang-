package model

import "math"

const (
	Wall     = -1
	Left     = 0
	Right    = 2
	North    = 1
	South    = 3
	CellSize = 16
)

type Node struct {
	Neighbours []*Node
	Weights    []float64
	X, Y       int
	Visited    bool
}

func CreateNode(x, y int) *Node {
	return &Node{X: x / CellSize, Y: y / CellSize, Neighbours: make([]*Node, 4), Weights: make([]float64, 4), Visited: false}

}
func (n Node) GetCoord() (int, int) {
	offset := CellSize / 2
	return n.X*CellSize + offset, n.Y*CellSize + offset
}

func (n *Node) Connect(dir int, node *Node) *Node {
	n.Neighbours[dir] = node
	if node != nil {
		w := (n.X-node.X)*(n.X-node.X) + (n.Y-node.Y)*(n.Y-node.Y)
		fw := float64(w)
		fw = math.Pow(fw, 0.5)
		n.Weights[dir] = fw

	} else {
		n.Weights[dir] = Wall
	}

	return n

}

func (n Node) GetConnection(dir int) *Node {
	return n.Neighbours[dir]
}

func (n Node) GetWeight(dir int) float64 {
	return n.Weights[dir]

}
func (n Node) IsNode() bool {
	sum := float64(0)
	for i := 0; i < len(n.Weights); i++ {
		sum += n.Weights[i]

	}
	//if it is a dead end aka surrounded by 3 walls
	if sum == -3 {
		return true
	}
	return n.Weights[Left] != Wall && n.Weights[North] != Wall || n.Weights[Left] != Wall && n.Weights[South] != Wall || n.Weights[Right] != Wall && n.Weights[North] != Wall || n.Weights[Right] != Wall && n.Weights[South] != Wall
}
