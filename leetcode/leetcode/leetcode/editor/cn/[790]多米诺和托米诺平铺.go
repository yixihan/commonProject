package main
//有两种形状的瓷砖：一种是 2 x 1 的多米诺形，另一种是形如 "L" 的托米诺形。两种形状都可以旋转。 
//
// 
//
// 给定整数 n ，返回可以平铺 2 x n 的面板的方法的数量。返回对 10⁹ + 7 取模 的值。 
//
// 平铺指的是每个正方形都必须有瓷砖覆盖。两个平铺不同，当且仅当面板上有四个方向上的相邻单元中的两个，使得恰好有一个平铺有一个瓷砖占据两个正方形。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入: n = 3
//输出: 5
//解释: 五种不同的方法如上所示。
// 
//
// 示例 2: 
//
// 
//输入: n = 1
//输出: 1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 1000 
// 
//
// Related Topics 动态规划 👍 229 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var f = [1001]int{1, 1, 2}

func init() {
	for i := 3; i <= 1000; i++ {
		f[i] = (f[i-1]*2 + f[i-3]) % (1e9 + 7)
	}
}

func numTilings(n int) int {
	return f[n]
}
//leetcode submit region end(Prohibit modification and deletion)
