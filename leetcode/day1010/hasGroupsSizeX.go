package main

func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)

	for _, val := range deck {
		m[val]++
	}

	tmp := 0
	for _, val := range m {
		if tmp = gcd(tmp, val); tmp == 1 {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
