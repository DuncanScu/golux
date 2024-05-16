package tree_test

import (
	"net/http"
	"testing"

	"github.com/DuncanScu/golux/tree"
	"github.com/aws/aws-lambda-go/events"
)

func TestNewRouteNode(t *testing.T) {
	path := "/test"
	node := tree.NewRouteNode(path)

	if node.PathSegment != path {
		t.Errorf("Expected PathSegment to be %s, got %s", path, node.PathSegment)
	}

	if node.Children == nil {
		t.Error("Expected Children map to be initialized")
	}

	if node.RequestHandlers == nil {
		t.Error("Expected RequestHandlers map to be initialized")
	}
}

func TestGetHandler(t *testing.T) {
	node := tree.NewRouteNode("/test")
	method := http.MethodGet
	handlerFunc := func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return events.APIGatewayProxyResponse{}, nil
	}

	_, err := node.GetHandler(method)
	if err == nil {
		t.Errorf("Expected: %s | Actual: %s", tree.HandlerDoesntExistError, err)
	}

	_ = node.AddHandler(method, handlerFunc)

	fn, err := node.GetHandler(method)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if fn == nil {
		t.Error("Expected handler function to be returned.")
	}
}

func TestAddHandler(t *testing.T) {
	node := tree.NewRouteNode("/test")
	method := "GET"
	handlerFunc := func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return events.APIGatewayProxyResponse{}, nil
	}

	err := node.AddHandler(method, handlerFunc)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = node.AddHandler(method, handlerFunc)
	if err == nil {
		t.Errorf("Expected %s error, got: %v", tree.HandlerExistsError, err)
	}
}
