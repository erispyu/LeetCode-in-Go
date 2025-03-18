package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	return nil
}

func main() {
	testCases := []struct {
		root   *TreeNode
		wanted [][]int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			wanted: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tc := range testCases {
		got := levelOrder(tc.root)
		fmt.Println(got)
		fmt.Println(tc.wanted)
	}
}
