package main

func bitwiseComplement(n int) int {
	tmp := toBinary(n)

	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 0 {
			tmp[i] = 1
		} else if tmp[i] == 1 {
			tmp[i] = 0
		}
	}

	return toDecimal(tmp)
}

func reverse(nums []int) {
	for l, r := 0, len(nums)-1; l < r; {
		tmp := nums[l]
		nums[l] = nums[r]
		nums[r] = tmp
		l++
		r--
	}
}

func toBinary(n int) []int {
	if n < 0 {
		return nil
	} else if n == 0 {
		return []int{0}
	}
	ans := make([]int, 0)
	for i := n; i > 0; i /= 2 {
		ans = append(ans, i%2)
	}
	reverse(ans)
	return ans
}

func toDecimal(arr []int) int {
	ans := 0
	tmp := 1
	for i := len(arr) - 1; i >= 0; i-- {
		ans += arr[i] * tmp
		tmp *= 2
	}

	return ans
}
