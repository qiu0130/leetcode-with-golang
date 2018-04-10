package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}
	p := make(map[string]int, len(s))
	count := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if _, ok := p[string(s[i])]; ok {
			// 重复
			count = Max(p[string(s[i])], count)
		}
		ans = Max(i - count + 1, ans)
		p[string(s[i])] = i + 1
		//fmt.Println(p)
	}
	return ans
}

func main() {

	var tests = []struct {
		s string
	}{
		{
			"aabc",
		},
		{
			"abcabcbb",
		},
		{
			"bbbbb",
		},
		{
			"",
		},
		{
			"au",
		},

	}

	for _, v := range tests {
		fmt.Println(lengthOfLongestSubstring(v.s))
	}
}