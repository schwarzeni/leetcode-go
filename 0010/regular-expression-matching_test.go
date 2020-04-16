package _10

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{s: "aa", p: "a"}, want: false},
		{name: "test2", args: args{s: "aaaaa", p: "a*"}, want: true},
		{name: "test3", args: args{s: "abab", p: "a.*b"}, want: true},
		{name: "test4", args: args{s: "aab", p: "c*a*b"}, want: true},
		{name: "test5", args: args{s: "mississippi", p: "mis*is*p*."}, want: false},
		{name: "test6", args: args{s: "abababababababa", p: "a.*b.*a.*.*.*.*a*b"}, want: false},
		{name: "test7", args: args{s: "", p: "."}, want: false},
		{name: "test8", args: args{s: "", p: ".*"}, want: true},
		{name: "test9", args: args{s: "bv", p: "c*a*b*...*b*"}, want: true},
		{name: "test10", args: args{s: "bv", p: "c*a*b*..*.*vb*"}, want: true},
		{name: "test11", args: args{s: "bv", p: "c*a*b*..*d.*ab*"}, want: false},
		{name: "test12", args: args{s: "aabbcc", p: "a*a.*b*a.*b*.*c*d*.*"}, want: true},
		{name: "test_1", args: args{s: "a", p: "ab*"}, want: true},
		{name: "test_2", args: args{s: "abcd", p: "d*"}, want: false},
		{name: "test_3", args: args{s: "acaabbaccbbacaabbbb", p: "a*.*b*.*a*aa*a*"}, want: false},
		{name: "test_4", args: args{s: "abcaaaaaaabaabcabac", p: ".*ab.a.*a*a*.*b*b*"}, want: true},
		{name: "test_5", args: args{s: "bcccccbaccccacaa", p: ".*bb*c*a*b*.*b*b*c*"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
