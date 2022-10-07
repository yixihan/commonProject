package main

func mySqrt(x int) int {
	l, r := 0, x

	for l <= r {
		mid := (l + r) >> 1
		tmp := mid * mid
		if tmp == x {
			return mid
		} else if tmp < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r
}
