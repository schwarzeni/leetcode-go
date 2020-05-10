package _0236

// 2020.05.10
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var (
		pPre, qPre []*TreeNode
		iterF      func(n, target *TreeNode, pre *[]*TreeNode) bool
	)
	iterF = func(n, target *TreeNode, pre *[]*TreeNode) bool {
		if n == nil {
			return false
		}
		if n.Val == target.Val || iterF(n.Left, target, pre) || iterF(n.Right, target, pre) {
			*pre = append(*pre, n)
			return true
		}
		return false
	}
	iterF(root, p, &pPre)
	iterF(root, q, &qPre)
	pi, qi := len(pPre)-1, len(qPre)-1
	for pi >= 0 && qi >= 0 {
		if pPre[pi].Val != qPre[qi].Val {
			break
		}
		pi--
		qi--
	}
	return pPre[pi+1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
