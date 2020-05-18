package _0152

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[]int{2, 3, -2, 4}}, want: 6},
		{name: "test2", args: args{[]int{-2, 0, -1}}, want: 0},
		{name: "test3", args: args{[]int{2, 3, -2, 4, -5}}, want: 240},
		{name: "test4", args: args{[]int{2, 3, -3, 0, 4, -5, -2}}, want: 40},
		{name: "test5", args: args{[]int{-2}}, want: -2},
		{name: "test6", args: args{[]int{-2, -3, -4}}, want: 12},
		{name: "test7", args: args{[]int{0, 0, 0, -2}}, want: 0},
		{name: "test8", args: args{[]int{-2, -2}}, want: 4},
		{name: "test9", args: args{[]int{0, -1}}, want: 0},
		{name: "test_1", args: args{[]int{-1, -1}}, want: 1},
		{name: "test_2", args: args{[]int{2, -5, -2, -4, 3}}, want: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
