package _0572

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		s *TreeNode
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{
			s: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}},
			t: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		}, want: true},
		{name: "test2", args: args{
			s: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}}, Right: &TreeNode{Val: 5}},
			t: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		}, want: false},
		{name: "test_1", args: args{
			s: &TreeNode{Val: 1},
			t: &TreeNode{Val: 0},
		}, want: false},
		{name: "test_2", args: args{
			s: &TreeNode{Val: 1, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}}}},
			t: &TreeNode{Val: 1, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}},
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
