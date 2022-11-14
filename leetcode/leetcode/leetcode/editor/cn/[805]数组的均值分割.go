package main
//给定你一个整数数组
// nums 
//
// 我们要将
// nums 数组中的每个元素移动到 A 数组 或者 B 数组中，使得 A 数组和
// B 数组不为空，并且
// average(A) == average(B) 。 
//
// 如果可以完成则返回true ， 否则返回 false 。 
//
// 注意：对于数组
// arr , 
// average(arr) 是
// arr 的所有元素除以
// arr 长度的和。 
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [1,2,3,4,5,6,7,8]
//输出: true
//解释: 我们可以将数组分割为 [1,4,5,8] 和 [2,3,6,7], 他们的平均值都是4.5。
// 
//
// 示例 2: 
//
// 
//输入: nums = [3,1]
//输出: false
// 
//
// 
//
// 提示: 
//
// 
// 1 <= nums.length <= 30 
// 0 <= nums[i] <= 10⁴ 
// 
//
// Related Topics 位运算 数组 数学 动态规划 状态压缩 👍 137 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}
	s := 0
	for _, v := range nums {
		s += v
	}
	for i, v := range nums {
		nums[i] = v*n - s
	}
	m := n >> 1
	vis := map[int]bool{}
	for i := 1; i < 1<<m; i++ {
		t := 0
		for j, v := range nums[:m] {
			if (i >> j & 1) == 1 {
				t += v
			}
		}
		if t == 0 {
			return true
		}
		vis[t] = true
	}
	for i := 1; i < 1<<(n-m); i++ {
		t := 0
		for j, v := range nums[m:] {
			if (i >> j & 1) == 1 {
				t += v
			}
		}
		if t == 0 || (i != (1<<(n-m))-1 && vis[-t]) {
			return true
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
