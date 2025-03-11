package main

// 在一个m*n的棋盘的每一格都放有一个礼物、每个礼物都有一定的价值（价值大于0）。
// 你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
// 给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

func maxValueDP(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = maxVal(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func maxValueDFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	var dfs func(r, c, val int)
	dfs = func(r, c, val int) {
		if r == m-1 && c == n-1 {
			ans = maxVal(ans, val+grid[r][c])
			return
		}

		if r+1 < m {
			dfs(r+1, c, val+grid[r][c])
		}

		if c+1 < n {
			dfs(r, c+1, val+grid[r][c])
		}
	}
	dfs(0, 0, 0)
	return ans
}

func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	grid := [][]int{
		{2, 4, 8, 1, 6, 2, 5},
		{1, 3, 5, 7, 9, 2, 4},
		{4, 6, 7, 8, 1, 3, 5},
		{2, 4, 6, 8, 1, 3, 5},
	}
	println(maxValueDP(grid))
	println(maxValueDFS(grid))
}
