package main

import "strings"

func isValid(s string) bool {
	stack := make([]string, 0)

	for _, v := range strings.Split(s, "") {
		if v == "(" || v == "[" || v == "{" {
			stack = append(stack, v)
		} else if len(stack) > 0 && v == ")" {
			if stack[len(stack)-1] != "(" {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if len(stack) > 0 && v == "]" {
			if stack[len(stack)-1] != "[" {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if len(stack) > 0 && v == "}" {
			if stack[len(stack)-1] != "{" {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}
