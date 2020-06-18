package _1028

import "strconv"

// 2020.06.18
// https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeInfo struct {
	Val   int
	Depth int
}

func recoverFromPreorder(S string) *TreeNode {
	if len(S) == 0 {
		return nil
	}
	nodeInfos := preProcess(S)
	return generateTree(nodeInfos)
}

func preProcess(S string) []*nodeInfo {
	var nodeInfos []*nodeInfo
	for len(S) > 0 {
		n := &nodeInfo{}
		idx := 0
		for S[idx] == '-' {
			idx++
		}
		n.Depth = idx
		i := idx
		for idx < len(S) && S[idx] >= '0' && S[idx] <= '9' {
			idx++
		}
		n.Val, _ = strconv.Atoi(S[i:idx])
		nodeInfos = append(nodeInfos, n)
		S = S[idx:]
	}
	return nodeInfos
}

func generateTree(nodeInfos []*nodeInfo) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	root := &TreeNode{Val: nodeInfos[0].Val}
	nodeInfos = nodeInfos[1:]
	nodeMap[0] = root
	for _, nInfo := range nodeInfos {
		parent := nodeMap[nInfo.Depth-1] // 默认必定存在
		n := &TreeNode{Val: nInfo.Val}
		if parent.Left == nil {
			parent.Left = n
		} else if parent.Right == nil {
			parent.Right = n
		}
		nodeMap[nInfo.Depth] = n
	}
	return root
}
