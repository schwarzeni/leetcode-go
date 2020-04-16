package _51

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "test1", args: args{n: 4}, want: [][]string{
			{".Q..", // 解法 1
				"...Q",
				"Q...",
				"..Q."},
			{"..Q.", // 解法 2
				"Q...",
				"...Q",
				".Q.."},
		}},
		{name: "test2", args: args{n: 5}, want: [][]string{
			{"Q....", "..Q..", "....Q", ".Q...", "...Q."},
			{"Q....", "...Q.", ".Q...", "....Q", "..Q.."},
			{".Q...", "...Q.", "Q....", "..Q..", "....Q"},
			{".Q...", "....Q", "..Q..", "Q....", "...Q."},
			{"..Q..", "Q....", "...Q.", ".Q...", "....Q"},
			{"..Q..", "....Q", ".Q...", "...Q.", "Q...."},
			{"...Q.", "Q....", "..Q..", "....Q", ".Q..."},
			{"...Q.", ".Q...", "....Q", "..Q..", "Q...."},
			{"....Q", ".Q...", "...Q.", "Q....", "..Q.."},
			{"....Q", "..Q..", "Q....", "...Q.", ".Q..."},
		}},
		{name: "test3", args: args{n: 0}, want: [][]string{{}}},
		{name: "test4", args: args{n: 1}, want: [][]string{{"Q"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

//   0 1 2 3 x
// 0 . 1 . .
// 1 . . . 1
// 2 . . 1 .
// 3 1 . . .
// y
