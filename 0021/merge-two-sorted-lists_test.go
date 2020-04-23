package _0021

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
			l1: generateLinkList([]int{1, 2, 4}),
			l2: generateLinkList([]int{1, 3, 4}),
		}, want: generateLinkList([]int{1, 1, 2, 3, 4, 4})},
		{name: "test2", args: args{
			l1: generateLinkList([]int{1, 2, 5, 8, 9}),
			l2: generateLinkList([]int{1, 3, 4}),
		}, want: generateLinkList([]int{1, 1, 2, 3, 4, 5, 8, 9})},
		{name: "test3", args: args{
			l1: generateLinkList([]int{1, 2, 5, 8, 9}),
			l2: generateLinkList([]int{}),
		}, want: generateLinkList([]int{1, 2, 5, 8, 9})},
		{name: "test4", args: args{
			l1: generateLinkList([]int{}),
			l2: generateLinkList([]int{}),
		}, want: generateLinkList([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
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
