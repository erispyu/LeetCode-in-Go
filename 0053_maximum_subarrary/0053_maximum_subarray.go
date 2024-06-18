package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	mid := n / 2
	leftSum := maxSubArray(nums[:mid])
	rightSum := maxSubArray(nums[mid:])
	crossSum := maxCross(nums, mid)

	// 最大值要使用>=短路
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftSum
	}
	if rightSum >= leftSum && rightSum >= crossSum {
		return rightSum
	}
	return crossSum
}

func maxCross(nums []int, mid int) int {
	n := len(nums)
	if mid >= n {
		// 元素可以为负数，所以不能返回0
		return math.MinInt
	}
	maxLeft, maxRight := 0, 0
	for i, sumLeft := mid-1, 0; i >= 0; i-- {
		sumLeft += nums[i]
		if sumLeft > maxLeft {
			maxLeft = sumLeft
		}
	}
	for j, sumRight := mid+1, 0; j < n; j++ {
		sumRight += nums[j]
		if sumRight > maxRight {
			maxRight = sumRight
		}
	}
	return nums[mid] + maxLeft + maxRight
}

func main() {
	testCases := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{-2, 0, -1},
		{0, -2, 0}, // 最大值要使用>=短路
		{-1, -2},
		{0},
	}
	for _, testCase := range testCases {
		fmt.Println(testCase, maxSubArray(testCase))
	}
}
