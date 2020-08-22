package _1530

// 2020.08.22
// https://leetcode-cn.com/problems/number-of-good-leaf-nodes-pairs/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	parentNode := make(map[*TreeNode]*TreeNode) // 记录父节点
	isVisited := make(map[*TreeNode]struct{})
	step := 0
	count := 0
	leafCount := 0

	// 先序遍历树的函数
	var iterTree func(root *TreeNode, hook func(curr *TreeNode))
	iterTree = func(root *TreeNode, hook func(curr *TreeNode)) {
		if root == nil {
			return
		}
		hook(root)
		iterTree(root.Left, hook)
		iterTree(root.Right, hook)
	}

	// 从 node 节点开始深度优先遍历整棵树
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		_, visited := isVisited[node]
		if visited || step > distance || node == nil {
			return
		}
		isVisited[node] = struct{}{}
		if node.Left == nil && node.Right == nil {
			count++
		}
		step++
		dfs(node.Left)
		dfs(node.Right)
		dfs(parentNode[node])
		step--
	}

	// 先遍历一遍，记录父节点
	parentNode[root] = nil
	iterTree(root, func(curr *TreeNode) {
		if curr.Left == nil && curr.Right == nil {
			leafCount++
		}
		if curr.Left != nil {
			parentNode[curr.Left] = curr
		}
		if curr.Right != nil {
			parentNode[curr.Right] = curr
		}
	})

	// 从 node 节点开始 dfs 的函数
	iterTree(root, func(curr *TreeNode) {
		if curr.Left == nil && curr.Right == nil {
			step = 0
			isVisited = make(map[*TreeNode]struct{})
			dfs(curr)
		}
	})

	return (count - leafCount) / 2
}
