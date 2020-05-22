package _0106

// 2020.05.22
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	pol := len(postorder)
	n, idx := &TreeNode{Val: postorder[pol-1]}, 0
	for idx < len(inorder) && inorder[idx] != n.Val {
		idx++
	}
	n.Left = buildTree(inorder[:idx], postorder[0:len(inorder[:idx])])
	n.Right = buildTree(inorder[idx+1:], postorder[len(inorder[:idx]):pol-1])
	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
