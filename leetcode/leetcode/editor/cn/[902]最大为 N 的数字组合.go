package main

import "strconv"

//给定一个按 非递减顺序 排列的数字数组
// digits 。你可以用任意次数 digits[i] 来写的数字。例如，如果
// digits = ['1','3','5']，我们可以写数字，如
// '13', '551', 和 '1351315'。 
//
// 返回 可以生成的小于或等于给定整数 n 的正整数的个数 。 
//
// 
//
// 示例 1： 
//
// 
//输入：digits = ["1","3","5","7"], n = 100
//输出：20
//解释：
//可写出的 20 个数字是：
//1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
// 
//
// 示例 2： 
//
// 
//输入：digits = ["1","4","9"], n = 1000000000
//输出：29523
//解释：
//我们可以写 3 个一位数字，9 个两位数字，27 个三位数字，
//81 个四位数字，243 个五位数字，729 个六位数字，
//2187 个七位数字，6561 个八位数字和 19683 个九位数字。
//总共，可以使用D中的数字写出 29523 个整数。 
//
// 示例 3: 
//
// 
//输入：digits = ["7"], n = 8
//输出：1
// 
//
// 
//
// 提示： 
// 
//
// 
// 1 <= digits.length <= 9 
// digits[i].length == 1 
// digits[i] 是从 '1' 到 '9' 的数 
// digits 中的所有值都 不同 
// digits 按 非递减顺序 排列 
// 1 <= n <= 10⁹ 
// 
//
// Related Topics 数组 数学 字符串 二分查找 动态规划 👍 130 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func atMostNGivenDigitSet(digits []string, n int) int {
	nS := strconv.Itoa(n)
	ln, ld := len(nS), len(digits)
	res := 0
	res += expSum(ld, ln-1)
	res += recursion(digits, []byte(nS))
	return res
}
func recursion(digits []string, d []byte) (ans int) {
	if len(d) == 0 {
		return 1
	}
	nLess, ans := 0, 0
	for _, digit := range digits {
		if digit < string(d[0]) {
			nLess += 1
		} else if digit == string(d[0]) {
			ans += recursion(digits, d[1:])
		}
	}
	ans += nLess * exp(len(digits), len(d)-1)
	return ans
}
func expSum(a, b int) int {
	last, r := 1, 0
	for i := 0; i < b; i++ {
		last *= a
		r += last
	}
	return r
}
func exp(a, b int) int {
	r := 1
	for i := 0; i < b; i++ {
		r *= a
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)
