package main

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

var _maxDepth = 0

func maxDepth(root *TreeNode) int {
	_maxDepth = 0
	Dfs(root, 0)
	return _maxDepth
}

func Dfs(node *TreeNode, depth int) {
	if node == nil {
		_maxDepth = Max(_maxDepth, depth)
		return
	}
	depth += 1
	Dfs(node.Left, depth)
	Dfs(node.Right, depth)
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
