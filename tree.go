package sgo

import (
	"net/url"
	"strings"
)

// Trie node.
type Trie struct {
	children  []*Trie
	param     byte
	component string
	methods   map[string]HandlerFunc
}

// Insert a node into the tree.
func (t *Trie) Insert(method, path string, handler HandlerFunc) {
	components := strings.Split(path, "/")[1:]
Next:
	for _, component := range components {
		for _, child := range t.children {
			if child.component == component {
				t = child
				continue Next
			}
		}
		newNode := &Trie{component: component,
			methods: make(map[string]HandlerFunc)}
		if len(component) > 0 {
			if component[0] == ':' || component[0] == '*' {
				newNode.param = component[0]
			}
		}
		t.children = append(t.children, newNode)
		t = newNode
	}
	t.methods[method] = handler
}

// Search the tree.
func (t *Trie) Search(components []string, params url.Values) *Trie {
Next:
	for _, component := range components {
		for _, child := range t.children {
			if child.component == component || child.param == ':' || child.param == '*' {
				if child.param == '*' {
					return child
				}
				if child.param == ':' {
					params.Add(child.component[1:], component)
				}
				t = child
				continue Next
			}
		}
		return nil // not found
	}
	return t
}
