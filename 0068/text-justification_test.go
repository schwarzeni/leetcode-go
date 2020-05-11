package _0068

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test1", args: args{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
		}, want: []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		}},
		{name: "test2", args: args{
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
		}, want: []string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				for idx, _ := range tt.want {
					g, w := got[idx], tt.want[idx]
					//fmt.Println([]byte(g))
					//fmt.Println([]byte(w))
					if !reflect.DeepEqual(g, w) {
						fmt.Printf("line %d not equal\n", idx)
						fmt.Println(g)
						fmt.Println(w)
					}
				}
				t.Errorf("fullJustify() = \n%v, want \n%v", printS(got), printS(tt.want))
			}
		})
	}
}

func printS(ss []string) string {
	sss := ""
	for _, s := range ss {
		sss += s
		sss += "\n"
	}
	return sss
}
