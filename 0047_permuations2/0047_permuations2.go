package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	vis := make([]bool, n)

	ans := [][]int{}
	temp := []int{}
	var backtrack func(int)

	backtrack = func(cur int) {
		if cur == n {
			perm := make([]int, n)
			copy(perm, temp)
			ans = append(ans, perm)
		}
		for i := 0; i < n; i++ {
			if vis[i] || i > 0 && nums[i-1] == nums[i] && !vis[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			vis[i] = true
			backtrack(cur + 1)
			temp = temp[:len(temp)-1]
			vis[i] = false
		}
	}

	backtrack(0)
	return ans
}
