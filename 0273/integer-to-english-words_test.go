package _273

import "testing"

func Test_parseHundred(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{num: 1}, want: "One"},
		{name: "test2", args: args{num: 12}, want: "Twelve"},
		{name: "test3", args: args{num: 21}, want: "Twenty One"},
		{name: "test4", args: args{num: 40}, want: "Forty"},
		{name: "test5", args: args{num: 99}, want: "Ninety Nine"},
		{name: "test6", args: args{num: 100}, want: "One Hundred"},
		{name: "test7", args: args{num: 123}, want: "One Hundred Twenty Three"},
		{name: "test8", args: args{num: 120}, want: "One Hundred Twenty"},
		{name: "test9", args: args{num: 103}, want: "One Hundred Three"},
		{name: "test10", args: args{num: 280}, want: "Two Hundred Eighty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseHundred(tt.args.num); got != tt.want {
				t.Errorf("parseHundred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{num: 1234567}, want: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{name: "test2", args: args{num: 1234567891}, want: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"},
		{name: "test3", args: args{num: 1000000001}, want: "One Billion One"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
