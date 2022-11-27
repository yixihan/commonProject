package main

import (
	"container/heap"
	"math"
)

//给你一个无向图（原始图），图中有 n 个节点，编号从 0 到 n - 1 。你决定将图中的每条边 细分 为一条节点链，每条边之间的新节点数各不相同。
//
// 图用由边组成的二维数组 edges 表示，其中 edges[i] = [ui, vi, cnti] 表示原始图中节点 ui 和 vi 之间存在一条边，
//cnti 是将边 细分 后的新节点总数。注意，cnti == 0 表示边不可细分。
//
// 要 细分 边 [ui, vi] ，需要将其替换为 (cnti + 1) 条新边，和 cnti 个新节点。新节点为 x1, x2, ..., xcnti ，
//新边为 [ui, x1], [x1, x2], [x2, x3], ..., [xcnti+1, xcnti], [xcnti, vi] 。
//
// 现在得到一个 新的细分图 ，请你计算从节点 0 出发，可以到达多少个节点？如果节点间距离是 maxMoves 或更少，则视为 可以到达 。
//
// 给你原始图和 maxMoves ，返回 新的细分图中从节点 0 出发 可到达的节点数 。
//
//
//
// 示例 1：
//
//
//输入：edges = [[0,1,10],[0,2,1],[1,2,2]], maxMoves = 6, n = 3
//输出：13
//解释：边的细分情况如上图所示。
//可以到达的节点已经用黄色标注出来。
//
//
// 示例 2：
//
//
//输入：edges = [[0,1,4],[1,2,6],[0,2,8],[1,3,1]], maxMoves = 10, n = 4
//输出：23
//
//
// 示例 3：
//
//
//输入：edges = [[1,2,4],[1,4,5],[1,3,1],[2,3,4],[3,4,5]], maxMoves = 17, n = 5
//输出：1
//解释：节点 0 与图的其余部分没有连通，所以只有节点 0 可以到达。
//
//
//
//
// 提示：
//
//
// 0 <= edges.length <= min(n * (n - 1) / 2, 10⁴)
// edges[i].length == 3
// 0 <= ui < vi < n
// 图中 不存在平行边
// 0 <= cnti <= 10⁴
// 0 <= maxMoves <= 10⁹
// 1 <= n <= 3000
//
//
// Related Topics 图 最短路 堆（优先队列） 👍 141 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func reachableNodes(edges [][]int, maxMoves, n int) (ans int) {
	g := make([][]neighbor, n)
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		g[u] = append(g[u], neighbor{v, cnt + 1})
		g[v] = append(g[v], neighbor{u, cnt + 1}) // 建图
	}

	dist := dijkstra(g, 0) // 从 0 出发的最短路

	for _, d := range dist {
		if d <= maxMoves { // 这个点可以在 maxMoves 步内到达
			ans++
		}
	}
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		a := max(maxMoves-dist[u], 0)
		b := max(maxMoves-dist[v], 0)
		ans += min(a+b, cnt) // 这条边上可以到达的节点数
	}
	return
}

// 以下为 Dijkstra 算法模板
type neighbor struct{ to, weight int }

func dijkstra(g [][]neighbor, start int) []int {
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.x
		if dist[x] < p.dist {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			if d := dist[x] + e.weight; d < dist[y] {
				dist[y] = d
				heap.Push(&h, pair{y, d})
			}
		}
	}
	return dist
}

type pair struct{ x, dist int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dist < h[j].dist }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
