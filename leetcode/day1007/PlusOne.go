package main

func plusOne(digits []int) []int {
	tmp := 0

	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			digits[i]++
		}
		if tmp > 0 {
			digits[i] += tmp
			tmp = 0
		}
		if digits[i] >= 10 {
			tmp = digits[i] / 10
			digits[i] %= 10
		}
	}
	if tmp > 0 {
		digits = append(digits, digits[len(digits)-1])
		for i := len(digits) - 2; i >= 0; i-- {
			if i != 0 {
				digits[i+1] = digits[i]
			} else {
				digits[i] = tmp
			}
		}
	}

	return digits
}
