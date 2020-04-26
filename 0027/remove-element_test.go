package _0027

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
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
		{name: "test1", args: args{
			nums: []int{3, 2, 2, 3},
			val:  3,
		}, want: want{
			l:    2,
			nums: []int{2, 2},
		}},
		{name: "test2", args: args{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
		}, want: want{
			l:    5,
			nums: []int{0, 1, 3, 0, 4},
		}},
		{name: "test3", args: args{
			nums: []int{},
			val:  2,
		}, want: want{
			l:    0,
			nums: []int{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want.l {
				t.Errorf("removeElement() = %v, want %v", got, tt.want.l)
			}
			w := tt.args.nums[:tt.want.l]
			if !reflect.DeepEqual(w, tt.want.nums) {
				t.Errorf("removeDuplicates() = %v, want %v", w, tt.want.nums)
			}
		})
	}
}
