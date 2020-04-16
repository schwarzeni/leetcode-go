package _146

import (
	"fmt"
	"testing"
)

func Test_common(t *testing.T) {
	obj := Constructor(2)
	if r := obj.Get(1); -1 != r {
		t.Errorf("get(%v) = %v, %v", 1, -1, r)
	}
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println()

	obj.Put(3, 3)
	fmt.Println()

	obj.Put(4, 4)
	fmt.Println()

	obj.Put(3, 10)
	fmt.Println()

	obj2 := Constructor(2)
	obj2.Put(1, 1)
	obj2.Put(2, 2)
	obj2.Get(1)
	obj2.Put(3, 3)
	obj2.Get(2)
	obj2.Put(4, 4)
	obj2.Get(1)
	obj2.Get(3)
	obj2.Get(4)
}
