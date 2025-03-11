package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	res := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			}
			res += leftMax - height[left]
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			}
			res += rightMax - height[right]
			right--
		}
	}
	return res
}

func main() {
	// [0,1,0,2,1,0,1,3,2,1,2,1]
	// 6
	testCases := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, tc := range testCases {
		got := trap(tc.height)
		if got != tc.want {
			fmt.Printf("failed: %v, %d, %d\n", tc.height, got, tc.want)
		} else {
			fmt.Printf("passed: %v, %d\n", tc.height, got)
		}
	}
}
