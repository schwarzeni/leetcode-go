package _0680

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{s: "aba"}, want: true},
		{name: "test2", args: args{s: "abca"}, want: true},
		{name: "test3", args: args{s: "adbcba"}, want: true},
		{name: "test4", args: args{s: "abcbda"}, want: true},
		{name: "test5", args: args{s: "abcbdea"}, want: false},
		{name: "test6", args: args{s: "adebcba"}, want: false},
		{name: "test7", args: args{s: ""}, want: true},
		{name: "test8", args: args{s: "a"}, want: true},
		{name: "test9", args: args{s: "abcbae"}, want: true},
		{name: "test10", args: args{s: "eabcba"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
