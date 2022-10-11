package main

func canThreePartsEqualSum(arr []int) bool {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	if sum%3 != 0 {
		return false
	}

	l, r := 0, len(arr)-1
	lSum, rSum := 0, 0

	avg := sum / 3
	for l < r {
		if lSum != avg {
			lSum += arr[l]
			l++
		}

		if rSum != avg {
			rSum += arr[r]
			r--
		}

		if lSum == avg && rSum == avg {
			break
		}
	}

	return lSum == avg && rSum == avg
}
