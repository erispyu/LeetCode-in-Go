package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}

	curr := 0

	for next := 1; next < len(nums); next++ {
		if nums[curr] != nums[next] {
			curr++
			if next - curr > 1 {
				nums[curr] = nums[next]
			}
		}
	}

	return curr + 1
}
