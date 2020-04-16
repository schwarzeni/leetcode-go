package _89

import (
	"reflect"
	"testing"
)

func Test_grayCode(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{n: 2}, want: []int{0, 2, 3, 1}},
		{name: "test2", args: args{n: 0}, want: []int{0}},
		//{name: "test3", args: args{n: 5}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grayCode(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toInt(t *testing.T) {
	n := 3
	dataSet := []byte{'1', '1', '0'}
	toInt := func() (result int) {
		for idx, _ := range dataSet {
			if dataSet[n-1-idx] == '1' {
				result += 1 << idx
			}
		}
		return
	}
	if toInt() != 6 {
		t.Errorf("expect 6, but got %v", toInt())
	}
}
