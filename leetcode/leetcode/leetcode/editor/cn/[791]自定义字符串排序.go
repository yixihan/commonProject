package main

import "sort"

//给定两个字符串 order 和 s 。order 的所有单词都是 唯一 的，并且以前按照一些自定义的顺序排序。
//
// 对 s 的字符进行置换，使其与排序的 order 相匹配。更具体地说，如果在 order 中的字符 x 出现字符 y 之前，那么在排列后的字符串中， x 
//也应该出现在 y 之前。 
//
// 返回 满足这个性质的 s 的任意排列 。 
//
// 
//
// 示例 1: 
//
// 
//输入: order = "cba", s = "abcd"
//输出: "cbad"
//解释: 
//“a”、“b”、“c”是按顺序出现的，所以“a”、“b”、“c”的顺序应该是“c”、“b”、“a”。
//因为“d”不是按顺序出现的，所以它可以在返回的字符串中的任何位置。“dcba”、“cdba”、“cbda”也是有效的输出。 
//
// 示例 2: 
//
// 
//输入: order = "cbafg", s = "abcd"
//输出: "cbad"
// 
//
// 
//
// 提示: 
//
// 
// 1 <= order.length <= 26 
// 1 <= s.length <= 200 
// order 和 s 由小写英文字母组成 
// order 中的所有字符都 不同 
// 
//
// Related Topics 哈希表 字符串 排序 👍 156 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func customSortString(order string, s string) string {
	d := [26]int{}
	for i := range order {
		d[order[i]-'a'] = i
	}
	cs := []byte(s)
	sort.Slice(cs, func(i, j int) bool { return d[cs[i]-'a'] < d[cs[j]-'a'] })
	return string(cs)
}
//leetcode submit region end(Prohibit modification and deletion)
