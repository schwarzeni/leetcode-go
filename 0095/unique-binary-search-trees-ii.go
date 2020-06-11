package _0095

// 2020.06.11
// https://leetcode-cn.com/problems/unique-binary-search-trees-ii
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var res []*TreeNode
	for i := 1; i <= n; i++ {
		res = append(res, generateTrees2(1, i, n)...)
	}
	return res
}

func generateTrees2(l, m, r int) (res []*TreeNode) {
	if l == m && m == r {
		return []*TreeNode{{Val: m}}
	}
	if l == m {
		for nr := m + 1; nr <= r; nr++ {
			for _, rnode := range generateTrees2(m+1, nr, r) {
				res = append(res, &TreeNode{Val: m, Left: nil, Right: rnode})
			}
		}
		return
	}
	if r == m {
		for nl := l; nl < m; nl++ {
			for _, lnode := range generateTrees2(l, nl, m-1) {
				res = append(res, &TreeNode{Val: m, Left: lnode, Right: nil})
			}
		}
		return
	}
	for nl := l; nl < m; nl++ {
		for _, lnode := range generateTrees2(l, nl, m-1) {
			for nr := m + 1; nr <= r; nr++ {
				for _, rnode := range generateTrees2(m+1, nr, r) {
					res = append(res, &TreeNode{Val: m, Left: lnode, Right: rnode})
				}
			}
		}
	}
	return
}
