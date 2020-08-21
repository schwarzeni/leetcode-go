package _0160

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lA, lB := listLen(headA), listLen(headB)
	longL, shortL, fast := headA, headB, lA-lB
	if lB > lA {
		longL, shortL, fast = shortL, longL, -fast
	}
	for fast > 0 {
		longL = longL.Next
		fast--
	}
	for longL != nil {
		if longL == shortL {
			return longL
		}
		longL, shortL = longL.Next, shortL.Next
	}
	return nil
}

func listLen(l *ListNode) int {
	tmp := l
	count := 0
	for tmp != nil {
		count++
		tmp = tmp.Next
	}
	return count
}
