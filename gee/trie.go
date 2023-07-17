package gee

import "strings"

type trie struct {
	pattern  string
	part     string
	children map[string]*trie
	isWild   bool
}

func (t *trie) insert(pattern string, parts []string) {
	node := t
	for _, part := range parts {

		if node.children[part] == nil {
			node.children[part] = &trie{
				part:     part,
				isWild:   part[0] == ':' || part[0] == '*',
				children: make(map[string]*trie),
			}
		}
		node = node.children[part]
	}
	node.pattern = pattern
}

// search 不能光考虑map[part]为不为空 还要考虑:param 以及路径深度问题
func (t *trie) search(parts []string, depth int) *trie {
	if depth == len(parts) || strings.HasPrefix(t.part, "*") {
		// 排除/p/:test/xx模式下 /p/hhh 匹配成功的情况 需要判断节点的模式是否注册过
		if t.pattern == "" {
			return nil
		}
		return t
	}
	part := parts[depth]
	nodes := make([]*trie, 0)
	for _, v := range t.children {
		if v.part == part || v.isWild {
			nodes = append(nodes, v)
		}
	}
	for _, node := range nodes {
		result := node.search(parts, depth+1)
		if result != nil {
			return result
		}
	}
	return nil
}
