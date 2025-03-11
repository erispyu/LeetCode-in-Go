package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int

	// 先确定两端，再确定中间
	// x: 可用value的第一个
	// y: 可用value的最后一个
	// z: 可用value的最后一个
	for x := 0; x < n; x++ {
		if x > 0 && nums[x] == nums[x-1] {
			continue
		}
		for z := n - 1; z > x+1; z-- {
			if z < n-1 && nums[z] == nums[z+1] {
				continue
			}
			val := -(nums[x] + nums[z])
			if nums[z-1] == val || nums[x+1] == val {
				ans = append(ans, []int{nums[x], val, nums[z]})
				continue
			}
			if nums[z-1] < val || nums[x+1] > val {
				continue
			}
			y := binarySearch(nums, x+1, z-1, val)
			if y != -1 {
				ans = append(ans, []int{nums[x], nums[y], nums[z]})
			}
		}
	}
	return ans
}

func binarySearch(nums []int, left, right, val int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if nums[mid] == val {
		return mid
	}
	if nums[mid] > val {
		return binarySearch(nums, left, mid-1, val)
	}
	return binarySearch(nums, mid+1, right, val)
}

func threeSumOfficialSolSpeedUp(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] { // 跳过重复数字
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 { // 优化一: 如果当前数加上最小的2个数之和大于0，当前数不可能再有解
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 { // 优化二: 如果当前数加上最大的2个数之和小于0，当前数不可能再有解
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 跳过重复数字
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // 跳过重复数字
			}
		}
	}
	return
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
	nums = []int{-4, 1, 1, 2, 2, 3}
	fmt.Println(threeSum(nums))
	nums = []int{0, 0, 0}
	fmt.Println(threeSum(nums))
	nums = []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(threeSum(nums))
}
