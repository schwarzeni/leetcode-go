package _65

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{s: " +23.54e-43  "}, want: true},
		{name: "test2", args: args{s: "    "}, want: false},
		{name: "test3", args: args{s: "+23.54e-4a3"}, want: false},
		{name: "test4", args: args{s: "+23.54e-4.3"}, want: false},
		{name: "test5", args: args{s: "+23.54e+43"}, want: true},
		{name: "test6", args: args{s: "-23.e+43"}, want: true},
		{name: "test7", args: args{s: "-.e+43"}, want: false},
		{name: "test8", args: args{s: "e3"}, want: false},
		{name: "test9", args: args{s: "-23.e+43"}, want: true},
		{name: "test10", args: args{s: " --6 "}, want: false},
		{name: "test11", args: args{s: "-+6 "}, want: false},
		{name: "test12", args: args{s: "53.5e93"}, want: true},
		{name: "test13", args: args{s: "95a54e53"}, want: false},
		{name: "test14", args: args{s: "0"}, want: true},
		{name: "test15", args: args{s: "0.1"}, want: true},
		{name: "test16", args: args{s: "0.1  a"}, want: false},
		{name: "test17", args: args{s: ".23"}, want: true},
		{name: "test18", args: args{s: "23e32   "}, want: true},
		{name: "test19", args: args{s: "23       "}, want: true},
		{name: "test20", args: args{s: "23e+2       "}, want: true},
		{name: "test21", args: args{s: "23e-23       "}, want: true},
		{name: "test22", args: args{s: "2.33a"}, want: false},
		{name: "test23", args: args{s: "2.3e1+2"}, want: false},
		{name: "test23", args: args{s: "2.3e1-2"}, want: false},
		{name: "test24", args: args{s: " -. "}, want: false},
		{name: "test25", args: args{s: "+   "}, want: false},
		{name: "test26", args: args{s: "+e   "}, want: false},
		{name: "test27", args: args{s: " 2e   "}, want: false},
		{name: "test28", args: args{s: "2e"}, want: false},
		{name: "test29", args: args{s: "2.3   "}, want: true},
		{name: "test_1", args: args{s: "2."}, want: true},
		{name: "test_2", args: args{s: ".  "}, want: false},
		{name: "test_3", args: args{s: "+.  "}, want: false},
		{name: "test_4", args: args{s: ".9 "}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
