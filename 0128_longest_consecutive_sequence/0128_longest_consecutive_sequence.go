package main

import "sort"

func longestConsecutiveError(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	sort.Ints(nums)
	temp, max := 1, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if max < temp {
				max = temp
			}
			temp = 1
		} else {
			temp++
		}
	}
	// 退出循环时再更新一下max，避免统计不到一直连续到数组最后一位的情况
	if max < temp {
		max = temp
	}
	return max
}

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	longest := 0
	for num := range cnt {
		if cnt[num-1] >= 1 {
			// not the start of a seq
			continue
		}
		currNum := num
		currLen := 1
		for cnt[currNum+1] >= 1 {
			currNum++
			currLen++
		}
		// seq ended
		if currLen > longest {
			longest = currLen
		}
	}
	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	println(longestConsecutive(nums))
	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	println(longestConsecutive(nums))
	nums = []int{1, 2, 0, 1} // 排序解决不了
	println(longestConsecutive(nums))
}
