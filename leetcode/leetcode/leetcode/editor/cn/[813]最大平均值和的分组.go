package main

import "math"

//给定数组 nums 和一个整数 k 。我们将给定的数组 nums 分成 最多 k 个相邻的非空子数组 。 分数 由每个子数组内的平均值的总和构成。
//
// 注意我们必须使用 nums 数组中的每一个数进行分组，并且分数不一定需要是整数。 
//
// 返回我们所能得到的最大 分数 是多少。答案误差在 10⁻⁶ 内被视为是正确的。 
//
// 
//
// 示例 1: 
//
// 
//输入: nums = [9,1,2,3,9], k = 3
//输出: 20.00000
//解释: 
//nums 的最优分组是[9], [1, 2, 3], [9]. 得到的分数是 9 + (1 + 2 + 3) / 3 + 9 = 20. 
//我们也可以把 nums 分成[9, 1], [2], [3, 9]. 
//这样的分组得到的分数为 5 + 2 + 6 = 13, 但不是最大值.
// 
//
// 示例 2: 
//
// 
//输入: nums = [1,2,3,4,5,6,7], k = 4
//输出: 20.50000
// 
//
// 
//
// 提示: 
//
// 
// 1 <= nums.length <= 100 
// 1 <= nums[i] <= 10⁴ 
// 1 <= k <= nums.length 
// 
//
// Related Topics 数组 动态规划 前缀和 👍 256 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	s := make([]int, n+1)
	f := [110][110]float64{}
	for i, v := range nums {
		s[i+1] = s[i] + v
	}
	var dfs func(i, k int) float64
	dfs = func(i, k int) float64 {
		if i == n {
			return 0
		}
		if k == 1 {
			return float64(s[n]-s[i]) / float64(n-i)
		}
		if f[i][k] > 0 {
			return f[i][k]
		}
		var ans float64
		for j := i; j < n; j++ {
			t := float64(s[j+1]-s[i])/float64(j-i+1) + dfs(j+1, k-1)
			ans = math.Max(ans, t)
		}
		f[i][k] = ans
		return ans
	}
	return dfs(0, k)
}
//leetcode submit region end(Prohibit modification and deletion)
