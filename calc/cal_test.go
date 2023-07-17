package calc

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	trie := Trie{}
	trie.insert("app")
	trie.insert("apphhha")
	trie.insert("bother")
	fmt.Println(trie.search("app"))
	fmt.Println(trie.search("fsdf"))
	fmt.Println(trie.startsWith("bot"))
}
