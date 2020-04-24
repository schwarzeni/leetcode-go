package _0024

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{generateLinkList([]int{1, 2, 3, 4, 5})},
			want: generateLinkList([]int{2, 1, 4, 3, 5}),
		},
		{
			name: "test2",
			args: args{generateLinkList([]int{1})},
			want: generateLinkList([]int{1}),
		},
		{
			name: "test3",
			args: args{generateLinkList([]int{1, 2})},
			want: generateLinkList([]int{2, 1}),
		},
		{
			name: "test4",
			args: args{generateLinkList([]int{1, 2, 3, 4})},
			want: generateLinkList([]int{2, 1, 4, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
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
