package _0019

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{
			head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}},
			n:    1,
		}, want: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}},
		{name: "test2", args: args{
			head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}},
			n:    2,
		}, want: &ListNode{Val: 0, Next: &ListNode{Val: 2, Next: nil}}},
		{name: "test3", args: args{
			head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}},
			n:    3,
		}, want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}},
		{name: "test4", args: args{
			head: &ListNode{Val: 0, Next: nil},
			n:    1,
		}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
