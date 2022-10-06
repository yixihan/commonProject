package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	node1 := list1
	node2 := list2

	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			node.Next = node1
			node1 = node1.Next
		} else {
			node.Next = node2
			node2 = node2.Next
		}
	}

	if node1 != nil {
		node.Next = node1
	}
	if node2 != nil {
		node.Next = node2
	}
	return head.Next
}
