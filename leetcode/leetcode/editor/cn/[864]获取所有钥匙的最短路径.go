package main
//给定一个二维网格 grid ，其中： 
//
// 
// '.' 代表一个空房间 
// '#' 代表一堵 
// '@' 是起点 
// 小写字母代表钥匙 
// 大写字母代表锁 
// 
//
// 我们从起点开始出发，一次移动是指向四个基本方向之一行走一个单位空间。我们不能在网格外面行走，也无法穿过一堵墙。如果途经一个钥匙，我们就把它捡起来。除非我们
//手里有对应的钥匙，否则无法通过锁。 
//
// 假设 k 为 钥匙/锁 的个数，且满足 1 <= k <= 6，字母表中的前 k 个字母在网格中都有自己对应的一个小写和一个大写字母。换言之，每个锁有唯一
//对应的钥匙，每个钥匙也有唯一对应的锁。另外，代表钥匙和锁的字母互为大小写并按字母顺序排列。 
//
// 返回获取所有钥匙所需要的移动的最少次数。如果无法获取所有钥匙，返回 -1 。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：grid = ["@.a.#","###.#","b.A.B"]
//输出：8
//解释：目标是获得所有钥匙，而不是打开所有锁。
// 
//
// 示例 2： 
//
// 
//
// 
//输入：grid = ["@..aA","..B#.","....b"]
//输出：6
// 
//
// 示例 3: 
// 
// 
//输入: grid = ["@Aa"]
//输出: -1 
//
// 
//
// 提示： 
//
// 
// m == grid.length 
// n == grid[i].length 
// 1 <= m, n <= 30 
// grid[i][j] 只含有 '.', '#', '@', 'a'-'f' 以及 'A'-'F' 
// 钥匙的数目范围是 [1, 6] 
// 每个钥匙都对应一个 不同 的字母 
// 每个钥匙正好打开一个对应的锁 
// 
//
// Related Topics 位运算 广度优先搜索 数组 矩阵 👍 148 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	var k, si, sj int
	for i, row := range grid {
		for j, c := range row {
			if c >= 'a' && c <= 'z' {
				// 累加钥匙数量
				k++
			} else if c == '@' {
				// 起点
				si, sj = i, j
			}
		}
	}
	type tuple struct{ i, j, state int }
	q := []tuple{tuple{si, sj, 0}}
	vis := map[tuple]bool{tuple{si, sj, 0}: true}
	dirs := []int{-1, 0, 1, 0, -1}
	ans := 0
	for len(q) > 0 {
		for t := len(q); t > 0; t-- {
			p := q[0]
			q = q[1:]
			i, j, state := p.i, p.j, p.state
			// 找到所有钥匙，返回当前步数
			if state == 1<<k-1 {
				return ans
			}
			// 往四个方向搜索
			for h := 0; h < 4; h++ {
				x, y := i+dirs[h], j+dirs[h+1]
				// 在边界范围内
				if x >= 0 && x < m && y >= 0 && y < n {
					c := grid[x][y]
					// 是墙，或者是锁，但此时没有对应的钥匙，无法通过
					if c == '#' || (c >= 'A' && c <= 'Z' && (state>>(c-'A')&1 == 0)) {
						continue
					}
					nxt := state
					// 是钥匙，更新状态
					if c >= 'a' && c <= 'z' {
						nxt |= 1 << (c - 'a')
					}
					// 此状态未访问过，入队
					if !vis[tuple{x, y, nxt}] {
						vis[tuple{x, y, nxt}] = true
						q = append(q, tuple{x, y, nxt})
					}
				}
			}
		}
		// 步数加一
		ans++
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)
