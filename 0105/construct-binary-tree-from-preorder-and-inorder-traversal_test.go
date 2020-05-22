package _0105

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{name: "test1", args: args{
			preorder: []int{3, 9, 17, 16, 20, 15, 7},
			inorder:  []int{17, 9, 16, 3, 15, 20, 7},
		}, want: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 17}, Right: &TreeNode{Val: 16}},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		}},
		{name: "test2", args: args{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
		}, want: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		}},
		{name: "test_1", args: args{
			preorder: []int{1, 2},
			inorder:  []int{1, 2},
		}, want: &TreeNode{Val: 1,
			Right: &TreeNode{Val: 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
