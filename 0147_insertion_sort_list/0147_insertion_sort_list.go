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

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if curr.Val >= lastSorted.Val {
			lastSorted = curr
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}

func main() {
	head := &ListNode{Val: 4}
	head.Append(2)
	head.Append(1)
	head.Append(3)
	head = insertionSortList(head)
	fmt.Println(head.ToSlice())
}
