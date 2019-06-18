package main

import (
	"algorithms/greedy/mazeSolver/algorithm"
	"algorithms/greedy/mazeSolver/maze"
	"fmt"
	"os"
)

func main() {
	inputFile, outpuFile := readArgs()
	mazeImage, dim := maze.ParseMaze(inputFile)
	startNode, targetNode := maze.CreatePath(dim, mazeImage)
	item := algorithm.FindPath(startNode, targetNode)
	maze.CreateImageResult(mazeImage, outpuFile, item)

}

func readArgs() (string, string) {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enaugh arguments ! make sure to type in the file name and the result file name")
		os.Exit(0)

	}
	return args[1], args[2]
}
