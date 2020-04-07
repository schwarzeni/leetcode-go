package _208

type Trie struct {
	Root *Node
}

const END = '\n'

type Node struct {
	Data byte
	Next map[byte]*Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{&Node{Data: END, Next: make(map[byte]*Node)}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.addWord(this.Root, word, 0)
}

func (this *Trie) addWord(currNode *Node, word string, startIdx int) {
	// 考虑边界条件

	if startIdx == len(word) {
		if currNode.Data == END {
		} else {
			if _, ok := currNode.Next[END]; !ok {
				currNode.Next[END] = &Node{Data: END}
			}
		}
		return
	}

	currChar := word[startIdx]
	if _, ok := currNode.Next[currChar]; !ok {
		currNode.Next[currChar] = &Node{
			Data: currChar,
			Next: make(map[byte]*Node),
		}
	}
	this.addWord(currNode.Next[currChar], word, startIdx+1)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.search(this.Root, word, 0)
}

func (this *Trie) search(currNode *Node, word string, startIdx int) bool {

	if startIdx == len(word) {
		_, ok := currNode.Next[END]
		return ok
	}

	currChar := word[startIdx]

	if _, ok := currNode.Next[currChar]; !ok {
		return false
	}

	return this.search(currNode.Next[currChar], word, startIdx+1)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.startsWith(this.Root, prefix, 0)
}

func (this *Trie) startsWith(currNode *Node, word string, startIdx int) bool {

	if startIdx == len(word) {
		return true
	}

	currChar := word[startIdx]

	if _, ok := currNode.Next[currChar]; !ok {
		return false
	}

	return this.startsWith(currNode.Next[currChar], word, startIdx+1)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
