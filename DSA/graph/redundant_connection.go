func findRedundantConnection(edges [][]int) []int {
    parents := make(map[int]int)
    for i:= range len(edges) {
        parents[i] = i
    }

    // find parent of the node
    var find func(int)int
    find = func(x int)int {
        if parents[x] != x {
            parents[x] = find(parents[x]) // path compression
        }
        return parents[x]
    } 

    // join nodes
    var union func(int, int) bool
    union = func(x, y int) bool {
        px, py := find(x), find(y)
        if px == py {
            return false
        }
        parents[px] = py
        return true
        
    }

    for _, edge := range edges {
        u, v := edge[0], edge[1]
        if !union(u, v) {
            return edge // redundant edge
        }
    }

    return []int{}
}