package _0117

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{name: "test1", args: args{&Node{Val: 1,
			Left:  &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}},
			Right: &Node{Val: 3, Right: &Node{Val: 7}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
