package main

import "fmt"

func main() {
	fmt.Println(isValid("{()}"))
}

func isValid(s string) bool {
	n := len(s)

	if n % 2 == 1 {
		return false
	}

	pairMap := map[byte]byte {
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte

	for i := 0; i < n; i++ {
		pair, ok := pairMap[s[i]]
		if ok {
			if len(stack) == 0 {
				return false
			}
			// need to consider empty stack
			top := stack[len(stack) - 1]
			if top == pair {
				stack = stack[:len(stack) - 1]
				continue
			} else {
				return false
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
