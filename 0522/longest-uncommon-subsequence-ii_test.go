package _0522

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[]string{"ab", "abc", "abcd", "abb"}}, want: 4},
		{name: "test2", args: args{[]string{"abe", "abc", "abc", "abb"}}, want: 3},
		{name: "test3", args: args{[]string{"abe", "abe", "abe", "abe"}}, want: -1},
		{name: "test_1", args: args{[]string{"aaaa", "aaaa", "aa"}}, want: -1},
		{name: "test_2", args: args{[]string{"aabbcc", "aabbcc", "cb", "abc"}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.strs); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
