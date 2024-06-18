package main

import "fmt"

func main() {
	fmt.Println(generateParenthesisCombinations(3))
	fmt.Println(generateParenthesisCombinations(5))
}

/**
 * @param n: n pairs of parentheses
 * @return: All combinations of well-formed parentheses
 */
func generateParenthesisCombinations(n int) []string {
	var results []string
	var backtrack func(left, right int, path string)

	backtrack = func(left, right int, path string) {
		if len(path) == 2*n {
			results = append(results, path)
			return
		}
		// add left parenthesis at first
		if left < n {
			backtrack(left+1, right, path+"(")
		}
		// add right parenthesis only if there are more left parenthesis waiting to be closed
		if right < left {
			backtrack(left, right+1, path+")")
		}
	}

	backtrack(0, 0, "")
	return results
}
