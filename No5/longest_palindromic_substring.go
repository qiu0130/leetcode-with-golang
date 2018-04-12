package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	l := len(s)
	k := int(l / 2)
	for i := 0; i < k; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {

	strLen := len(s)
	if strLen == 1 {
		return s
	}
	if strLen == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}
	if strLen == 3 {
		if s[0] == s[2] {
			return s
		}
		if s[0] == s[1] {
			return s[:2]
		}
		if s[1] == s[2] {
			return s[1:]
		}
	}
	if isPalindrome(s) {
		return s
	}
	max := 0
	answer := s

	for i := 1; i < strLen-1; i++ {
		tmp := ""
		j, k := i-1, i
		for ; ; j, k = j-1, k+1 {
			if j < 0 || k >= strLen || s[j] != s[k] {
				break
			}
			tmp = string(s[j]) + tmp + string(s[k])
		}
		count := len(tmp)
		if count > max {
			max = count
			answer = tmp
		}

		tmp = ""
		j, k = i-1, i+1
		for ; ; j, k = j-1, k+1 {
			if j < 0 || k >= strLen || s[j] != s[k] {
				break
			}
			tmp = string(s[j]) + tmp + string(s[k])
		}
		count = len(tmp)+1
		if count > max {
			max = count
			ll := int(len(tmp)/2)
			answer = tmp[:ll] + string(s[i]) + tmp[ll:]

		}
	}
	return answer
}


func main() {
	var tests  = []string {
		"q",
		"aa",
		"ab",
		"aba",
		"aab",
		"abb",
		"babad",
		"cbbd",
		"aaaaaa",
		"aaraa",
		"aaataacaa",
		"ddtattarrattatddetartrateedredivi",
		"ddtattarrattatdd",
		"aaata",
		"aapttaattp",
		"afqtqtgadaeq2tgafaaaabaaacccadadadadadcccccccccccccccccccccdadadadd",
		"caba",
		"dertyjhgdfytdaba",
		"azwdzwmwcqzgcobeeiphemqbjtxzwkhiqpbrprocbppbxrnsxnwgikiaqutwpftbiinlnpyqstkiqzbggcsdzzjbrkfmhgtnbujzszxsycmvipjtktpebaafycngqasbbhxaeawwmkjcziybxowkaibqnndcjbsoehtamhspnidjylyisiaewmypfyiqtwlmejkpzlieolfdjnxntonnzfgcqlcfpoxcwqctalwrgwhvqvtrpwemxhirpgizjffqgntsmvzldpjfijdncexbwtxnmbnoykxshkqbounzrewkpqjxocvaufnhunsmsazgibxedtopnccriwcfzeomsrrangufkjfzipkmwfbmkarnyyrgdsooosgqlkzvorrrsaveuoxjeajvbdpgxlcrtqomliphnlehgrzgwujogxteyulphhuhwyoyvcxqatfkboahfqhjgujcaapoyqtsdqfwnijlkknuralezqmcryvkankszmzpgqutojoyzsnyfwsyeqqzrlhzbc",

	}
	for _, test := range tests {
		fmt.Println(longestPalindrome(test))
	}
}