package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	vis := make([]int, n)
	temp := []int{}

	var dfs func(int)
	dfs = func(cur int) {
		if cur == n {
			perm := make([]int, n)
			copy(perm, temp)
			ans = append(ans, perm)
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] == 0 {
				temp = append(temp, nums[i])
				vis[i] = 1
				dfs(cur + 1)
				vis[i] = 0
				temp = temp[:len(temp)-1]
			}
		}
	}

	dfs(0)

	return ans
}
