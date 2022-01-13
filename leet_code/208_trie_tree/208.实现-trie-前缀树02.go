package trietree

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type TrieM struct {
	children map[uint8]*TrieM
	isWord   bool
}

func ConstructorM() TrieM {
	return TrieM{children: map[uint8]*TrieM{}}
}

func (this *TrieM) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		if _, ok := root.children[word[i]]; !ok {
			trieM := ConstructorM()
			root.children[word[i]] = &trieM
		}
		root = root.children[word[i]]
	}
	root.isWord = true
}

func (this *TrieM) SearchPrefix(prefix string) *TrieM {
	root := this
	for i := 0; i < len(prefix); i++ {
		if _, ok := root.children[prefix[i]]; ok {
			root = root.children[prefix[i]]
			continue
		}
		return nil
	}
	return root
}

func (this *TrieM) Search(word string) bool {
	if trie := this.SearchPrefix(word); trie != nil {
		return trie.isWord
	}
	return false
}

func (this *TrieM) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
