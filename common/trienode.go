package common

type TrieNode struct {
	Links []*TrieNode
	Flag  bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Links: make([]*TrieNode, 26)}
}

func (tN *TrieNode) ContainsKey(key byte) bool {
	return tN.Links[key-'a'] != nil
}

func (tN *TrieNode) Put(key byte, node *TrieNode) {
	tN.Links[key-'a'] = node
}

func (tN *TrieNode) Get(key byte) *TrieNode {
	return tN.Links[key-'a']
}

func (tN *TrieNode) SetEnd() {
	tN.Flag = true
}
