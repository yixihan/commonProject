package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct {
	parent *TreeNode
	depth  int
}

func (n *node) String() string {
	return fmt.Sprintf("parent: %v, depth: %d\n", n.parent.Val, n.depth)
}

func isCousins(root *TreeNode, x int, y int) bool {
	m := make(map[int]*node)
	m[root.Val] = &node{
		parent: nil,
		depth:  0,
	}
	getDepth(m, root, 0)

	return m[x].depth == m[y].depth && m[x].parent != m[y].parent
}

func getDepth(m map[int]*node, root *TreeNode, depth int) {
	if root == nil {
		return
	}
	depth += 1
	if root.Left != nil {
		m[root.Left.Val] = &node{
			parent: root,
			depth:  depth,
		}
		getDepth(m, root.Left, depth)
	}
	if root.Right != nil {
		m[root.Right.Val] = &node{
			parent: root,
			depth:  depth,
		}
		getDepth(m, root.Right, depth)
	}
}
