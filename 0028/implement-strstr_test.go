package _0028

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{
			haystack: "hello",
			needle:   "ll",
		}, want: 2},
		{name: "test2", args: args{
			haystack: "aaaaa",
			needle:   "bba",
		}, want: -1},
		{name: "test3", args: args{
			haystack: "aaaaa",
			needle:   "",
		}, want: 0},
		{name: "test4", args: args{
			haystack: "aaaaaaaaaaabc",
			needle:   "aaaaabc",
		}, want: 6},
		{name: "test_1", args: args{
			haystack: "",
			needle:   "",
		}, want: 0},
		{name: "test_2", args: args{
			haystack: "baabbbbababbbabab",
			needle:   "abbab",
		}, want: -1},
		{name: "test_3", args: args{
			haystack: "abbabaaaabbbaabaabaabbbaaabaaaaaabbbabbaabbabaabbabaaaaababbabbaaaaabbbbaaabbaaabbbbabbbbaaabbaaaaababbaababbabaaabaabbbbbbbaabaabaabbbbababbbababbaaababbbabaabbaaabbbba",
			needle:   "bbbbbbaa",
		}, want: 118},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
