package _0521

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{a: "aa", b: "bb"}, want: 2},
		{name: "test2", args: args{a: "aa", b: "bbc"}, want: 3},
		{name: "test3", args: args{a: "aa", b: "b"}, want: 2},
		{name: "test4", args: args{a: "aa", b: "aa"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
