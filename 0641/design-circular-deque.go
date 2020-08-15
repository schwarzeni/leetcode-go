package _0641

// 2020.08.15
// https://leetcode-cn.com/problems/design-circular-deque/
type MyCircularDeque struct {
    size int
    cap int
    header int
    tail int
    queue []int
}

const NULL_VAL = -1

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    queue := make([]int, k)
    for i := range queue {
        queue[i] = NULL_VAL
    }
    return MyCircularDeque{
        size: 0,
        cap: k,
        header: -1,
        tail: -1,
        queue: queue,
    }
}

func (this *MyCircularDeque) initPointer() (header int, tail int) {
    this.header, this.tail = 0, 0
    return this.header, this.tail
}

func (this *MyCircularDeque) nextHeader() int {
    if this.header == -1 {
        this.initPointer()
    } else if this.header == 0 {
        this.header = this.cap - 1
    } else {
        this.header --
    }
    return this.header
}

func (this *MyCircularDeque) nextTail() int {
    if this.tail == -1 {
        this.initPointer()
    } else if this.tail == this.cap - 1 {
        this.tail = 0
    } else {
        this.tail ++
    }
    return this.tail
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    this.queue[this.nextHeader()] = value
    this.size ++
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    this.queue[this.nextTail()] = value
    this.size ++
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    this.header ++
    if this.header == this.cap {
        this.header = 0
    }
    this.size --
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    this.tail --
    if this.tail == -1 {
        this.tail = this.cap - 1
    }
    this.size --
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.header]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.tail]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.size == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return this.size == this.cap
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
