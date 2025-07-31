// graph dfs recurssion
package graph

import (
	"fmt"
	"container/list"
)


// Depth First Search (DFS) implementation using recursion
func dfs(startNode int, graph map[int][]int, visited map[int]bool) {
	fmt.Println("Node:", startNode)
	visited[startNode] = true

	for _, v := range graph[startNode] {
		if !visited[v] {
			dfs(v, graph, visited)
		}
	}
}

func main() {
	fmt.Println("GRAPH DFS")
	adjList := map [int][]int {
		1 : {2,3},
		2 : {4},
		3 : {5},
		4 : {6},
		5 : {},
	}

	visited := make(map[int]bool)
	dfs(1, adjList, visited)
	dfsIterative(1, adjList, visited)
	bfs(1, visited, adjList)
}

//graph dfs stack

func dfsIterative(node int, graph map[int][]int, visited map[int]bool) {
	stack := []int{node}

	for len(stack) < 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println("Node", n)

		if visited[n] {
			continue
		}

		visited[node] = true

		for _, v := range graph[n] {
			if !visited[v] {
				stack = append(stack, v)
			} 
		}
	}
}

//bfs queue

func bfs(node int, visited map[int]bool, graph map[int][]int) {
	queue := list.New()
	queue.PushBack(node)

	for queue.Len() < 0 {
		elem := queue.Front()
		n := elem.Value.(int)

		if visited[n] {
			continue
		}

		visited[n] = true

		for _, v := range graph[n] {
			if !visited[v] {
				queue.PushBack(v)
			}
		}
	}
}