package _0102

// 2020.05.13
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []TreeNodeInfo{{Node: root, Level: 1}}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if len(result) < n.Level {
			result = append(result, []int{})
		}
		result[n.Level-1] = append(result[n.Level-1], n.Node.Val)
		if n.Node.Left != nil {
			queue = append(queue, TreeNodeInfo{Node: n.Node.Left, Level: n.Level + 1})
		}
		if n.Node.Right != nil {
			queue = append(queue, TreeNodeInfo{Node: n.Node.Right, Level: n.Level + 1})
		}
	}
	return result
}

type TreeNodeInfo struct {
	Node  *TreeNode
	Level int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
