package graph

import (
	"container/list"
)

func orangesRotting(grid [][]int) int {
	//identify fresh and rotten oranges
	rows, columns := len(grid), len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	fresh := 0
	minutes := 0
	queue := list.New()

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			switch grid[r][c] {
			case 2:
				queue.PushBack([2]int{r, c})
			case 1:
				fresh++
			}
		}
	}

	for queue.Len() > 0 && fresh > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			front := queue.Remove(queue.Front()).([2]int)
			r, c := front[0], front[1]

			// checking all directions and rotting adjacent fresh oranges
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < rows && nc >= 0 && nc < columns && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					queue.PushBack([2]int{nr, nc})
				}
			}
		}

		minutes++

	}

	if fresh == 0 {
		return minutes
	} else {
		return -1
	}

}
