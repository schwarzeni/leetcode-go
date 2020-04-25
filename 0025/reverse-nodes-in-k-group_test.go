package _0025

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{
			head: generateLinkList([]int{1, 2, 3, 4, 5, 6, 7}),
			k:    3,
		}, want: generateLinkList([]int{3, 2, 1, 6, 5, 4, 7})},
		{name: "test2", args: args{
			head: generateLinkList([]int{1, 2, 3, 4, 5, 6, 7}),
			k:    1,
		}, want: generateLinkList([]int{1, 2, 3, 4, 5, 6, 7})},
		{name: "test3", args: args{
			head: generateLinkList([]int{1, 2, 3, 4, 5, 6, 7}),
			k:    2,
		}, want: generateLinkList([]int{2, 1, 4, 3, 6, 5, 7})},
		{name: "test4", args: args{
			head: generateLinkList([]int{1, 2, 3, 4, 5, 6, 7}),
			k:    4,
		}, want: generateLinkList([]int{4, 3, 2, 1, 5, 6, 7})},
		{name: "test5", args: args{
			head: generateLinkList([]int{1, 2, 3, 4, 5, 6, 7}),
			k:    7,
		}, want: generateLinkList([]int{7, 6, 5, 4, 3, 2, 1})},
		{name: "test6", args: args{
			head: generateLinkList([]int{1, 2, 3, 4, 5, 6, 7}),
			k:    8,
		}, want: generateLinkList([]int{1, 2, 3, 4, 5, 6, 7})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateLinkList(args []int) *ListNode {
	header := &ListNode{Next: nil}
	h := header

	for _, v := range args {
		header.Next = &ListNode{Val: v}
		header = header.Next
	}

	return h.Next
}
