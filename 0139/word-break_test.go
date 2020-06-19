package _0139

import "testing"

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{
			s:        "goalspecial",
			wordDict: []string{"go", "goal", "goals", "special"},
		}, want: true},
		{name: "test2", args: args{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
		}, want: false},
		{name: "test3", args: args{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
		}, want: true},
		{name: "test4", args: args{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
