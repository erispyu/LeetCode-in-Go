package main

func main() {
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	println(maxAreaOfIsland(grid))
}

var (
	dx = []int{1, 0, 0, -1}
	dy = []int{0, 1, -1, 0}
)

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j, grid)
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}

func dfs(r int, c int, grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if r < 0 || r == m || c < 0 || c == n || grid[r][c] != 1 {
		return 0
	}
	area := 1

	for i := 0; i < 4; i++ {
		nr, nc := r+dx[i], c+dy[i]
		area += dfs(nr, nc, grid)
	}
	return area
}
