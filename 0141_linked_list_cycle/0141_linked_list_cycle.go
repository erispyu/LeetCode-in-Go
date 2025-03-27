package main

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

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	// fast meets slow
	return true
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	// fast finds an end
	return false
}
