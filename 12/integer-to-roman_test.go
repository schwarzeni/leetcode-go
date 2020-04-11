package _12

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{num: 2}, want: "II"},
		{name: "test2", args: args{num: 12}, want: "XII"},
		{name: "test3", args: args{num: 27}, want: "XXVII"},
		{name: "test4", args: args{num: 1994}, want: "MCMXCIV"},
		{name: "test5", args: args{num: 9}, want: "IX"},
		{name: "test6", args: args{num: 4}, want: "IV"},
		{name: "test7", args: args{num: 3}, want: "III"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
