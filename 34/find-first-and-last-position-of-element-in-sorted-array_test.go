package _34

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
		}, want: []int{3, 4}},
		{name: "test2", args: args{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
		}, want: []int{-1, -1}},
		{name: "test3", args: args{
			nums:   []int{1, 2, 3, 4, 5, 5, 5, 5, 5, 6, 7, 8, 9, 10, 11},
			target: 5,
		}, want: []int{4, 8}},
		{name: "test4", args: args{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 10, 11},
			target: 10,
		}, want: []int{9, 13}},
		{name: "test5", args: args{
			nums:   []int{},
			target: 10,
		}, want: []int{-1, -1}},
		{name: "test6", args: args{
			nums:   []int{10},
			target: 10,
		}, want: []int{0, 0}},
		{name: "test7", args: args{
			nums:   []int{9, 10},
			target: 10,
		}, want: []int{1, 1}},
		{name: "test8", args: args{
			nums:   []int{9, 10, 10},
			target: 10,
		}, want: []int{1, 2}},
		{name: "test9", args: args{
			nums:   []int{9, 9, 10},
			target: 9,
		}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
