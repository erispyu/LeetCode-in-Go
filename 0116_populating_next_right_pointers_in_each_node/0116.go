package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	println(connect(node1))
}

func connect(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	q := []*Node{root}
	for len(q) > 0 {
		level := q
		q = []*Node{}
		for i, node := range level {
			if i+1 < len(level) {
				node.Next = level[i+1]
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}
