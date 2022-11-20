package main

import "sort"

//一个序列的 宽度 定义为该序列中最大元素和最小元素的差值。
//
// 给你一个整数数组 nums ，返回 nums 的所有非空 子序列 的 宽度之和 。由于答案可能非常大，请返回对 10⁹ + 7 取余 后的结果。
//
// 子序列 定义为从一个数组里删除一些（或者不删除）元素，但不改变剩下元素的顺序得到的数组。例如，[3,6,2,7] 就是数组 [0,3,1,6,2,2,7]
// 的一个子序列。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,1,3]
//输出：6
//解释：子序列为 [1], [2], [3], [2,1], [2,3], [1,3], [2,1,3] 。
//相应的宽度是 0, 0, 0, 1, 1, 2, 2 。
//宽度之和是 6 。
//
//
// 示例 2：
//
//
//输入：nums = [2]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
//
// Related Topics 数组 数学 排序 👍 181 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func sumSubseqWidths(nums []int) (ans int) {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	n := len(nums)
	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = pow2[i-1] * 2 % mod // 预处理 2 的幂次
	}
	for i, x := range nums {
		ans += (pow2[i] - pow2[n-1-i]) * x // 在题目的数据范围下，这不会溢出
	}
	return (ans%mod + mod) % mod // 注意上面有减法，ans 可能为负数
}

//leetcode submit region end(Prohibit modification and deletion)
