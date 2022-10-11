package main

import "fmt"

func main() {
	fmt.Println(areAlmostEqual("bank", "kanb"))
}

func areAlmostEqual(s1 string, s2 string) bool {
	var c1, c2 byte = 0, 0
	flag := true

	for i := 0; i < len(s1); i++ {
		if c1 == c2 && s1[i] != s2[i] {
			if flag {
				c1 = s1[i]
				c2 = s2[i]
				continue
			} else {
				return false
			}
		}
		if c1 != c2 && s1[i] != s2[i] {
			if c1 == s2[i] && c2 == s1[i] {
				flag = false
				c1 = 0
				c2 = 0
			} else {
				return false
			}
		}
	}

	return c1 == c2
}
