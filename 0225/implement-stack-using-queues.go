package _0225

import "container/list"

// 2020.09.25
// https://leetcode-cn.com/problems/implement-stack-using-queues/
type MyStack struct {
	queue *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{queue: list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	tmpQ := list.New()
	ll := this.queue.Len() - 1
	for i := 0; i < ll; i++ {
		e := this.queue.Front()
		tmpQ.PushBack(e.Value)
		this.queue.Remove(e)
	}

	topE := this.queue.Front()
	res := topE.Value.(int)
	this.queue.Remove(topE)

	for tmpQ.Len() > 0 {
		e := tmpQ.Front()
		this.queue.PushBack(e.Value)
		tmpQ.Remove(e)
	}
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	tmpQ := list.New()
	ll := this.queue.Len() - 1
	for i := 0; i < ll; i++ {
		e := this.queue.Front()
		tmpQ.PushBack(e.Value)
		this.queue.Remove(e)
	}
	topE := this.queue.Front()
	res := topE.Value.(int)
	this.queue.Remove(topE)
	tmpQ.PushBack(res)

	for tmpQ.Len() > 0 {
		e := tmpQ.Front()
		this.queue.PushBack(e.Value)
		tmpQ.Remove(e)
	}
	return res
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

//  ["MyStack","push","push","top","pop","empty", "top", "pop", "empty"]
// [[],[1],[2],[],[],[], [], [], []]
