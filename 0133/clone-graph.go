package _0133

type Node struct {
	Val       int
	Neighbors []*Node
}

// 2020.06.12
// https://leetcode-cn.com/problems/clone-graph/
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	mapping := make(map[*Node]*Node)
	queue := []*Node{node}
	//mapping[node] = &Node{Val: node.Val, Neighbors: node.Neighbors}
	mapping[node] = &Node{Val: node.Val}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, nb := range n.Neighbors {
			if _, ok := mapping[nb]; !ok {
				//mapping[nb] = &Node{Val: nb.Val, Neighbors: nb.Neighbors}
				mapping[nb] = &Node{Val: nb.Val}
				queue = append(queue, nb)
			}
		}
	}

	for k, _ := range mapping {
		for _, nb := range k.Neighbors {
			mapping[k].Neighbors = append(mapping[k].Neighbors, mapping[nb])
		}
	}

	return mapping[node]
}
