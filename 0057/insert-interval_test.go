package _0057

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test1", args: args{
			intervals:   [][]int{{5, 7}, {8, 10}, {15, 18}},
			newInterval: []int{1, 2},
		}, want: [][]int{{1, 2}, {5, 7}, {8, 10}, {15, 18}}},
		{name: "test2", args: args{
			intervals:   [][]int{{5, 7}, {8, 10}, {15, 18}},
			newInterval: []int{20, 22},
		}, want: [][]int{{5, 7}, {8, 10}, {15, 18}, {20, 22}}},
		{name: "test3", args: args{
			intervals:   [][]int{{5, 7}, {8, 10}, {15, 18}},
			newInterval: []int{11, 13},
		}, want: [][]int{{5, 7}, {8, 10}, {11, 13}, {15, 18}}},
		{name: "test4", args: args{
			intervals:   [][]int{{5, 7}, {8, 10}, {15, 18}},
			newInterval: []int{9, 18},
		}, want: [][]int{{5, 7}, {8, 18}}},
		{name: "test5", args: args{
			intervals:   [][]int{{5, 7}, {8, 10}, {15, 18}},
			newInterval: []int{4, 13},
		}, want: [][]int{{4, 13}, {15, 18}}},
		{name: "test6", args: args{
			intervals:   [][]int{{5, 6}, {8, 10}, {15, 18}},
			newInterval: []int{7, 11},
		}, want: [][]int{{5, 6}, {7, 11}, {15, 18}}},
		{name: "test7", args: args{
			intervals:   [][]int{{5, 6}, {8, 10}, {15, 18}},
			newInterval: []int{8, 10},
		}, want: [][]int{{5, 6}, {8, 10}, {15, 18}}},
		{name: "test8", args: args{
			intervals:   [][]int{},
			newInterval: []int{8, 10},
		}, want: [][]int{{8, 10}}},
		{name: "test9", args: args{
			intervals:   [][]int{{9, 10}},
			newInterval: []int{8, 11},
		}, want: [][]int{{8, 11}}},
		{name: "test10", args: args{
			intervals:   [][]int{{12, 13}},
			newInterval: []int{8, 11},
		}, want: [][]int{{8, 11}, {12, 13}}},
		{name: "test11", args: args{
			intervals:   [][]int{{8, 11}},
			newInterval: []int{12, 13},
		}, want: [][]int{{8, 11}, {12, 13}}},
		{name: "test12", args: args{
			intervals:   [][]int{{5, 6}, {8, 10}, {15, 18}},
			newInterval: []int{6, 14},
		}, want: [][]int{{5, 14}, {15, 18}}},
		{name: "test13", args: args{
			intervals:   [][]int{{5, 6}, {8, 10}, {15, 18}},
			newInterval: []int{6, 15},
		}, want: [][]int{{5, 18}}},
		{name: "test_1", args: args{
			intervals:   [][]int{{5, 6}, {8, 10}, {15, 18}},
			newInterval: []int{9, 15},
		}, want: [][]int{{5, 6}, {8, 18}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
