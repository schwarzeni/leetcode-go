package _0041

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[]int{3, 4, -1, 1}}, want: 2},
		{name: "test2", args: args{[]int{2, 6, 0, 8, 4, 7, -1, 3, 1}}, want: 5},
		{name: "test3", args: args{[]int{-1, 0, -3}}, want: 1},
		{name: "test4", args: args{[]int{4, 5, 6, 9}}, want: 1},
		{name: "test5", args: args{[]int{1, 2, 3, 4, 5, 6}}, want: 7},
		{name: "test6", args: args{[]int{}}, want: 1},
		{name: "test_1", args: args{[]int{0, 1}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
