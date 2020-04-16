package _5

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{s: "babad"}, want: "bab"},
		{name: "test2", args: args{s: "cbbd"}, want: "bb"},
		{name: "test3", args: args{s: ""}, want: ""},
		{name: "test4", args: args{s: "a"}, want: "a"},
		{name: "test5", args: args{s: "gabcdefedcba"}, want: "abcdefedcba"},
		{name: "test6", args: args{s: "abcdefghi"}, want: "i"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
