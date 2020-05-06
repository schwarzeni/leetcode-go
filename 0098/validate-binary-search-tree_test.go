package _0098

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{
			&TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				}}}, want: false},
		{name: "test2", args: args{
			&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}}, want: true},
		{name: "test_1", args: args{
			&TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 1},
			}}, want: false},
		{name: "test_2", args: args{
			&TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 1},
			}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
