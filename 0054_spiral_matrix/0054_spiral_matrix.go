package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	total := m * n
	left, right, top, bottom := 0, n-1, 0, m-1
	res := make([]int, total)
	i := 0
	for i < total {
		// left -> right, top++
		for j := left; j <= right; j++ {
			res[i] = matrix[top][j]
			i++
		}
		if i == total {
			break
		}
		top++
		// top -> bottom, right--
		for j := top; j <= bottom; j++ {
			res[i] = matrix[j][right]
			i++
		}
		if i == total {
			break
		}
		right--
		// right -> left, bottom--
		for j := right; j >= left; j-- {
			res[i] = matrix[bottom][j]
			i++
		}
		if i == total {
			break
		}
		bottom--
		// bottom -> top, left++
		for j := bottom; j >= top; j-- {
			res[i] = matrix[j][left]
			i++
		}
		if i == total {
			break
		}
		left++
	}
	return res
}

func main() {
	testCases := []struct {
		matrix [][]int
		want   []int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tc := range testCases {
		got := spiralOrder(tc.matrix)
		for i := 0; i < len(tc.want); i++ {
			if got[i] != tc.want[i] {
				fmt.Printf("want: %v, got: %v\n", tc.want, got)
			}
		}
	}
}
