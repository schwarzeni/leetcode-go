package _0002

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{
			l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
		}, want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}},
		{name: "test2", args: args{
			l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}},
			l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
		}, want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: nil}}}},
		{name: "test3", args: args{
			l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}},
			l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}},
		}, want: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			if !equal(got, tt.want) {
				t.Errorf("Error!, got %v, but expected %v", listToArr(got), listToArr(tt.want))
			}
		})
	}
}

func equal(got *ListNode, want *ListNode) bool {
	for got != nil && want != nil {
		if got.Val != want.Val {
			return false
		}
		got = got.Next
		want = want.Next
	}
	return got == nil && want == nil
}

func listToArr(l *ListNode) (result []int) {
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return
}
