package _208

import (
	"testing"
)

func TestTrie_StartsWith(t *testing.T) {

	obj := Constructor()
	obj.Insert("apple")
	if obj.Search("app") != false {
		t.Errorf("Search app in apple failed")
	}
	if obj.Search("apple") != true {
		t.Errorf("Search apple in apple failed")
	}
	if obj.StartsWith("app") != true {
		t.Errorf("StartsWith app in apple failed")
	}
}
