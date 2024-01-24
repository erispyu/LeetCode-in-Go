package main

import "fmt"

// gcdEuclidean Euclidean algorithm, 辗转相除法
func gcdEuclidean(a, b int) int {
	fmt.Printf("a: %d\tb: %d\n", a, b)
	if b == 0 {
		return a
	}
	return gcdEuclidean(b, a%b)
}

// gcdSubtraction 更相减损术
func gcdSubtraction(a, b int) int {
	fmt.Printf("a: %d\tb: %d\n", a, b)
	if a == b {
		return a
	}
	if a < b {
		return gcdSubtraction(b-a, a)
	}
	return gcdSubtraction(a-b, b)
}

func main() {
	//fmt.Printf("gcd: %d", gcdEuclidean(1071, 462))
	fmt.Printf("gcd: %d", gcdSubtraction(1071, 462))
}
