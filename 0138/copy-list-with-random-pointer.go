package _0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 2020.06.12
// https://leetcode-cn.com/problems/copy-list-with-random-pointer/
func copyRandomList(head *Node) *Node {
	var newNodes []*Node
	var record = make(map[*Node]*Node)
	if head == nil {
		return nil
	}
	p := head
	for p != nil {
		i, n := len(newNodes), &Node{Val: p.Val, Random: p.Random}
		record[p] = n
		newNodes = append(newNodes, n)
		if i > 0 {
			newNodes[i-1].Next = n
		}
		p = p.Next
	}
	for i, _ := range newNodes {
		if rn := newNodes[i].Random; rn != nil {
			newNodes[i].Random = record[rn]
		}
	}
	return newNodes[0]
}
