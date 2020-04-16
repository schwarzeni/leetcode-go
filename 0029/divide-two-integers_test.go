package _29

import (
	"math"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{
			dividend: 12,
			divisor:  5,
		}, want: 2},
		{name: "test2", args: args{
			dividend: -12,
			divisor:  5,
		}, want: -2},
		{name: "test3", args: args{
			dividend: -12,
			divisor:  -5,
		}, want: 2},
		{name: "test4", args: args{
			dividend: math.MinInt32,
			divisor:  -1,
		}, want: math.MaxInt32},
		{name: "test5", args: args{
			dividend: 0,
			divisor:  -1,
		}, want: 0},
		{name: "test6", args: args{
			dividend: math.MaxInt32,
			divisor:  2,
		}, want: math.MaxInt32 / 2},
		{name: "test7", args: args{
			dividend: math.MaxInt32,
			divisor:  math.MinInt32,
		}, want: 0},
		{name: "test8", args: args{
			dividend: math.MinInt32,
			divisor:  math.MinInt32,
		}, want: 1},
		{name: "test_1_超出时间限制", args: args{
			dividend: -2147483648,
			divisor:  -2,
		}, want: 1073741824},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
