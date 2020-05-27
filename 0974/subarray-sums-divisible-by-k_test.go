package _0974

import "testing"

func Test_subarraysDivByK(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{A: []int{4, 5, 0, -2, -3, 1}, K: 5}, want: 7},
		{name: "test2", args: args{A: []int{2, 3, -1, 1, 5}, K: 5}, want: 6},
		{name: "test3", args: args{A: []int{4, 7, 8, 6, 3, 2, 4, 6, 8, 7, 6, 3, 2, 4, 6, 7, 2, 7, 5, 4, 25, 6, 2}, K: 2}, want: 141},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysDivByK(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
