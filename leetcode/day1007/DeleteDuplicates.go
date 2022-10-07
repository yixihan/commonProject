package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	m := make(map[int]int)
	node, next := head, head.Next
	m[node.Val] = 1
	for next != nil {
		if _, ok := m[next.Val]; !ok {
			m[next.Val] = 1
			node = next
			next = node.Next
		} else {
			node.Next = next.Next
			next = node.Next
		}
	}

	return head
}
