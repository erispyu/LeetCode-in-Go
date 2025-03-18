package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strListToGroup []string) [][]string {
	record := make(map[string][]string)
	for _, s := range strListToGroup {
		k := getKey(s)
		anagrams, ok := record[k]
		if !ok {
			record[k] = []string{s}
		} else {
			record[k] = append(anagrams, s)
		}
	}
	ans := make([][]string, 0, len(record))
	for _, anagrams := range record {
		ans = append(ans, anagrams)
	}
	return ans
}

func getKey(str string) string {
	s := []byte(str)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return string(s)
}

func getKey2(str string) string {
	cl := strings.Split(str, "")
	sort.Strings(cl)
	return strings.Join(cl, "")
}

func getKeyAsArray(str string) [26]int {
	key := [26]int{}
	for _, c := range str {
		key[c-'a']++
	}
	return key
}

func main() {
	strListToGroup := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strListToGroup))
}
