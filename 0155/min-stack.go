package _0155

import "math"

// 2020.05.12
// https://leetcode-cn.com/problems/min-stack/

type Node struct {
	v int
	m int
}

type MinStack struct {
	stack []Node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	n := Node{v: x, m: x}
	if len(this.stack) > 0 && this.stack[len(this.stack)-1].m < x {
		n.m = this.stack[len(this.stack)-1].m
	}
	this.stack = append(this.stack, n)
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return math.MinInt32
	}
	return this.stack[len(this.stack)-1].v
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return math.MinInt32
	}
	return this.stack[len(this.stack)-1].m
}

//type MinStack struct {
//	data     []int
//	sortData []int
//}
//
//func Constructor() MinStack {
//	return MinStack{}
//}
//
//func (this *MinStack) Push(x int) {
//	this.data = append(this.data, x)
//	this.sortData = append(this.sortData, x)
//	sort.Ints(this.sortData)
//}
//
//func (this *MinStack) Pop() {
//	if len(this.data) == 0 {
//		return
//	}
//	d := this.data[len(this.data)-1]
//	this.data = this.data[:len(this.data)-1]
//	i, j := 0, len(this.sortData)-1
//	for i <= j {
//		mid := i + (j-i)/2
//		if this.sortData[mid] == d {
//			this.sortData = append(this.sortData[:mid], this.sortData[mid+1:]...)
//			break
//		}
//		if this.sortData[mid] < d {
//			i = mid + 1
//		} else {
//			j = mid - 1
//		}
//	}
//}
//
//func (this *MinStack) Top() int {
//	if len(this.data) == 0 {
//		return 0
//	}
//	return this.data[len(this.data)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	if len(this.data) == 0 {
//		return math.MinInt32
//	}
//	return this.sortData[0]
//}
