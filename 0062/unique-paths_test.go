package _0062

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{m: 7, n: 3}, want: 28},
		{name: "test2", args: args{m: 3, n: 2}, want: 3},
		{name: "test_1", args: args{m: 23, n: 12}, want: 193536720},
		{name: "test_2", args: args{m: 1, n: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
