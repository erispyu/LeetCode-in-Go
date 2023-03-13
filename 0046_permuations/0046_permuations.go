package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	visited := make([]bool, n)
	var backtrack func(next int, path []int)
	backtrack = func(next int, path []int) {
		if next == n {
			perm := make([]int, n)
			copy(perm, path)
			ans = append(ans, perm)
			return
		}
		for i, val := range nums {
			if !visited[i] {
				path = append(path, val)
				visited[i] = true
				backtrack(next+1, path)
				visited[i] = false
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0, []int{})

	return ans
}
