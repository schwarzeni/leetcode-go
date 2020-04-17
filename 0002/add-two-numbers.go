package _0002

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2020.04.17
// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		incre  = 0
		result = &ListNode{Next: nil}
		next   *ListNode
		pre    *ListNode
	)
	pre = result
	for l1 != nil && l2 != nil {
		r := l1.Val + l2.Val + incre
		next = &ListNode{Val: r % 10, Next: nil}
		incre = r / 10
		pre.Next = next
		pre = next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		r := l1.Val + incre
		next = &ListNode{Val: r % 10, Next: nil}
		incre = r / 10
		pre.Next = next
		pre = next
		l1 = l1.Next
	}
	for l2 != nil {
		r := l2.Val + incre
		next = &ListNode{Val: r % 10, Next: nil}
		incre = r / 10
		pre.Next = next
		pre = next
		l2 = l2.Next
	}
	for incre != 0 {
		r := incre
		next = &ListNode{Val: r % 10, Next: nil}
		incre = r / 10
		pre.Next = next
		pre = next
	}
	return result.Next
}
