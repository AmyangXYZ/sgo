package sweetygo

import (
	"net/url"
	"strings"
)

type node struct {
	children  []*node
	hasParams bool
	prefix    string
	methods   map[string]HandlerFunc
}

func (n *node) addNode(method, path string, handler HandlerFunc) {
	prefixes := strings.Split(path, "/")[1:]

	for count := len(prefixes); count > 0; count-- {
		aNode, prefix := n.traverse(prefixes, nil)
		if aNode.prefix == prefix && count == 1 { // updating an existing node
			aNode.methods[method] = handler
			return
		}

		newNode := node{
			prefix:  prefix,
			methods: make(map[string]HandlerFunc),
		}

		if len(prefix) > 0 {
			if prefix[0] == ':' {
				newNode.hasParams = true
			}
		}

		if count == 1 {
			newNode.methods[method] = handler
		}

		aNode.children = append(aNode.children, &newNode)
	}
}

func (n *node) traverse(prefixes []string, params url.Values) (*node, string) {
	prefix := prefixes[0]

	if len(n.children) > 0 { // not leaf
		for _, child := range n.children {

			if prefix == child.prefix || child.hasParams {
				if child.hasParams && params != nil {
					params.Add(child.prefix[1:], prefix)
				}

				next := prefixes[1:]
				if len(next) > 0 {
					return child.traverse(next, params) // evil recursion
				}

				return child, prefix
			}
		}
	}
	return n, prefix
}
