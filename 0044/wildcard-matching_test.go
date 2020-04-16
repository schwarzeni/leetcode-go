package _44

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
		{name: "test1", args: args{s: "", p: ""}, want: true},
		{name: "test2", args: args{s: "aa", p: "a"}, want: false},
		{name: "test3", args: args{s: "aa", p: "*"}, want: true},
		{name: "test4", args: args{s: "cb", p: "*a"}, want: false},
		{name: "test5", args: args{s: "cb", p: "?a"}, want: false},
		{name: "test6", args: args{s: "abcdef", p: "a*"}, want: true},
		{name: "test7", args: args{s: "abcdef", p: "a*"}, want: true},
		{name: "test8", args: args{s: "aacaef", p: "a*aef"}, want: true},
		{name: "test9", args: args{s: "aacaef", p: "a*f"}, want: true},
		{name: "test10", args: args{s: "adceb", p: "*a*b"}, want: true},
		{name: "test11", args: args{s: "acdcb", p: "a*c?b"}, want: false},
		{name: "test12", args: args{s: "aacaef", p: "??????"}, want: true},
		{name: "test13", args: args{s: "aacaef", p: "*"}, want: true},
		{name: "test14", args: args{s: "", p: "*"}, want: true},
		{name: "test15", args: args{s: "", p: "?"}, want: false},
		{name: "test_1", args: args{s: "ho", p: "ho**"}, want: true},
		{name: "test_2", args: args{s: "mississippi", p: "m??*ss*?i*pi"}, want: false},
		// 超出时间限制
		{name: "test_3", args: args{s: "aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", p: "a*******b"}, want: false},
		{name: "test_4", args: args{s: "abbbb", p: "?*b**"}, want: true},
		{name: "test_5", args: args{s: "aaba", p: "?***"}, want: true},
		// 超出时间限制
		{name: "test_6", args: args{
			s: "abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
			p: "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"},
			want: false},
	}
	for _, tt := range tests {
		//if tt.name == "test_6" {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
		//}
	}
}
