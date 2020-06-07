package _0126

import (
	"reflect"
	"testing"
)

func Test_findLadders(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "test1", args: args{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
		}, want: [][]string{
			{"hit", "hot", "dot", "dog", "cog"},
			{"hit", "hot", "lot", "log", "cog"},
		}},
		{name: "test2", args: args{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
		}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
