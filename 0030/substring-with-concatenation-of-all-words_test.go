package _0030

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
		}, want: []int{0, 9}},
		{name: "test2", args: args{
			s:     "aaaaa",
			words: []string{"aa", "aa"},
		}, want: []int{0, 1}},
		{name: "test3", args: args{
			s:     "aaaaa",
			words: []string{"a", "a"},
		}, want: []int{0, 1, 2, 3}},
		{name: "test4", args: args{
			s:     "aaaaa",
			words: []string{""},
		}, want: []int{0, 1, 2, 3, 4, 5}},
		{name: "test_1", args: args{
			s:     "abababab",
			words: []string{"ab", "ba"},
		}, want: nil},
		{name: "test_2", args: args{
			s:     "barfoofoobarthefoobarman",
			words: []string{"bar", "foo", "the"},
		}, want: []int{6, 9, 12}},
	}
	for _, tt := range tests {
		//if tt.name == "test_1" {

		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
		//}
	}
}
