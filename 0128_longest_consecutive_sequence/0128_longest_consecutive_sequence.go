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
	// edge case
	if len(nums) <= 1 {
		return len(nums)
	}

	// 1st loop: init hashmap
	record := make(map[int]bool)
	for _, val := range nums {
		record[val] = true
	}

	// 2nd loop: calculate sequences
	ans := 1
	// ENHANCE: loop hashmap instead of nums
	for _, val := range nums {
		if record[val-1] {
			// enhance performance, only start seq with the min val
			continue
		}
		cnt := 0
		for record[val] {
			val++
			cnt++
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func longestConsecutiveAuth(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	println(longestConsecutive(nums))
	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	println(longestConsecutive(nums))
	nums = []int{1, 2, 0, 1} // 排序解决不了
	println(longestConsecutiveError(nums))
	nums = []int{}
	println(longestConsecutive(nums))
}
