package _0007

import (
	"math"
	"testing"
)

func Test_isPlusOverflow(t *testing.T) {
	type args struct {
		raw   int
		plus  int
		bound int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{
			raw:   -11,
			plus:  10,
			bound: -100,
		}, want: true},
		{name: "test2", args: args{
			raw:   -9,
			plus:  10,
			bound: -100,
		}, want: false},
		{name: "test3", args: args{
			raw:   11,
			plus:  10,
			bound: 100,
		}, want: true},
		{name: "test4", args: args{
			raw:   9,
			plus:  10,
			bound: 100,
		}, want: false},
		{name: "test5", args: args{
			raw:   10,
			plus:  10,
			bound: 100,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPlusOverflow(tt.args.raw, tt.args.plus, tt.args.bound); got != tt.want {
				t.Errorf("isPlusOverflow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAddOverflow(t *testing.T) {
	type args struct {
		raw   int
		incre int
		bound int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{
			raw:   10,
			incre: 9,
			bound: 20,
		}, want: false},
		{name: "test2", args: args{
			raw:   10,
			incre: 11,
			bound: 20,
		}, want: true},
		{name: "test3", args: args{
			raw:   -10,
			incre: -9,
			bound: -20,
		}, want: false},
		{name: "test4", args: args{
			raw:   -10,
			incre: -11,
			bound: -20,
		}, want: true},
		{name: "test5", args: args{
			raw:   -10,
			incre: -10,
			bound: -20,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAddOverflow(tt.args.raw, tt.args.incre, tt.args.bound); got != tt.want {
				t.Errorf("isAddOverflow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{x: 123}, want: 321},
		{name: "test2", args: args{x: -123}, want: -321},
		{name: "test3", args: args{x: math.MaxInt32}, want: 0},
		{name: "test4", args: args{x: 1147483647}, want: 0},
		{name: "test5", args: args{x: 0}, want: 0},
		{name: "test6", args: args{x: -123456}, want: -654321},
		{name: "test7", args: args{x: -2147483648}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
