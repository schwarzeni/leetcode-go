package _0009

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{x: 121}, want: true},
		{name: "test2", args: args{x: 1210121}, want: true},
		{name: "test3", args: args{x: -12}, want: false},
		{name: "test4", args: args{x: 12345}, want: false},
		{name: "test5", args: args{x: 0}, want: true},
		{name: "test6", args: args{x: 1}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
