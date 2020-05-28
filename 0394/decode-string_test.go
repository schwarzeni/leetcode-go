package _0394

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{s: "3[a2[c]]"}, want: "accaccacc"},
		{name: "test2", args: args{s: "a3[2[c]a4[dfb]6[ed]12[a]a3[2[d]a[vb]]a]av"}, want: "accadfbdfbdfbdfbededededededaaaaaaaaaaaaaddaddaddaaccadfbdfbdfbdfbededededededaaaaaaaaaaaaaddaddaddaaccadfbdfbdfbdfbededededededaaaaaaaaaaaaaddaddaddaaav"},
		{name: "test3", args: args{s: "abcde"}, want: "abcde"},
		{name: "test4", args: args{s: ""}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
