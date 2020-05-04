package _0058

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{s: "  aa  bbbc  dd  "}, want: 2},
		{name: "test2", args: args{s: "   "}, want: 0},
		{name: "test3", args: args{s: ""}, want: 0},
		{name: "test4", args: args{s: "a b cc"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
