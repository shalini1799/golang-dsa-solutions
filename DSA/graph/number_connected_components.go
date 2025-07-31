func countComponents(n int, edges [][]int) int {
	//create a graph
	graph := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	visited := make(map[int]bool)
	var dfs func(int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}
	
	components := 0
	for i := 0; i < n; i++ {		
		if !visited[i] {
			dfs(i)
			components++
		}
	}
	return components
}