package _0061

// 2020.05.05
// https://leetcode-cn.com/problems/rotate-list/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	size := 1
	n := head
	for n.Next != nil {
		size++
		n = n.Next
	}
	tail := n
	if k = k % size; k == 0 {
		return head
	}
	tail.Next = head
	k, n = size-k-1, head
	for k > 0 {
		n = n.Next
		k--
	}
	head = n.Next
	n.Next = nil
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
