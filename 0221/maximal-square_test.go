package _0221

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{[][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}}, want: 4},
		{name: "test2", args: args{[][]byte{
			{'1', '0', '1', '1', '1'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '0'},
		}}, want: 9},
		{name: "test3", args: args{[][]byte{
			{'1', '0', '1', '1', '1'},
		}}, want: 1},
		{name: "test4", args: args{[][]byte{
			{'1'},
			{'0'},
			{'1'},
		}}, want: 1},
		{name: "test5", args: args{[][]byte{
			{'0'},
			{'0'},
			{'0'},
		}}, want: 0},
		{name: "test6", args: args{[][]byte{
			{'0', '0', '0', '0', '0'},
		}}, want: 0},
		{name: "test7", args: args{[][]byte{
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '0'},
		}}, want: 16},
		{name: "test8", args: args{[][]byte{
			{'1', '0', '1', '1', '0'},
			{'1', '1', '0', '1', '1'},
			{'1', '0', '1', '0', '1'},
			{'1', '1', '1', '1', '0'},
		}}, want: 1},
		{name: "test9", args: args{[][]byte{
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
		}}, want: 25},
		{name: "test_1", args: args{[][]byte{
			{'0'},
		}}, want: 0},
		{name: "test_2", args: args{[][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
