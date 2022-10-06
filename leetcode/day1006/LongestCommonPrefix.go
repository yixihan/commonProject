package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var minStr string
	minLen := 201

	for _, str := range strs {
		if minLen > len(str) {
			minLen = len(str)
			minStr = str
		}
	}
	var sb strings.Builder
	for i := 0; i < minLen; i++ {
		if helper(sb.String() + string(minStr[i]), strs) {
			sb.WriteString(string(minStr[i]))
		} else {
			break
		}
	}
	return sb.String()
}

func helper (s string, strs []string) bool {
	for _, str := range strs {
		if !strings.HasPrefix(str, s) {
			return false
		}
	}

	return true
}
