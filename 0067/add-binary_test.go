package _0067

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{
			a: "11",
			b: "1",
		}, want: "100"},
		{name: "test2", args: args{
			a: "11",
			b: "11",
		}, want: "110"},
		{name: "test3", args: args{
			a: "1010",
			b: "1011",
		}, want: "10101"},
		{name: "test4", args: args{
			a: "10001",
			b: "1",
		}, want: "10010"},
		{name: "test_1", args: args{
			a: "1",
			b: "111",
		}, want: "1000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
