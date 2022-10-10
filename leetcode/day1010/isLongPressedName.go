package main

func isLongPressedName(name string, typed string) bool {
	n1, n2 := len(name), len(typed)
	if n1 > n2 {
		return false
	}

	i, j := 0, 0
	for i < n1 && j < n2 {
		if name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}

	for j < n2 {
		if j > 0 && typed[j] != typed[j-1] {
			return false
		}
		j++
	}

	return i == n1
}
