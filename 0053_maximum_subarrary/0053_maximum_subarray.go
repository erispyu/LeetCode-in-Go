package main

import "math"

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	mid := n/2 - 1
	crossSum := math.MinInt
	if mid+1 < n {
		crossSum = maxCross(nums, 0, mid, n-1)
	}
	leftSum := maxSubArray(nums[:mid+1])
	rightSum := maxSubArray(nums[mid+1:])
	if leftSum >= crossSum && leftSum >= rightSum {
		return leftSum
	}
	if rightSum >= crossSum && rightSum >= leftSum {
		return rightSum
	}
	return crossSum
}

func maxCross(nums []int, low, mid, high int) int {
	leftSum, rightSum := nums[mid], nums[mid+1]
	for i, sum := mid, 0; i >= low; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}
	for j, sum := mid+1, 0; j <= high; j++ {
		sum += nums[j]
		if sum > rightSum {
			rightSum = sum
		}
	}
	return leftSum + rightSum
}

func main() {
	nums := []int{-2, -1}
	println(maxSubArray(nums))
}
