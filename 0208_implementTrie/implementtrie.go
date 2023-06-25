package implementtrie

import "vudangkhoa2906-leetcode-go/common"

type Trie struct {
	Root *common.TrieNode
}

func Constructor() Trie {
	return Trie{Root: common.NewTrieNode()}
}

func (this *Trie) Insert(word string) {
	node := this.Root
	for _, val := range []byte(word) {
		if !node.ContainsKey(val) {
			node.Put(val, common.NewTrieNode())
		}
		node = node.Get(val)
	}
	node.SetEnd()
}

func (this *Trie) Search(word string) bool {
	node := this.Root
	for _, val := range []byte(word) {
		if !node.ContainsKey(val) {
			return false
		}
		node = node.Get(val)
	}
	return node.Flag
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.Root
	for _, val := range []byte(prefix) {
		if !node.ContainsKey(val) {
			return false
		}
		node = node.Get(val)
	}
	return true
}
