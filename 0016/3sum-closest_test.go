package _0016

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{
			nums:   []int{-1, 2, 1, -4},
			target: 1,
		}, want: 2},
		{name: "test2", args: args{
			nums:   []int{1, 1, 1, 0},
			target: 100,
		}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
