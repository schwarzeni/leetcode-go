package _1371

import "testing"

func Test_findTheLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{s: "eleetminicoworoep"}, want: 13},
		{name: "test2", args: args{s: "leetcodeisgreat"}, want: 5},
		{name: "test3", args: args{s: "bcbcbc"}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("findTheLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
