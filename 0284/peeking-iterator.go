package _0284

// 2020.09.05
// https://leetcode-cn.com/problems/peeking-iterator/

type PeekingIterator struct {
	cache []int
	iter  *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	pIter := &PeekingIterator{iter: iter}
	return pIter
}

func (this *PeekingIterator) hasNext() bool {
	return len(this.cache) > 0 || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if len(this.cache) > 0 {
		v := this.cache[0]
		this.cache = this.cache[1:]
		return v
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if len(this.cache) > 0 {
		return this.cache[0]
	}
	v := this.iter.next()
	this.cache = append(this.cache, v)
	return v
}

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	panic("")
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	panic("")
}
