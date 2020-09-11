package _1206

import (
	"math/rand"
	"time"
)

// 2020.09.10 - 2020.09.11
// https://leetcode-cn.com/problems/design-skiplist/
const MaxLevel = 16

type node struct {
	val   int
	nexts []*node
}

func (n *node) Level() int {
	return len(n.nexts) - 1
}

func newNode(val int, prevs []*node) *node {
	n := &node{val: val}
	n.nexts = make([]*node, MaxLevel+1)
	if l := len(prevs); l > 0 {
		n.nexts = make([]*node, n.getRandLevel()+1)
		for i := 1; i <= n.Level(); i++ {
			n.nexts[i] = prevs[i].nexts[i]
			prevs[i].nexts[i] = n
		}
	}
	return n
}

func (n *node) getRandLevel() int {
	level := 1
	rand.Seed(time.Now().UnixNano())
	for rand.Intn(2) != 0 && level < MaxLevel {
		level++
	}
	return level
}

type Skiplist struct {
	header *node
}

func Constructor() Skiplist {
	return Skiplist{header: &node{
		nexts: make([]*node, MaxLevel+1),
	}}
}

func (this *Skiplist) search(num int) ([]*node, *node) {
	var (
		prevs     = make([]*node, MaxLevel+1)
		foundNode *node
		curr      = this.header
	)
	for l := MaxLevel; l > 0; l-- {
		for {
			if curr.nexts[l] != nil && curr.nexts[l].val == num {
				foundNode = curr.nexts[l]
				break
			}
			if curr.nexts[l] == nil ||
				(curr == this.header && curr.nexts[l].val > num) ||
				(curr.val < num && curr.nexts[l].val > num) {
				break
			}
			curr = curr.nexts[l]
		}
		prevs[l] = curr
	}
	return prevs, foundNode
}

func (this *Skiplist) Add(num int) {
	prevs, _ := this.search(num)
	newNode(num, prevs)
}

func (this *Skiplist) Search(target int) bool {
	_, n := this.search(target)
	return n != nil
}

func (this *Skiplist) Erase(num int) bool {
	prevs, n := this.search(num)
	if n == nil {
		return false
	}
	if l := len(prevs); l > 0 {
		for i := 1; i <= n.Level(); i++ {
			prevs[i].nexts[i] = n.nexts[i]
		}
	}
	n = nil
	return true
}
