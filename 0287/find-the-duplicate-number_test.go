package _0287

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[]int{1, 3, 4, 2, 2}}, want: 2},
		{name: "test2", args: args{[]int{3, 1, 3, 4, 2}}, want: 3},
		{name: "test3", args: args{[]int{75, 75, 75, 34, 32}}, want: 75},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
