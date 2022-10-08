package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	str1 := ParseStr(s)
	str2 := ReverseStr(str1)

	return strings.EqualFold(str1, str2)
}

func ParseStr(s string) string {
	sb := new(strings.Builder)

	for _, c := range strings.ToLower(s) {
		if c >= 'a' && c <= 'z' {
			sb.WriteString(string(c))
		} else if c >= '0' && c <= '9' {
			sb.WriteString(string(c))
		}
	}

	return sb.String()
}

func ReverseStr(s string) string {
	sb := new(strings.Builder)

	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(string(s[i]))
	}

	return sb.String()
}
