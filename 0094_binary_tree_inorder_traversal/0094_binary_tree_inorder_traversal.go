package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	ans := inorderTraversal(root.Left)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal(root.Right)...)
	return ans
}

func inorderTraversalWithHelperFunc(root *TreeNode) []int {
	var ans []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return ans
}

func inorderTraversalIteration(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 { // 深度遍历结束 or 栈为空
		for root != nil { // 一直向左走到底
			stack = append(stack, root)
			root = root.Left
		}
		// 从栈中取出最后一个元素，这时认为其左子树已经遍历完毕
		peak := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, peak.Val)
		// 遍历右子树
		root = peak.Right
	}
	return ans
}
