package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (i *ListNode) Append(val int) {
	temp := i
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = &ListNode{Val: val}
}

func (i *ListNode) ToSlice() []int {
	temp := i
	var res []int
	for temp != nil {
		res = append(res, temp.Val)
		temp = temp.Next
	}
	return res
}

func isPalindrome(head *ListNode) bool {
	fast, mid := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		mid = mid.Next
	}
	left, right := head, reverseList(mid)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func main() {
	head := &ListNode{Val: 1}
	head.Append(2)
	head.Append(2)
	head.Append(1)
	fmt.Println(head.ToSlice(), isPalindrome(head))

	head = &ListNode{Val: 1}
	head.Append(2)
	head.Append(1)
	fmt.Println(head.ToSlice(), isPalindrome(head))

	head = &ListNode{Val: 1}
	head.Append(2)
	head.Append(3)
	head.Append(2)
	head.Append(1)
	fmt.Println(head.ToSlice(), isPalindrome(head))
}
