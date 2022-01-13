package trietree

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start

type Trie struct {
	children [26]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	root := t
	for i := 0; i < len(word); i++ {
		if root.children[word[i]-'a'] == nil {
			root.children[word[i]-'a'] = new(Trie)
		}
		root = root.children[word[i]-'a']
	}
	root.isWord = true
}

func (t *Trie) Search(word string) bool {
	if trie := t.SearchPrefix(word); trie == nil {
		return trie.isWord
	}
	return false
}

func (t *Trie) SearchPrefix (prefix string) *Trie {
	root := t
	for i := 0; i < len(prefix); i++ {
		if root.children[prefix[i]-'a'] != nil {
			root = root.children[prefix[i]-'a']
			continue
		}
		return nil
	}
	return root
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
