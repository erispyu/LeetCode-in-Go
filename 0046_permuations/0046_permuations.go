package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	visited := make([]bool, n)
	var backtrack func(path []int)
	backtrack = func(path []int) {
		pathLen := len(path) // use the length of current path to record progress, no need for another param
		if pathLen == n {
			perm := make([]int, n)
			copy(perm, path)
			ans = append(ans, perm)
			return
		}
		for i, val := range nums {
			if !visited[i] {
				// forward
				visited[i] = true
				path = append(path, val)
				backtrack(path)
				// backward
				visited[i] = false
				path = path[:len(path)-1] // remove nums[i] from path
			}
		}
	}
	backtrack([]int{})
	return ans
}
