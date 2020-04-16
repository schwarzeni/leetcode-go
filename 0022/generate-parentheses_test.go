package _22

import (
	"reflect"
	"sort"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test1", args: args{n: 3}, want: []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()"}},
		{name: "test2", args: args{n: 1}, want: []string{
			"()",
		}},
		{name: "test3", args: args{n: 2}, want: []string{
			"()()",
			"(())",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.args.n)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
