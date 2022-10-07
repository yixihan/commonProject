package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans = make([]int, 0)

func inorderTraversal(root *TreeNode) []int {
	ans = make([]int, 0)
	if root != nil {
		inorder(root)
	}
	return ans
}

func inorder(node *TreeNode) {
	if node.Left != nil {
		inorder(node.Left)
	}

	ans = append(ans, node.Val)

	if node.Right != nil {
		inorder(node.Right)
	}
}
