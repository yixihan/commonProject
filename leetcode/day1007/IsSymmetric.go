package main

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs(root.Left, root.Right)
}

func dfs(p, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	} else {
		return false
	}
}
