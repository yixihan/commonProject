package main

import "math"

//给你一个长度为 n 的整数数组 nums ，请你返回 nums 中最 接近 0 的数字。如果有多个答案，请你返回它们中的 最大值 。
//
//
//
// 示例 1：
//
// 输入：nums = [-4,-2,1,4,8]
//输出：1
//解释：
//-4 到 0 的距离为 |-4| = 4 。
//-2 到 0 的距离为 |-2| = 2 。
//1 到 0 的距离为 |1| = 1 。
//4 到 0 的距离为 |4| = 4 。
//8 到 0 的距离为 |8| = 8 。
//所以，数组中距离 0 最近的数字为 1 。
//
//
// 示例 2：
//
// 输入：nums = [2,-1,1]
//输出：1
//解释：1 和 -1 都是距离 0 最近的数字，所以返回较大值 1 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 1000
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics 数组 👍 7 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func findClosestNumber(nums []int) int {
	ans := 0
	minDiff := math.MaxInt

	var max func(a, b int) int
	max = func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		} else {
			return num2
		}
	}

	for _, val := range nums {
		num := abs(val)
		if num < minDiff {
			minDiff = num
			ans = val
		} else if num == minDiff {
			ans = max(ans, val)
		}
	}

	return ans
}

func abs(num int) int {
	if num > 0 {
		return num
	} else {
		return -num
	}
}

//leetcode submit region end(Prohibit modification and deletion)
