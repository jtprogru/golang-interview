/*
There are `n` servers numbered from `0` to `n - 1` connected
by undirected server-to-server connections forming a network
where `connections[i] = [ai, bi]` represents a connection
between servers `ai` and `bi`. Any server can reach other
servers directly or indirectly through the network.

A critical connection is a connection that, if removed,
will make some servers unable to reach some other server.

Return all critical connections in the network in any order.
*/

package task0012

func Solution(n int, connections [][]int) [][]int {
	var buildGraph func() [][]int
	buildGraph = func() [][]int {
		graph := make([][]int, n, n)
		for _, v := range connections {
			src, dest := v[0], v[1]
			graph[src] = append(graph[src], dest)
			graph[dest] = append(graph[dest], src)
		}

		return graph
	}

	graph := buildGraph()
	res, timer := [][]int{}, 0
	visited, timeStamps := make([]bool, n), make([]int, n)
	var criticalConnectionsHelperDFS func(curVertex, prevVertex int)
	criticalConnectionsHelperDFS = func(curVertex, prevVertex int) {
		visited[curVertex] = true
		timeStamps[curVertex] = timer
		timer++
		curTimeStamp := timeStamps[curVertex]
		for _, v := range graph[curVertex] {
			if v == prevVertex {
				continue
			}
			if !visited[v] {
				criticalConnectionsHelperDFS(v, curVertex)
			}

			timeStamps[curVertex] = findMin(timeStamps[curVertex], timeStamps[v])
			if curTimeStamp < timeStamps[v] {
				res = append(res, []int{curVertex, v})
			}
		}
	}

	criticalConnectionsHelperDFS(0, -1)
	return res
}

func findMin(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
