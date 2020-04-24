package _0023

// 2020.04.24
// https://leetcode-cn.com/problems/merge-k-sorted-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeKLists(lists []*ListNode) *ListNode {
//	var h *ListNode
//	for _, l := range lists {
//		h = mergeTwoLists(h, l)
//	}
//	return h
//}

// 使用归并，提升 10 倍效率
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if start > end {
		return nil
	}
	if start == end {
		return lists[start]
	}
	mid := start + (end-start)/2
	return mergeTwoLists(merge(lists, start, mid), merge(lists, mid+1, end))
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
