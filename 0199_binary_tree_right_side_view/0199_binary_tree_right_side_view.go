package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	q := []*TreeNode{root}
	for len(q) > 0 {
		var p []*TreeNode
		for _, node := range q {
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		res = append(res, q[len(q)-1].Val)
		q = p
	}
	return res
}

func main() {
	// Example usage
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	result := rightSideView(root)
	fmt.Println(result)
}
