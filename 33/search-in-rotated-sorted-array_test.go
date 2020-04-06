package _33

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0}, want: 4},
		{name: "test2", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3}, want: -1},
		{name: "test3", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 4}, want: 0},
		{name: "test4", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 5}, want: 1},
		{name: "test5", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 6}, want: 2},
		{name: "test6", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 7}, want: 3},
		{name: "test7", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 1}, want: 5},
		{name: "test8", args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 2}, want: 6},
		{name: "test9", args: args{nums: []int{2, 1}, target: 2}, want: 0},
		{name: "test10", args: args{nums: []int{2, 1}, target: 1}, want: 1},
		{name: "test_1", args: args{nums: []int{5, 1, 3}, target: 0}, want: -1},
		{name: "test_2", args: args{nums: []int{5, 1, 3}, target: 3}, want: 2},
		{name: "test_3", args: args{nums: []int{3, 4, 5, 6, 1, 2}, target: 2}, want: 5},
		{name: "test_4", args: args{nums: []int{6, 7, 1, 2, 3, 4, 5}, target: 6}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
