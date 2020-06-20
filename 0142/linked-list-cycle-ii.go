package _0142

// 2020.06.20
// https://leetcode-cn.com/problems/linked-list-cycle-ii/
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	flag := false
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			flag = true
			break
		}
	}
	if !flag {
		return nil
	}
	index := head
	for index != fast {
		fast = fast.Next
		index = index.Next
	}
	return fast
}
