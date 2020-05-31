package _0101

// 2020.05.31
// https://leetcode-cn.com/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	//return same(root, reverse(root))
	if root == nil {
		return true
	}
	return same2(root.Left, root.Right)
}

func reverse(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nroot := &TreeNode{Val: root.Val}
	nroot.Left = reverse(root.Right)
	nroot.Right = reverse(root.Left)
	return nroot
}

func same(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil {
		return t1.Val == t2.Val && same(t1.Left, t2.Left) && same(t1.Right, t2.Right)
	}
	return false
}

func same2(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil {
		return t1.Val == t2.Val && same2(t1.Left, t2.Right) && same2(t1.Right, t2.Left)
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
