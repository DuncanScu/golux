package golux

import (
	"github.com/DuncanScu/golux/common"
	"github.com/DuncanScu/golux/tree"
	"github.com/aws/aws-lambda-go/events"
)

type (
	Router struct {
		PathTree *tree.PathTree
	}
	HandlerFunc common.HandlerFunc
)

func NewRouter() *Router {
	return &Router{
		PathTree: tree.NewPathTree(),
	}
}

func (r *Router) Handle(method string, path string, fn common.HandlerFunc) {
	err := r.PathTree.Insert(method, path, fn)
	if err != nil {
		panic(err)
	}
}

func (r *Router) HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	routeNode, err := r.PathTree.Search(req.Path)
	if err != nil {
		panic(err)
	}

	handlerFn, err := routeNode.GetHandler(req.HTTPMethod)
	if err != nil {
		panic(err)
	}

	return handlerFn(req)
}
