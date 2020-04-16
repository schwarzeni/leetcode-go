package _54

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}}, want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{name: "test2", args: args{[][]int{
			{1, 2, 3},
			{5, 6, 7},
			{9, 10, 11},
			{20, 21, 22},
		}}, want: []int{1, 2, 3, 7, 11, 22, 21, 20, 9, 5, 6, 10}},
		{name: "test2", args: args{[][]int{
			{1},
			{5},
			{9},
			{20},
		}}, want: []int{1, 5, 9, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
