package _0106

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		//{name: "test1", args: args{
		//    preorder: []int{3, 9, 17, 16, 20, 15, 7},
		//    inorder:  []int{17, 9, 16, 3, 15, 20, 7},
		//}, want: &TreeNode{Val: 3,
		//    Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 17}, Right: &TreeNode{Val: 16}},
		//    Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		//}},
		{name: "test2", args: args{
			postorder: []int{9, 15, 7, 20, 3},
			inorder:   []int{9, 3, 15, 20, 7},
		}, want: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		}},
		{name: "test3", args: args{
			postorder: []int{2, 1},
			inorder:   []int{1, 2},
		}, want: &TreeNode{Val: 1,
			Right: &TreeNode{Val: 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
