package main

import "fmt"

/*
超运算:
a[1]b=a+b
a[2]b=a[1]a[1]...a[1]a(重复b次，右结合，即运算顺序自右至左)
a[n]b=a[n-1]a[n-1]...a[n-1]a（重复b次，右结合，即运算顺序自右至左）

例：
2[1]3=2+3=5
2[2]3=2[1]2[1]2=2[1]4=6（2x3,此运算表示乘法）
2[3]3=2[2]2[2]2=2[2]4=8（2^3,此预算表示指数）
2[4]3=2[3]2[3]2=2[3]4=2^4=16(2^2^2,此运算表示重幂)

	...
*/
func superOperation(a, b, n int) int {
	if n == 1 {
		return a + b
	}
	if b == 1 {
		return a
	}
	return superOperation(a, superOperation(a, b-1, n), n-1)
}

func main() {
	// 测试超运算函数
	fmt.Println(superOperation(2, 3, 1)) // 5
	fmt.Println(superOperation(2, 3, 2)) // 6
	fmt.Println(superOperation(2, 3, 3)) // 8
	fmt.Println(superOperation(2, 3, 4)) // 16
}
