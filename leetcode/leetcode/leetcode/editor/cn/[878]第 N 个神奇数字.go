package main

import "sort"

//一个正整数如果能被 a 或 b 整除，那么它是神奇的。
//
// 给定三个整数 n , a , b ，返回第 n 个神奇的数字。因为答案可能很大，所以返回答案 对 10⁹ + 7 取模 后的值。 
//
// 
//
// 
// 
//
// 示例 1： 
//
// 
//输入：n = 1, a = 2, b = 3
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：n = 4, a = 2, b = 3
//输出：6
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 10⁹ 
// 2 <= a, b <= 4 * 10⁴ 
// 
//
// 
//
// Related Topics 数学 二分查找 👍 125 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func nthMagicalNumber(n int, a int, b int) int {
	c := a * b / gcd(a, b)
	const mod int = 1e9 + 7
	r := (a + b) * n
	return sort.Search(r, func(x int) bool { return x/a+x/b-x/c >= n }) % mod
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
//leetcode submit region end(Prohibit modification and deletion)
