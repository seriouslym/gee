package calc

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}

func (t *Trie) startsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}
func (t *Trie) search(prefix string) bool {
	node := t.searchPrefix(prefix)
	return node != nil && node.isEnd
}
