package main

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)
}

func dfs(root *TreeNode, sum, targetSum int) bool {
	if root == nil {
		return false
	}

	sum += root.Val
	if root.Left == nil && root.Right == nil && sum == targetSum {
		return true
	}

	return dfs(root.Left, sum, targetSum) && dfs(root.Right, sum, targetSum)
}
