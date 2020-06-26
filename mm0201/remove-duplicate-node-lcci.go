package mm0201

// 2020.06.26
// https://leetcode-cn.com/problems/remove-duplicate-node-lcci/
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cache := make(map[int]struct{})
	cache[head.Val] = struct{}{}
	pre, curr := head, head.Next
	for curr != nil {
		next := curr.Next
		if _, ok := cache[curr.Val]; !ok {
			pre.Next = curr
			pre = curr
			pre.Next = nil
			cache[curr.Val] = struct{}{}
		}
		curr = next
	}
	return head
}
