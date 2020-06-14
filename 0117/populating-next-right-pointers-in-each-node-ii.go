package _0117

// 2020.06.14
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 非递归形式
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var curr, lstart *Node
	lstart = root
	for lstart != nil {
		curr = lstart
		var pchild *Node
		for curr != nil {
			if pchild != nil {
				if curr.Left != nil && curr.Right != nil {
					pchild.Next = curr.Left
					curr.Left.Next = curr.Right
					pchild = curr.Right
				} else if curr.Left != nil && curr.Right == nil {
					pchild.Next = curr.Left
					pchild = curr.Left
				} else if curr.Left == nil && curr.Right != nil {
					pchild.Next = curr.Right
					pchild = curr.Right
				}
			} else {
				if curr.Left != nil {
					curr.Left.Next = curr.Right
					pchild = curr.Left
				}
				if curr.Right != nil {
					pchild = curr.Right
				}
				if curr.Left != nil {
					lstart = curr.Left
				} else {
					lstart = curr.Right
				}
			}
			curr = curr.Next
		}
	}
	return root
}

// 递归形式
//func connect(root *Node) *Node {
//	if root == nil {
//		return nil
//	}
//	if root.Left != nil {
//		if root.Right != nil {
//			root.Left.Next = root.Right
//		} else {
//			root.Left.Next = nextNode(root.Next)
//		}
//	}
//	if root.Right != nil {
//		root.Right.Next = nextNode(root.Next)
//	}
//	connect(root.Right)
//	connect(root.Left)
//
//	return root
//}
//
//func nextNode(node *Node)*Node {
//	for node != nil {
//		if node.Left != nil {
//			return node.Left
//		}
//		if node.Right != nil {
//			return node.Right
//		}
//		node = node.Next
//	}
//	return nil
//}
