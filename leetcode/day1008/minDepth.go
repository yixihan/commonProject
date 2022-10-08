package main

import "math"

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

var _minDepth = math.MaxInt

func minDepth(root *TreeNode) int {
	_minDepth = math.MaxInt
	if root == nil {
		return 0
	}
	getDepth(root, 1)
	return _minDepth
}

func getDepth(root *TreeNode, depth int) {
	if root.Left == nil && root.Right == nil {
		_minDepth = min(_minDepth, depth)
		return
	}

	if root.Left != nil {
		getDepth(root.Left, depth+1)
	}
	if root.Right != nil {
		getDepth(root.Right, depth+1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
