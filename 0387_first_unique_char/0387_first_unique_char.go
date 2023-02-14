package main

import "fmt"

func firstUniqCharLoopTwice(s string) int {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for i, c := range s {
		if cnt[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqueCharLoopOnce(s string) int {
	n := len(s)
	record := [26]int{}
	for i := range record[:] {
		record[i] = n
	}
	for i, c := range s {
		if record[c-'a'] == n {
			record[c-'a'] = i
		} else if record[c-'a'] >= 0 {
			record[c-'a'] = -1
		}
	}
	ans := n
	for _, i := range record {
		if i < 0 || i == n {
			continue
		}
		if i < ans {
			ans = i
		}
	}
	if ans == n {
		ans = -1
	}
	return ans
}

func firstUniqCharQueue(s string) int {
	n := len(s)
	record := [26]int{}
	q := []int{}
	for i := range record[:] {
		record[i] = n
	}
	for i, c := range s {
		ch := c - 'a'
		if record[ch] == n {
			record[ch] = i
			q = append(q, i)
		} else {
			record[ch] = -1
			for len(q) > 0 && record[s[q[0]]-'a'] == -1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return q[0]
	} else {
		return -1
	}
}

func main() {
	fmt.Println(firstUniqCharLoopTwice("aadadaad"))
	fmt.Println(firstUniqueCharLoopOnce("aadadaad"))
	fmt.Println(firstUniqCharQueue("aadadaad"))
}
