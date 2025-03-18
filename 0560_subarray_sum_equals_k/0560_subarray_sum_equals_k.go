package main

import "fmt"

func subarraySumSlow(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	leftSum := make([]int, len(nums))
	leftSum[0] = nums[0]
	ans := 0
	for i := 1; i < len(nums); i++ {
		leftSum[i] = leftSum[i-1] + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if leftSum[i] == k {
			// [0:i]
			ans++
		}
		for j := i + 1; j < len(nums); j++ {
			sum := leftSum[j] - leftSum[i]
			if sum == k {
				// [i:j]
				ans++
			}
		}
	}
	return ans
}

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	ans, pre := 0, 0
	record := make(map[int]int, 0) // record cnt of prefix sums
	record[0] = 1                  // initial case: no element selected, prefix sum is 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if cnt, ok := record[pre-k]; ok { // pre[j-1] == pre[i] - k, subarray [j, i]
			ans += cnt
		}
		record[pre]++
	}
	return ans
}

func main() {
	testCases := []struct {
		nums   []int
		k      int
		wanted int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 10, 1},
		{[]int{1, 2, 3, 4, 5}, 11, 0},
	}
	for _, tc := range testCases {
		got := subarraySum(tc.nums, tc.k)
		fmt.Printf("got: %v, wanted: %v\n", got, tc.wanted)
	}
}
