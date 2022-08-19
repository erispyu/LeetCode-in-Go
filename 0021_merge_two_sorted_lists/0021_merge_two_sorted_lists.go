package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := ListNode{}
	temp := &dummyHead
	for list1 != nil && list2 != nil {
		val1 := list1.Val
		val2 := list2.Val
		if val1 < val2 {
			temp.Next = list1
			list1 = list1.Next
			temp = temp.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
			temp = temp.Next
		}
	}

	if list1 != nil {
		temp.Next = list1
	}

	if list2 != nil {
		temp.Next = list2
	}

	return dummyHead.Next
}


func mergeRecursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeRecursion(list1.Next, list2)
		return list1
	}

	list2.Next = mergeRecursion(list1, list2.Next)
	return list2
}
