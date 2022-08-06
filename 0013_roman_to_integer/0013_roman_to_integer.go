package main

import "fmt"

func main() {
	fmt.Println(romanToInt("IIIIII"))
}

var romanMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i, c := range s {
		fmt.Println(i, c)
		value := romanMap[s[i]]
		if i < n - 1 && value < romanMap[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}
