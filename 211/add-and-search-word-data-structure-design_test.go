package _211

import (
	"fmt"
	"testing"
)

func TestWordDictionary_addWord(t *testing.T) {

	obj := Constructor()
	obj.AddWord("abc")
	obj.AddWord("abed")
	obj.AddWord("abcc")

	ok1 := obj.Search("abc")
	ok2 := obj.Search("abed")
	ok3 := obj.Search("abcc")
	ok4 := obj.Search("a..c")
	ok5 := obj.Search("...")
	ok6 := obj.Search("a.e.")

	fmt.Println(ok1, ok2, ok3, ok4, ok5, ok6) // expected to be true

	ok7 := obj.Search("cas")
	ok8 := obj.Search("abcd")
	ok9 := obj.Search("a.")
	fmt.Println(ok7, ok8, ok9) // expected to be false

	//
	//input := []string{"WordDictionary", "addWord", "addWord", "search", "search", "search", "search", "search", "search"}
	//searchAndExpect := struct {
	//	s        []string
	//	expected []bool
	//}{s: []string{"", "a", "a", ".", "a", "aa", "a", ".a", "a."},
	//	expected: []bool{false, false, false, false}}

	// ["WordDictionary","addWord","addWord","search","search","search","search","search","search"]
	//[[],["a"],["a"],["."],["a"],["aa"],["a"],[".a"],["a."]]

	//["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
	//[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
}
