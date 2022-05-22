package main

func findRoad(nodeNumber int, graphResidual [][]int) int {

	minCapacity := 1<<63 - 1
	visited := make([]bool, nodeNumber)
	parent := make([]int, nodeNumber)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}

	var dfs func(x int)
	dfs = func(x int) {
		for i := 0; i < nodeNumber; i++ {
			if !visited[i] && graphResidual[x][i] > 0 {

				visited[i] = true
				parent[i] = x
				dfs(i)
			}
		}
	}
	visited[0] = true
	dfs(0)

	for x := 5; parent[x] != -1; x = parent[x] {
		minCapacity = min(minCapacity, graphResidual[parent[x]][x])
	}

	if minCapacity == 1<<63-1 {
		return 0
	}

	for x := 5; parent[x] != -1; x = parent[x] {
		graphResidual[parent[x]][x] -= minCapacity
		graphResidual[x][parent[x]] += minCapacity
	}

	return minCapacity
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
