package main

import "strings"

func lengthOfLastWord(s string) int {
	splits := strings.Split(strings.Trim(s, " "), " ")

	return len(splits[len(splits)-1])

}
