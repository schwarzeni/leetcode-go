package _0003

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{s: "abcabcbb"}, want: 3},
		{name: "test2", args: args{s: "bbbbbb"}, want: 1},
		{name: "test3", args: args{s: "pwwkew"}, want: 3},
		{name: "test4", args: args{s: "abcde"}, want: 5},
		{name: "test5", args: args{s: "t"}, want: 1},
		{name: "test6", args: args{s: "tammzuxta"}, want: 6},
		{name: "test_1", args: args{s: "tmmzuxt"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
