package main

import "fmt"

func allLengthPermute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	vis := make([]int, n)
	var temp []int

	var dfs func(cur, k int)
	dfs = func(cur, k int) {
		if cur == k {
			perm := make([]int, k)
			copy(perm, temp)
			ans = append(ans, perm)
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] == 0 {
				temp = append(temp, nums[i])
				vis[i] = 1
				dfs(cur+1, k)
				vis[i] = 0
				temp = temp[:len(temp)-1]
			}
		}
	}

	for k := 1; k <= n; k++ {
		dfs(0, k)
	}

	return ans
}

func main() {
	fmt.Println(allLengthPermute([]int{1, 2, 3}))
}
