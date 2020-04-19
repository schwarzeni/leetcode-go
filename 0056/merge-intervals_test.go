package _0056

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test1", args: args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{name: "test2", args: args{[][]int{{1, 4}, {4, 5}}}, want: [][]int{{1, 5}}},
		{name: "test3", args: args{[][]int{{1, 3}, {15, 20}, {5, 10}, {15, 18}, {5, 6}}}, want: [][]int{{1, 3}, {5, 10}, {15, 20}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
