package main

//给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。
//
// 岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。
//
//
//
// 你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。
//
//
//
// 返回必须翻转的 0 的最小数目。
//
//
//
// 示例 1：
//
//
//输入：grid = [[0,1],[1,0]]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
//输出：2
//
//
// 示例 3：
//
//
//输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
//输出：1
//
//
//
//
// 提示：
//
//
// n == grid.length == grid[i].length
// 2 <= n <= 100
// grid[i][j] 为 0 或 1
// grid 中恰有两个岛
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 矩阵 👍 354 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func shortestBridge(grid [][]int) int {
	var landArr [][]int
	var findLand func(int, int)
	var bridgeNum int
	var endNum int
	var queue [][]int
	findLand = func(i int, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid) {
			return
		}
		if grid[i][j] == 1 {
			grid[i][j] = 2
			landArr = append(landArr, []int{i, j})
			findLand(i+1, j)
			findLand(i-1, j)
			findLand(i, j+1)
			findLand(i, j-1)
		} else if grid[i][j] == 0 {
			queue = append(queue, []int{i, j})
			return
		}
		return
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				findLand(i, j)
				goto Loop
			}
		}
	}
Loop:
	var findNext func(int, int)
	findNext = func(i int, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return
		}
		if grid[i][j] == 0 {
			grid[i][j] = 2
			queue = append(queue, []int{i, j})
		} else {
			if grid[i][j] == 1 {
				endNum = bridgeNum
			}
		}
		return
	}
	for len(queue) > 0 {
		if endNum > 0 {
			break
		}
		bridgeNum++
		for tmpLen := len(queue) - 1; tmpLen >= 0; tmpLen-- {
			tmpi := queue[0][0]
			tmpj := queue[0][1]
			findNext(tmpi+1, tmpj)
			findNext(tmpi-1, tmpj)
			findNext(tmpi, tmpj+1)
			findNext(tmpi, tmpj-1)
			queue = queue[1:]
		}
	}

	return endNum
}

//leetcode submit region end(Prohibit modification and deletion)
