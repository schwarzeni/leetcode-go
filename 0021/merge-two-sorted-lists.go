package _0021

// 2020.04.23
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	header := &ListNode{Next: nil}
	h := header

	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val < l2.Val) {
			header.Next = l1
			l1 = l1.Next
		} else {
			header.Next = l2
			l2 = l2.Next
		}
		header = header.Next
	}

	return h.Next
}
