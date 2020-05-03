package _0050

import (
	"math"
	"testing"
)

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "test1", args: args{
			x: 2.00000,
			n: 3,
		}, want: math.Pow(2.0, 3)},
		{name: "test2", args: args{
			x: 2.00000,
			n: -2,
		}, want: math.Pow(2.0, -2)},
		{name: "test3", args: args{
			x: 2.10000,
			n: 3,
		}, want: math.Pow(2.1, 3)},
		{name: "test4", args: args{
			x: 0.00001,
			n: 2147483647,
		}, want: math.Pow(0.00001, 2147483647)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
