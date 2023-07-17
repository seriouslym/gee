package gee

import (
	"fmt"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "/assets/*filepath", nil)
	return r
}
func TestName(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "geektutu" {
		t.Fatal("name should be equal to 'geektutu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
	//trie := trie{
	//	children: make(map[string]*trie),
	//}
	//trie.insert("/test/help", []string{"test", "help"})
	//trie.insert("/test/:lang/", []string{"test", ":lang"})
	//trie.insert("/test/help/hhhh", []string{"test", "help", "hhhh"})
	//node := trie.search([]string{"test", ":lang"})
	//fmt.Println(node.pattern, node.part)

}
