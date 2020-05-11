package _0155

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.GetMin() == -3)
	obj.Pop()
	fmt.Println(obj.Top() == 0)
	fmt.Println(obj.GetMin() == -2)

	obj2 := Constructor()
	obj2.Push(2147483646)
	obj2.Push(2147483646)
	obj2.Push(2147483647)
	obj2.Pop()
	obj2.Push(2147483647)
	fmt.Println(obj2.Top())
	fmt.Println(obj2.GetMin())

}
