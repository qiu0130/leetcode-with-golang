package main

import (
	"fmt"
)

//type Trie struct {
//	character []string
//
//}

type Characters struct {
	p map[string]int
	l int
}

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}
	max := 1
	p := make(map[string]int, len(s))
	for i := 0; i < len(s); i++ {
		key := string(s[i])
		if _, ok := p[key]; ok {
			p[key] = p[key] + 1
		} else {
			p[key] = 1
		}
	}

	ch := make([]Characters, len(s))
	ll := len(s)
	for i := 0; i < ll; i++ {
		ch[i].p = p
		ch[i].l = ll-i
	}
	//fmt.Println(ch)

	for i := 1; i < ll; i++ {
		key := string(s[i])
		if _, ok := ch[i].p[key]; ok {
			ch[i].p[key] = ch[i].p[key] - 1
		}
	}
	fmt.Println(ch)

	var length int
	for i := 0; i < ll; i++ {
		key := string(s[i])
		if _, ok := ch[i].p[key]; ok {
			length = ch[i].p[key] - ch[i].l
		} else {
			length = ch[i].p[key]
		}
		if length > max {
			max = length
		}
	}

	return max
}


func main() {

	var tests = []struct {
		s string
	}{
		{
			"aabc",
		},
		//{
		//	"abcabcbb",
		//},
		//{
		//	"bbbbb",
		//},
		//{
		//	"",
		//},
		//{
		//	"au",
		//},

	}

	for _, v := range tests {
		fmt.Println(lengthOfLongestSubstring(v.s))
	}

}