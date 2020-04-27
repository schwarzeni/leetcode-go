package _0032

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{s: "(()"}, want: 2},
		{name: "test2", args: args{s: ")()())"}, want: 4},
		{name: "test3", args: args{s: "(()))(())(())(("}, want: 8},
		{name: "test4", args: args{s: "())"}, want: 2},
		{name: "test5", args: args{s: "))"}, want: 0},
		{name: "test_1", args: args{s: "()(()"}, want: 2},
		{name: "test_2", args: args{s: "(()()"}, want: 4},
		{name: "test_太慢", args: args{s: func() string {
			s := ""
			for i := 0; i < 100000; i++ {
				s += "("
			}
			return s + ")"
		}()}, want: 2},
		{name: "test7", args: args{s: "())((())"}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
