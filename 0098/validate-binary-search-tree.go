package _0098

// 2020.05.05
// https://leetcode-cn.com/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	var (
		iter func(root *TreeNode) bool
		pre  int
		flag = -1
	)
	if root == nil {
		return true
	}
	iter = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if r := iter(root.Left); !r {
			return false
		}
		if flag == -1 {
			flag = 1
		} else if pre >= root.Val {
			return false
		}
		pre = root.Val
		if r := iter(root.Right); !r {
			return false
		}
		return true
	}

	return iter(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
