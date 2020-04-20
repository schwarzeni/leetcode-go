package _0014

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{[]string{"flower", "flow", "flight"}}, want: "fl"},
		{name: "test2", args: args{[]string{"dog", "racecar", "car"}}, want: ""},
		{name: "test3", args: args{[]string{"dog"}}, want: "dog"},
		{name: "test4", args: args{[]string{""}}, want: ""},
		{name: "test5", args: args{[]string{"flower", "flower", "flowerrrr"}}, want: "flower"},
		{name: "test6", args: args{[]string{"flower", "", "flowerrrr"}}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
