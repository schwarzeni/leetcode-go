package _211

// // 2020.04.07
// https://leetcode-cn.com/problems/add-and-search-word-data-structure-design/
type WordDictionary struct {
	Root *Node
}

const END = '\n'
const ANY = '.'

type Node struct {
	Data byte
	Next map[byte]*Node
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{&Node{Data: END, Next: make(map[byte]*Node)}}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.addWord(this.Root, word, 0)
}

func (this *WordDictionary) addWord(currNode *Node, word string, startIdx int) {
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

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.search(this.Root, word, 0)
}

func (this *WordDictionary) search(currNode *Node, word string, startIdx int) bool {

	if startIdx == len(word) {
		_, ok := currNode.Next[END]
		return ok
	}

	currChar := word[startIdx]
	if currChar == ANY {
		for _, v := range currNode.Next {
			if this.search(v, word, startIdx+1) {
				return true
			}
		}
		return false
	}

	if _, ok := currNode.Next[currChar]; !ok {
		return false
	}

	return this.search(currNode.Next[currChar], word, startIdx+1)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
