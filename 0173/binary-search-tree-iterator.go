package _0173

// 2020.09.05
// https://leetcode-cn.com/problems/binary-search-tree-iterator/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{}
	iter.initWithRoot(root)
	return iter
}

func (this *BSTIterator) initWithRoot(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	curr := this.stack[len(this.stack)-1]
	res := curr.Val
	this.stack = this.stack[:len(this.stack)-1]
	this.initWithRoot(curr.Right)
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
