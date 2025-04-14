package main

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n {
			return
		}
		if grid[r][c] == '0' {
			return
		}
		grid[r][c] = '0'
		dfs(r, c+1)
		dfs(r, c-1)
		dfs(r+1, c)
		dfs(r-1, c)
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

func main() {
	testCases := []struct {
		grid [][]byte
		num  int
	}{
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			num: 3,
		},
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			num: 1,
		},
	}

	for _, testCase := range testCases {
		result := numIslands(testCase.grid)
		if result != testCase.num {
			panic("Test case failed")
		}
	}
}
