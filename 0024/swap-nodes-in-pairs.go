package _0024

// 2020.04.24
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var (
		pre  *ListNode
		h    = head
		curr = head
	)
	for {
		if curr == nil || curr.Next == nil {
			break
		}
		if pre == nil {
			h = head.Next
		}
		swap(pre, curr, curr.Next)
		pre = curr
		curr = curr.Next
	}
	return h
}

func swap(pre *ListNode, n1 *ListNode, n2 *ListNode) {
	if pre != nil {
		pre.Next = n2
	}
	n1.Next = n2.Next
	n2.Next = n1
}

// 1 -> 2 -> 3 -> 4 -> 5
//
