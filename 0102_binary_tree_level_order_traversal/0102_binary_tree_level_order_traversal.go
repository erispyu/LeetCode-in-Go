package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	currLevel := []*TreeNode{root}
	for currLevel != nil {
		var level []int
		var nextLevel []*TreeNode
		for _, node := range currLevel {
			level = append(level, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		ans = append(ans, level)
		currLevel = nextLevel
	}
	return ans
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
