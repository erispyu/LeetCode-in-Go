package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	record := make(map[string][]string, 0)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sorted := string(s)
		record[sorted] = append(record[sorted], str)
	}
	ans := make([][]string, 0, len(record))
	for _, group := range record {
		ans = append(ans, group)
	}
	return ans
}

type PayIndex struct {
	SourceId    uint64
	ReferenceId uint64
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	var err error
	fmt.Println(err)
	fmt.Println(err == nil)
	var payIndex *PayIndex
	fmt.Println(payIndex)
	fmt.Println(payIndex == nil)
}
