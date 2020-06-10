package _0081

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 0}, want: true},
		{name: "test1", args: args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 2}, want: true},
		{name: "test1", args: args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 5}, want: true},
		{name: "test1", args: args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 6}, want: true},
		{name: "test1", args: args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 1}, want: true},
		{name: "test1", args: args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 4}, want: false},
		{name: "test1", args: args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 3}, want: false},
		{name: "test2", args: args{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 3}, want: false},
		{name: "test_1", args: args{nums: []int{1}, target: 0}, want: false},
		{name: "test_2", args: args{nums: []int{3, 1}, target: 1}, want: true},
		{name: "test_3", args: args{nums: []int{1, 1, 3, 1}, target: 3}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
