package _0105

// 2020.05.22
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 性质
// 先序遍历：[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
// 中序遍历：[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果]]
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	n, idx := &TreeNode{Val: preorder[0]}, 0
	for idx < len(inorder) && inorder[idx] != preorder[0] {
		idx++
	}
	n.Left = buildTree(preorder[1:len(inorder[:idx])+1], inorder[:idx])
	n.Right = buildTree(preorder[len(inorder[:idx])+1:], inorder[idx+1:])
	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
