package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	diffMap := make(map[int]int)
	for index, value := range nums {
		diff := target - value
		diffIndex, ok := diffMap[diff]
		if ok && index != diffIndex{
			return []int{index, diffIndex}
		} else {
			diffMap[value] = index
		}
	}
	return nil
}
