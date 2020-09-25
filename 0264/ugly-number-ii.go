package _0264

import "container/heap"

// 2020.09.25
// https://leetcode-cn.com/problems/ugly-number-ii/
func nthUglyNumber(n int) int {
	mark := make(map[int]struct{})
	mark[1] = struct{}{}
	count := 0
	i := 0
	h := &IntHeap{1}
	for {
		i = heap.Pop(h).(int)
		count++
		if count == n {
			break
		}
		for _, v := range []int{2, 3, 5} {
			if _, ok := mark[i*v]; !ok {
				mark[i*v] = struct{}{}
				heap.Push(h, i*v)
			}
		}
	}
	return i
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
