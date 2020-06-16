package _0297

import (
	"strconv"
	"strings"
)

// 2020.06.16
// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var val []string
	var queue []*TreeNode

	if root == nil {
		return "[]"
	}
	val = append(val, strconv.Itoa(root.Val))
	queue = append(queue, root)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n.Left != nil {
			val = append(val, strconv.Itoa(n.Left.Val))
			queue = append(queue, n.Left)
		} else {
			val = append(val, "null")
		}
		if n.Right != nil {
			val = append(val, strconv.Itoa(n.Right.Val))
			queue = append(queue, n.Right)
		} else {
			val = append(val, "null")
		}
	}
	for val[len(val)-1] == "null" {
		val = val[:len(val)-1]
	}
	return "[" + strings.Join(val, ",") + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 2 {
		return nil
	}
	data = data[1 : len(data)-1]
	vals := strings.Split(data, ",")
	var roots []*TreeNode

	v, _ := strconv.Atoi(vals[0])
	vals = vals[1:]
	root := &TreeNode{Val: v}
	roots = append(roots, root)
	for {
		if len(roots) == 0 {
			break
		}
		r := roots[0]
		roots = roots[1:]
		if len(vals) == 0 {
			break
		}
		if vals[0] != "null" {
			v, _ := strconv.Atoi(vals[0])
			lr := &TreeNode{Val: v}
			r.Left = lr
			roots = append(roots, lr)
		}
		vals = vals[1:]

		if len(vals) == 0 {
			break
		}
		if vals[0] != "null" {
			v, _ := strconv.Atoi(vals[0])
			rr := &TreeNode{Val: v}
			r.Right = rr
			roots = append(roots, rr)
		}
		vals = vals[1:]
	}

	return root
}
