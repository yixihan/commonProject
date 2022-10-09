package main

func scoreOfParentheses(s string) int {
	ans, n := 0, 0
	flag := false

	for _, c := range s {
		if c == '(' {
			if n == 0 {
				n = 1
			} else {
				n <<= 1
			}
			flag = true
		} else if c == ')' {
			if flag {
				ans += n
			}
			n >>= 1
			flag = false
		}
	}

	return ans
}
