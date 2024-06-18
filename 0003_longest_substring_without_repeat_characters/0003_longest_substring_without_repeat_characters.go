package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	indexMap := map[byte]int{}
	left, right, n := 0, 1, len(s)
	indexMap[s[left]] = left
	ans := 0
	for right < n {
		pre, ok := indexMap[s[right]]
		if ok {
			ans = max(ans, right-left)
			newLeft := pre + 1 // 重复字符的下一个位置
			for left < newLeft {
				delete(indexMap, s[left])
				left++
			}
		}
		indexMap[s[right]] = right
		right++
	}
	ans = max(ans, right-left)
	return ans
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func main() {
	testCases := []string{
		"A",
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"ab",
	}
	for _, testCase := range testCases {
		fmt.Println(testCase, lengthOfLongestSubstring(testCase))
	}
}
