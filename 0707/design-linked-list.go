package _0707

// 2020.08.15
// https://leetcode-cn.com/problems/design-linked-list/
type MyNode struct {
    val int
    prev *MyNode
    next *MyNode
}

func NewMyNode (val int) *MyNode {
    return &MyNode{val: val}
}


type MyLinkedList struct {
    head *MyNode
    tail *MyNode
    size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if n := this.getNode(index); n != nil {
        return n.val
    }
    return -1
}

func (this *MyLinkedList) getNode(index int) *MyNode {
    if this.size <= index {
        return nil
    }
    curr := this.head
    for curr != nil && index >= 0{
        if index == 0 {
            return curr
        }
        curr = curr.next
        index --
    }
    return nil
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    n := NewMyNode(val)
    this.size ++
    if this.head == nil {
        this.head = n
        this.tail = n
        return
    }
    n.next = this.head
    this.head.prev = n
    this.head = n
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    n := NewMyNode(val)
    this.size ++
    if this.tail == nil {
        this.head = n
        this.tail = n
        return
    }
    n.prev = this.tail
    this.tail.next = n
    this.tail = n
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 { // ...
        this.AddAtHead(val)
        return
    }
    if index > this.size {
        return
    }
    if index == this.size {
        this.AddAtTail(val)
        return
    }
    n := NewMyNode(val)
    nextNode := this.getNode(index)
    n.next = nextNode
    n.prev = nextNode.prev
    if nextNode.prev != nil {
        nextNode.prev.next = n
    }
    nextNode.prev = n
    if nextNode == this.head {
        this.head = n
    }
    this.size ++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.size {
        return
    }
    n := this.getNode(index)
    prevN, nextN := n.prev, n.next
    if prevN != nil {
        prevN.next = nextN
    }
    if nextN != nil {
        nextN.prev = prevN
    }
    if n == this.head {
        this.head = nextN
    }
    if n == this.tail {
        this.tail = prevN
    }
    this.size --
    n = nil
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
