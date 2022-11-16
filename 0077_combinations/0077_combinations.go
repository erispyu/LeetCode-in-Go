package main

import "fmt"

var ans [][]int

func combine(n int, k int) [][]int {
	dfs(1, n, k, []int{}, ans)
	return ans
}

func dfs(cur int, n int, k int, temp []int, ans [][]int) {
	if len(temp)+n-cur+1 < k {
		return
	}
	if len(temp) == k {
		ans = append(ans, temp)
		return
	}
	temp = append(temp, cur)
	dfs(cur+1, n, k, temp, ans)
	temp = temp[:len(temp)-1]
	dfs(cur+1, n, k, temp, ans)
	return
}

func main() {
	fmt.Println(combine(4, 2))
}
