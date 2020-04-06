package _31

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{name: "test1", args: args{nums: []int{1, 1, 5}}, expected: []int{1, 5, 1}},
		{name: "test2", args: args{nums: []int{1, 2, 3}}, expected: []int{1, 3, 2}},
		{name: "test3", args: args{nums: []int{3, 2, 1}}, expected: []int{1, 2, 3}},
		{name: "test4", args: args{nums: []int{1, 2, 3, 4}}, expected: []int{1, 2, 4, 3}},
		{name: "test5", args: args{nums: []int{1, 2, 4, 3}}, expected: []int{1, 3, 2, 4}},
		{name: "test6", args: args{nums: []int{1, 4, 3, 2}}, expected: []int{2, 1, 3, 4}},
		{name: "test7", args: args{nums: []int{1, 2}}, expected: []int{2, 1}},
		{name: "test8", args: args{nums: []int{2, 1}}, expected: []int{1, 2}},
		{name: "test9", args: args{nums: []int{2, 3, 4, 1}}, expected: []int{2, 4, 1, 3}},
		{name: "test10", args: args{nums: []int{2, 3, 4, 1}}, expected: []int{2, 4, 1, 3}},
		{name: "test_1", args: args{nums: []int{5, 1, 1}}, expected: []int{1, 1, 5}},
		{name: "test_2", args: args{nums: []int{2, 3, 1, 3, 3}}, expected: []int{2, 3, 3, 1, 3}},
		{name: "test_3", args: args{nums: []int{2, 3, 3, 3, 1}}, expected: []int{3, 1, 2, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, tt.args.nums)
			}
		})
	}
}
