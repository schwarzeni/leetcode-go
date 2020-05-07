package _0572

// 2020.05.07
// https://leetcode-cn.com/problems/subtree-of-another-tree/
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isEqual(s, t) {
		return true
	}
	if s.Left != nil && isSubtree(s.Left, t) {
		return true
	}
	if s.Right != nil && isSubtree(s.Right, t) {
		return true
	}
	return false
}

func isEqual(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s != nil && t != nil {
		return s.Val == t.Val && isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
