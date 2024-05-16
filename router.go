// Package golux provides a lightweight HTTP router designed for serverless applications.
//
// golux allows you to define routes and handlers for incoming HTTP requests in AWS Lambda functions,
// using a simple and efficient routing mechanism based on a tree structure.
package golux

import (
	"net/http"

	"github.com/DuncanScu/golux/common"
	"github.com/DuncanScu/golux/tree"
	"github.com/aws/aws-lambda-go/events"
)

type (
	// Router is a lightweight HTTP router.
	Router struct {
		PathTree *tree.PathTree
	}
	// HandlerFunc defines the signature for HTTP request handlers.
	HandlerFunc common.HandlerFunc
)

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	return &Router{
		PathTree: tree.NewPathTree(),
	}
}

// Handle registers a new handler for the specified HTTP method and path pattern.
func (r *Router) Handle(method string, path string, fn common.HandlerFunc) {
	err := r.PathTree.Insert(method, path, fn)
	if err != nil {
		panic(err)
	}
}

// HandleRequest handles an incoming API Gateway Proxy request and dispatches it to the appropriate handler.
// It returns an API Gateway Proxy response and an error, if any.
func (r *Router) HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	routeNode, err := r.PathTree.Search(req.Path)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadGateway}, err
	}

	handlerFn, err := routeNode.GetHandler(req.HTTPMethod)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadGateway}, err
	}

	return handlerFn(req)
}
