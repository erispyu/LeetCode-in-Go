package main

import "fmt"

func main() {
	fmt.Println()
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for i := range strs[0] {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || c != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
