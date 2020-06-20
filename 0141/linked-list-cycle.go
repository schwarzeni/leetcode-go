package _0141

// 2020.06.20
// https://leetcode-cn.com/problems/linked-list-cycle/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slowPtr, fastPtr := head, head
	for {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if fastPtr == nil || fastPtr.Next == nil {
			return false
		}
		if fastPtr == slowPtr || fastPtr.Next == slowPtr {
			return true
		}
	}
}
