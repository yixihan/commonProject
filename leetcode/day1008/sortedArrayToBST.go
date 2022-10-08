package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(0, len(nums), nums)
}

func buildTree(l, r int, nums []int) *TreeNode {
	if l == r {
		return nil
	}

	mid := (l + r) >> 1
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildTree(l, mid, nums)
	root.Right = buildTree(mid+1, r, nums)
	return root
}
