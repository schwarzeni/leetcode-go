package _0023

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "test1", args: args{[]*ListNode{
			generateLinkList([]int{1, 4, 5}),
			generateLinkList([]int{1, 3, 4}),
			generateLinkList([]int{2, 6}),
		}}, want: generateLinkList([]int{1, 1, 2, 3, 4, 4, 5, 6})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
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
