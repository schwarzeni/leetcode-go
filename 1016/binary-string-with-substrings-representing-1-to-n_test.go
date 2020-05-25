package _1016

import (
	"strconv"
	"testing"
)

func Test_queryString(t *testing.T) {
	type args struct {
		S string
		N int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{
			S: "0110",
			N: 3,
		}, want: true},
		{name: "test2", args: args{
			S: "0110",
			N: 4,
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryString(tt.args.S, tt.args.N); got != tt.want {
				t.Errorf("queryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toBinaryString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{n: 234532}, want: strconv.FormatInt(int64(234532), 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toBinaryString(tt.args.n); got != tt.want {
				t.Errorf("toBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
