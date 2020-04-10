package _8

import (
	"math"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{str: "42"}, want: 42},
		{name: "test2", args: args{str: "     -42aabc"}, want: -42},
		{name: "test3", args: args{str: "4193 with words"}, want: 4193},
		{name: "test4", args: args{str: "words and 987"}, want: 0},
		{name: "test5", args: args{str: "-91283472332"}, want: -2147483648},
		{name: "test6", args: args{str: "1111111111111111111a v "}, want: math.MaxInt32},
		{name: "test7", args: args{str: "-a23"}, want: 0},
		{name: "test8", args: args{str: "+-23"}, want: 0},
		{name: "test9", args: args{str: "-2147483649"}, want: math.MinInt32},
		{name: "test10", args: args{str: "2147483648"}, want: math.MaxInt32},
		{name: "test8", args: args{str: "+a23"}, want: 0},
		{name: "test_1", args: args{str: "+23 "}, want: 23},
		{name: "test_2", args: args{str: "+-23"}, want: 0},
		{name: "test_3", args: args{str: "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"}, want: math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
