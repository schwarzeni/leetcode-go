package _0990

import "testing"

func Test_equationsPossible(t *testing.T) {
	type args struct {
		equations []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{[]string{"a==b", "b!=a"}}, want: false},
		{name: "test2", args: args{[]string{"b==a", "a==b"}}, want: true},
		{name: "test3", args: args{[]string{"a==b", "b==c", "a==c"}}, want: true},
		{name: "test4", args: args{[]string{"a==b", "b!=c", "c==a"}}, want: false},
		{name: "test5", args: args{[]string{"c==c", "b==d", "x!=z"}}, want: true},
		{name: "test_1", args: args{[]string{"c==c", "b==d", "a!=a", "x!=z"}}, want: false},
		{name: "test_2", args: args{[]string{"a!=a"}}, want: false},
		{name: "test_3", args: args{[]string{"c==c", "f!=a", "f==b", "b==c"}}, want: true},
		{name: "test_4", args: args{[]string{"e==d", "e==a", "f!=d", "b!=c", "a==b"}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equationsPossible(tt.args.equations); got != tt.want {
				t.Errorf("equationsPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
