func validGraph(n int, edges [][]int) bool {
	// a valid tree of n node must have n-1 edges
	if len(edges) != n-1 {
		return false
	}

	//build a graph
	graph := make(map[int][]int)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }

	visited := make(map[int]bool)
	var dfs func(int, int)bool
	dfs = func(node int, parent int) bool{
		if visited[node] {
			return false // cycle detected
		}	
		visited[node] = true
		for _, neighbor := range graph[node] {
			if neighbor == parent{
				continue // skip the parent node
			}
			if !dfs(neighbor, node) {
                return false
            }
			
		}
		return true // no cycle detected
	}

	//loop over the edges
	if !dfs(0, -1) {
        return false // cycle detected
    }

	if len(visited) != n {
		return false // not all nodes are visited
	}
	return true // all nodes are visited and edges count is n-1

}