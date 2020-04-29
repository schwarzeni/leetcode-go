package _0038

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{n: 5}, want: "111221"},
		{name: "test2", args: args{n: 6}, want: "312211"},
		{name: "test3", args: args{n: 1}, want: "1"},
		{name: "test4", args: args{n: 2}, want: "11"},
		{name: "test5", args: args{n: 3}, want: "21"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
