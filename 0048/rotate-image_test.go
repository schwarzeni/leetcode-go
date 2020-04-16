package _48

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test1", args: args{matrix: [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}}, want: [][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		}},
		{name: "test2", args: args{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}}, want: [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}},
		{name: "test3", args: args{matrix: [][]int{
			{2, 3},
			{5, 6},
		}}, want: [][]int{
			{5, 2},
			{6, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("expect \n%s, but got \n%s", printArray(tt.want), printArray(tt.args.matrix))
			}
		})
	}
}

func printArray(a [][]int) string {
	result := ""
	for _, cs := range a {
		for _, v := range cs {
			result += fmt.Sprintf("%d ", v)
			if v <= 9 {
				result += " "
			}
		}
		result += "\n"
	}
	return result
}
