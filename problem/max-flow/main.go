package main

import "fmt"

func main() {

	nodeNumber := 6

	graphResidual := make([][]int, nodeNumber)

	for i := 0; i < nodeNumber; i++ {
		graphResidual[i] = make([]int, nodeNumber)
	}

	graphResidual[0][1] = 4
	graphResidual[0][2] = 2
	graphResidual[1][2] = 1
	graphResidual[1][3] = 2
	graphResidual[1][4] = 4
	graphResidual[2][4] = 2
	graphResidual[3][5] = 3
	graphResidual[4][5] = 3

	maxFlow := 0

	for {
		minCapacity := findRoad(nodeNumber, graphResidual)
		if minCapacity == 0 {
			break
		}
		maxFlow += minCapacity
	}
	fmt.Println(maxFlow)
}
