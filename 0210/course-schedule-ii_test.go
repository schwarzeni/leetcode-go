package _0210

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{
			numCourses: 4,
			prerequisites: [][]int{
				{1, 0},
				{2, 0},
				{3, 1},
				{3, 2},
			},
		}, want: []int{0, 1, 2, 3}},
		{name: "test2", args: args{
			numCourses: 5,
			prerequisites: [][]int{
				{1, 0},
				{2, 0},
				{3, 4},
				{3, 2},
			},
		}, want: []int{0, 4, 1, 2, 3}},
		{name: "test3", args: args{
			numCourses: 4,
			prerequisites: [][]int{
				{1, 0},
				{2, 0},
				{0, 1},
				{3, 2},
			},
		}, want: []int{}},
		{name: "test4", args: args{
			numCourses: 7,
			prerequisites: [][]int{
				{1, 0},
				{2, 0},
				{3, 4},
				{3, 2},
				{6, 5},
			},
		}, want: []int{0, 4, 5, 1, 2, 6, 3}},
		{name: "test5", args: args{
			numCourses: 7,
			prerequisites: [][]int{
				{1, 0},
				{2, 0},
			},
		}, want: []int{0, 3, 4, 5, 6, 1, 2}},
		{name: "test6", args: args{
			numCourses: 3,
			prerequisites: [][]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
		}, want: []int{2, 1, 0}},
		{name: "test_1", args: args{
			numCourses:    7,
			prerequisites: [][]int{},
		}, want: []int{0, 1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
