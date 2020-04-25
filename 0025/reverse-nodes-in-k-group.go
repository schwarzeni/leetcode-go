package _0025

// 2020.04.25
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		h     = &ListNode{Next: head}
		pre   = h
		curr  = head
		count = 0
		gHead = head
	)

	if k == 1 {
		return head
	}

	for curr != nil {
		count++
		if count == k {
			curr = reverseGroup(pre, gHead, curr)
			pre = curr
			gHead = curr.Next
			count = 0
		}
		curr = curr.Next
	}

	return h.Next
}

func reverseGroup(pre, head, tail *ListNode) (t *ListNode) {
	c := head
	p := pre
	for {
		tmp := c.Next
		c.Next = p
		p = c
		c = tmp
		if c == tail {
			break
		}
	}
	head.Next = tail.Next
	tail.Next = p
	pre.Next = tail

	return head
}
