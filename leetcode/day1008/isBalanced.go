package main

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

func isBalanced(root *TreeNode) bool {
	return Helper(root)
}

func Helper(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := GetHeight(root.Left, 0)
	rightHeight := GetHeight(root.Right, 0)

	if Abs(leftHeight, rightHeight) > 1 {
		return false
	} else if leftHeight == 0 && rightHeight == 0 {
		return true
	}

	return Helper(root.Left) && Helper(root.Right)
}

func GetHeight(root *TreeNode, height int) int {
	if root == nil {
		return height
	}

	return Max(GetHeight(root.Left, height+1), GetHeight(root.Right, height+1))
}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
