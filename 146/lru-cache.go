package _146

type Node struct {
	K        int
	V        int
	PreNode  *Node
	NextNode *Node
}

// 2020.04.08
// https://leetcode-cn.com/problems/lru-cache/
type LRUCache struct {
	NodeMap  map[int]*Node
	Header   *Node
	Tail     *Node
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		NodeMap:  make(map[int]*Node),
		Header:   nil,
		Tail:     nil,
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.NodeMap[key]
	if !ok {
		return -1
	}
	this.addHeader(n)
	return n.V
}

func (this *LRUCache) Put(key int, value int) {
	n := &Node{
		K:        key,
		V:        value,
		PreNode:  nil,
		NextNode: nil,
	}
	if n, ok := this.NodeMap[key]; ok {
		n.V = value
		this.addHeader(n)
		return
	}
	if len(this.NodeMap) == this.Capacity {
		delN := this.removeTail()
		delete(this.NodeMap, delN.K)
		delN = nil
	}
	this.addHeader(n)
	if this.Tail == nil {
		this.Tail = n
	}
	this.NodeMap[key] = n
}

func (this *LRUCache) addHeader(n *Node) {
	if this.Header == nil {
		this.Header = n
		return
	}
	if n != this.Header {
		if n.PreNode != nil {
			n.PreNode.NextNode = n.NextNode
			if this.Tail == n {
				this.Tail = n.PreNode
			}
		}
		if n.NextNode != nil {
			n.NextNode.PreNode = n.PreNode
		}
		this.Header.PreNode = n
		n.PreNode = nil
		n.NextNode = this.Header
		this.Header = n
	}
}

func (this *LRUCache) removeTail() *Node {
	if this.Tail == nil {
		return nil
	}
	n := this.Tail
	if n.PreNode != nil {
		n.PreNode.NextNode = nil
	}
	this.Tail = n.PreNode
	n.PreNode = nil
	n.NextNode = nil
	return n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// ["LRUCache","put","put","get","put","get","put","get","get","get"]
//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

//["LRUCache","put", "put", "get", "get","get","get","get","get", "put", "get"]
//[[2], [1,1], [2,2], [1],[1],[1],[1],[1], [2], [3,3],[2]]
