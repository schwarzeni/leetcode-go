package _0109

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	pre, p, q := head, head, head
	for q != nil && q.Next != nil {
		pre = p
		p = p.Next
		q = q.Next.Next
	}
	root := &TreeNode{Val: p.Val}
	if head.Next == nil {
		return root
	}
	pre.Next = nil
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(p.Next)
	return root
}
