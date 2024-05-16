package tree

import (
	"errors"

	"github.com/DuncanScu/golux/common"
)

type RouteNode struct {
	PathSegment     string
	Children        map[string]*RouteNode
	RequestHandlers map[string]common.HandlerFunc
}

func NewRouteNode(path string) *RouteNode {
	return &RouteNode{
		PathSegment:     path,
		Children:        map[string]*RouteNode{},
		RequestHandlers: make(map[string]common.HandlerFunc),
	}
}

func (n *RouteNode) AddHandler(method string, fn common.HandlerFunc) error {
	if _, exists := n.RequestHandlers[method]; exists {
		return errors.New("A handler for the path already exists")
	}
	n.RequestHandlers[method] = fn
	return nil
}

func (n *RouteNode) GetHandler(method string) (common.HandlerFunc, error) {
	handler, exists := n.RequestHandlers[method]
	if !exists {
		return nil, errors.New("A handler for that path does not exist")
	}
	return handler, nil
}
