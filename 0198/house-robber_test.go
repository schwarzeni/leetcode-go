package _198

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[]int{1, 2}}, want: 2},
		{name: "test2", args: args{[]int{2}}, want: 2},
		{name: "test3", args: args{[]int{}}, want: 0},
		{name: "test4", args: args{[]int{1, 2, 3, 1}}, want: 4},
		{name: "test5", args: args{[]int{2, 7, 9, 3, 1}}, want: 12},
		{name: "test_1", args: args{[]int{2, 1, 1, 2}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
