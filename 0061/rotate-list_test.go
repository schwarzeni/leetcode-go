package _0061

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
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
			head: generateLinkList([]int{1, 2, 3, 4, 5}),
			k:    2,
		}, want: generateLinkList([]int{4, 5, 1, 2, 3})},
		{name: "test2", args: args{
			head: generateLinkList([]int{1, 2, 3, 4, 5}),
			k:    7,
		}, want: generateLinkList([]int{4, 5, 1, 2, 3})},
		{name: "test3", args: args{
			head: generateLinkList([]int{0, 1, 2}),
			k:    4,
		}, want: generateLinkList([]int{2, 0, 1})},
		{name: "test4", args: args{
			head: generateLinkList([]int{}),
			k:    4,
		}, want: generateLinkList([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
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
