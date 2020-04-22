package _0019

// 2020.04.22
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fn, bn, pn *ListNode = head, head, nil

	for n > 1 {
		fn = fn.Next
		n--
	}
	if fn.Next == nil { // if head node removed
		head = head.Next
		return head
	}
	for { // other case
		pn = bn
		bn = bn.Next
		if fn = fn.Next; fn.Next == nil {
			pn.Next = bn.Next
			return head
		}
	}
}
