package main

func romanToInt(s string) int {
	m := initMap()
	ans := 0
	length := len(s)

	for i := 0; i < length; i++ {
		if i < length-1 && s[i] == 'I' && s[i+1] == 'V' {
			ans += 4
			i++
		} else if i < length-1 && s[i] == 'I' && s[i+1] == 'X' {
			ans += 9
			i++
		} else if i < length-1 && s[i] == 'X' && s[i+1] == 'L' {
			ans += 40
			i++
		} else if i < length-1 && s[i] == 'X' && s[i+1] == 'C' {
			ans += 90
			i++
		} else if i < length-1 && s[i] == 'C' && s[i+1] == 'D' {
			ans += 400
			i++
		} else if i < length-1 && s[i] == 'C' && s[i+1] == 'M' {
			ans += 900
			i++
		} else {
			ans += m[s[i]]
		}
	}

	return ans
}

func initMap () map[byte]int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	return m
}
