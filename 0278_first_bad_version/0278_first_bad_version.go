package main

var first = 1

func isBadVersion(version int) bool {
	return version >= first
}

//func firstBadVersion(n int) int {
//	left, right := 1, n
//	for left < right {
//		mid := left + (right-left)/2
//		if isBadVersion(mid) {
//			right = mid
//		} else {
//			left = mid + 1
//		}
//	}
//	return left
//}

//func firstBadVersion(n int) int {
//	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
//}

func firstBadVersion(n int) int {
	left, right := 1, n
	ans := n
	for left < right {
		mid := (right-left)/2 + left
		if isBadVersion(mid) {
			ans = mid
			right = mid
		} else {
			left = mid
		}
	}
	return ans
}

func main() {
	first = 2
	println(firstBadVersion(2))
}
