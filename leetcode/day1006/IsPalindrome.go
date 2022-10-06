package main

import (
	"fmt"
)

func main() {
	println(isPalindrome(1234554321))
}

// 不转化为字符串
func isPalindrome(x int) bool {
	nums := make([]int, 0)
	for i := x ; i > 0; i /= 10 {
		nums = append (nums, i % 10)
	}
	reverseNums := reverse(nums)
	fmt.Println(reverseNums, nums)
	return equals(nums, reverseNums)
}

func reverse(nums []int) []int {
	ans := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		ans = append(ans, nums[i])
	}
	return ans
}

func equals (a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 转化为字符串
//func isPalindrome(x int) bool {
//	str := strconv.Itoa(x)
//	l := 0
//	r := len(str) - 1
//	for l < r {
//		fmt.Println(str[l], str[r])
//		if str[l] != str[r] {
//			return false
//		}
//		l++
//		r--
//	}
//	return true
//}