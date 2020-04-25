package _0026

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	type want struct {
		l    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{name: "test1", args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, want: want{
			l:    5,
			nums: []int{0, 1, 2, 3, 4},
		}},
		{name: "test2", args: args{nums: []int{0, 0, 1}}, want: want{
			l:    2,
			nums: []int{0, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want.l {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want.l)
			}
			w := tt.args.nums[:tt.want.l]
			if !reflect.DeepEqual(w, tt.want.nums) {
				t.Errorf("removeDuplicates() = %v, want %v", w, tt.want.nums)
			}
		})
	}
}
