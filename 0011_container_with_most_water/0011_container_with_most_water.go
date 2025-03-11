package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left <= right {
		curr := (right - left) * minVal(height[left], height[right])
		ans = maxVal(ans, curr)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type testCase struct {
	height []int
	want   int
}

func main() {
	testCases := []testCase{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{height: []int{1, 1}, want: 1},
	}
	for _, tc := range testCases {
		got := maxArea(tc.height)
		if got != tc.want {
			print("FAIL: got: %v, want: %v\n", got, tc.want)
			return
		}
	}
	print("PASS\n")
}
