package _0071

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{path: "/home/"}, want: "/home"},
		{name: "test2", args: args{path: "/../"}, want: "/"},
		{name: "test3", args: args{path: "/a/../../../../b"}, want: "/b"},
		{name: "test4", args: args{path: "/home//foo/"}, want: "/home/foo"},
		{name: "test5", args: args{path: "/home//foo//"}, want: "/home/foo"},
		{name: "test6", args: args{path: "/a/./b/../../c/"}, want: "/c"},
		{name: "test7", args: args{path: "/a/../../b/../c//.//"}, want: "/c"},
		{name: "test8", args: args{path: "/a//b////c/d//././/.."}, want: "/a/b/c"},
		{name: "test9", args: args{path: "/"}, want: "/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
