package main

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return equals(p, q)
}

func equals(p, q *TreeNode) bool {
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
		return equals(p.Left, q.Left) && equals(p.Right, q.Right)
	} else {
		return false
	}
}
