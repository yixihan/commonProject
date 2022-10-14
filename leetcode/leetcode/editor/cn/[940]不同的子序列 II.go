package main
//给定一个字符串 s，计算 s 的 不同非空子序列 的个数。因为结果可能很大，所以返回答案需要对 10^9 + 7 取余 。 
//
// 字符串的 子序列 是经由原字符串删除一些（也可能不删除）字符但不改变剩余字符相对位置的一个新字符串。 
//
// 
// 例如，"ace" 是 "abcde" 的一个子序列，但 "aec" 不是。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "abc"
//输出：7
//解释：7 个不同的子序列分别是 "a", "b", "c", "ab", "ac", "bc", 以及 "abc"。
// 
//
// 示例 2： 
//
// 
//输入：s = "aba"
//输出：6
//解释：6 个不同的子序列分别是 "a", "b", "ab", "ba", "aa" 以及 "aba"。
// 
//
// 示例 3： 
//
// 
//输入：s = "aaa"
//输出：3
//解释：3 个不同的子序列分别是 "a", "aa" 以及 "aaa"。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 2000 
// s 仅由小写英文字母组成 
// 
//
// 
//
// Related Topics 字符串 动态规划 👍 167 👎 0



//leetcode submit region begin(Prohibit modification and deletion)
const mod = 1e9 + 7

func distinctSubseqII(s string) int {
	/*
	   > aba
	   {} <------------------------+       哨兵
	   {}, a                       |       a           pre = [{}]              len(pre) = 1
	   {}, a, ab, b                |       b           pre = [{}, a]           len(pre) = 2
	   {}, a, ab, b, aa, aba, ba, *a       a           pre = [{}, a, ab, b]    len(pre) = 4

	   > abaa
	   {} <------------------------+                                    哨兵
	   {}, a                       |                                    a           pre = [{}]                         len(pre) = 1
	   {}, a, ab, b                |                                    b           pre = [{}, a]                      len(pre) = 2
	   {}, a, ab, b, aa, aba, ba, *a                                    a           pre = [{}, a, ab, b]               len(pre) = 4
	   {}, a, ab, b, aa, aba, ba, *aa, *aba, *ba, aaa, abaa, baa, *a    a           pre = [{}, a, ab, b, aa, aba, ba]  len(pre) = 7
	*/
	// 保存相同字符结尾的重复值，因为cur = pre（之前的值） + pre（之前值组合当前的字符）- 重复值
	cache := [26]int{}
	// 设置一个哨兵
	pre, cur := 1, 0
	for _, c := range s {
		i := c - 'a'
		// 从缓存中取之前重复过的（相同字符结尾的值）
		dup := cache[i]
		// 执行上面的公式
		cur = (pre + pre - dup + mod) % mod
		// 这里存pre的原因是，当前遍历的字符串已经把「前面的字符串(pre)」都组合过一遍了
		// 下一次遇到时，需要把之前组合过的部分给删掉
		cache[i] = pre
		pre = cur
	}
	// 把之前设置的哨兵取消
	return cur - 1
}

//leetcode submit region end(Prohibit modification and deletion)
