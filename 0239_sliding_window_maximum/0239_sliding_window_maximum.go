package main

import (
	"container/heap"
	"fmt"
	"sort"
)

var a []int

// import "container/heap"
//
//	type Interface interface {
//		sort.Interface
//		Push(x any) // add x as element Len()
//		Pop() any   // remove and return element Len() - 1.
//	}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

type TestCase struct {
	nums []int
	k    int
	ans  []int
}

func main() {
	testCases := []TestCase{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{3, 2, 1, 0, -1}, 3, []int{3, 2, 1}},
		{[]int{1}, 1, []int{1}},
	}
	for _, testCase := range testCases {
		ans := maxSlidingWindow(testCase.nums, testCase.k)
		fmt.Println(testCase, ans)
	}
}
