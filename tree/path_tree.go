package tree

import (
	"errors"
	"strings"

	"github.com/DuncanScu/golux/common"
)

type PathTree struct {
	Root *RouteNode
}

func NewPathTree() *PathTree {
	root := NewRouteNode("")
	return &PathTree{
		Root: root,
	}
}

func (t *PathTree) Insert(method string, path string, fn common.HandlerFunc) error {
	current := t.Root
	partSegments := strings.Split(path, "/")
	for _, segment := range partSegments {
		if _, childFound := current.Children[segment]; !childFound {
			current.Children[segment] = NewRouteNode(segment)
		}
		current = current.Children[segment]
	}
	return current.AddHandler(method, fn)
}

func (t *PathTree) Search(path string) (*RouteNode, error) {
	current := t.Root
	partSegments := strings.Split(path, "/")
	for _, segment := range partSegments {
		if _, exist := current.Children[segment]; !exist {
			return nil, errors.New("the handler for that request does not exist")
		}
		current = current.Children[segment]
	}
	return current, nil
}
