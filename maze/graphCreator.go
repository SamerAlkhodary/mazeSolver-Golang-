package maze

import (
	"algorithms/greedy/mazeSolver/model"
	"image"
)

const (
	cellSize = 16
	offSet   = 2
)

func isWall(r, g, b, a uint32) bool {
	var black uint32
	black = 0
	return r == black && g == black && b == black
}
func CreatePath(dim model.Dimensions, image image.Image) (*model.Node, *model.Node) {

	numberOfVCell := (dim.Height-1)/cellSize + 2
	numberOfHCell := (dim.Width-1)/cellSize + 1
	var startNode *model.Node
	var goalNode *model.Node
	unconnected := make(map[int]*model.Node)
	graph := make([][]*model.Node, numberOfVCell)

	graph[0] = make([]*model.Node, numberOfHCell)
	for j := 0; j < dim.Width-offSet; j += cellSize {

		if !isWall(image.At(j+offSet, 0).RGBA()) {
			node := model.CreateNode(0, j)
			startNode = node.Connect(model.North, nil).Connect(model.Right, nil).Connect(model.Left, nil)
			graph[0][j/cellSize] = startNode
			unconnected[startNode.Y] = startNode
			break

		}

	}

	for i := 0; i < dim.Height-offSet; i += cellSize {
		graph[i/cellSize+1] = make([]*model.Node, numberOfHCell)

		var prev *model.Node
		for j := 0; j < dim.Width-offSet; j += cellSize {

			n := createNodes(image, i, j)
			if prev != nil && n != nil && prev.GetWeight(model.Right) != -1 && n.GetWeight(model.Left) != -1 {
				prev.Connect(model.Right, n)
				n.Connect(model.Left, prev)

			}

			graph[i/cellSize+1][j/cellSize] = n
			if n != nil {
				prev = n
				connectVertical(n, unconnected)

			}

		}

	}
	graph[len(graph)-1] = make([]*model.Node, numberOfHCell)

	for j := 0; j < dim.Width-offSet; j += cellSize {

		if !isWall(image.At(j+offSet, dim.Height-1).RGBA()) {
			node := model.CreateNode(dim.Height+cellSize, j)
			goalNode = node.Connect(model.South, nil).Connect(model.Right, nil).Connect(model.Left, nil)
			graph[len(graph)-1][j/cellSize] = goalNode
			break

		}
	}
	connectVertical(goalNode, unconnected)

	//prints out the node table for debugging reasons

	/*for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == nil {
				fmt.Print("-", " ")

			} else {
				fmt.Print("v", " ")
			}

		}
		fmt.Println()
	}*/

	return startNode, goalNode
}
func createNodes(image image.Image, i, j int) *model.Node {

	tmpNode := model.CreateNode(i+16, j)

	if isWall(image.At(j, i+offSet).RGBA()) {

		tmpNode.Connect(model.Left, nil)

	}
	if isWall(image.At(j+cellSize, i+offSet).RGBA()) {
		tmpNode.Connect(model.Right, nil)

	}
	if isWall(image.At(j+offSet, i).RGBA()) {
		tmpNode.Connect(model.North, nil)
	}
	if isWall(image.At(j+offSet, i+cellSize).RGBA()) {
		tmpNode.Connect(model.South, nil)

	}
	if tmpNode.IsNode() {
		return tmpNode
	}
	return nil

}
func connectVertical(n *model.Node, unconnected map[int]*model.Node) {
	if node, ok := unconnected[n.Y]; ok {
		if node.GetWeight(model.South) != model.Wall && n.GetWeight(model.North) != model.Wall {
			n.Connect(model.North, node)
			node.Connect(model.South, n)

		}
	}
	unconnected[n.Y] = n
}
