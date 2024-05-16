# Golux

Golux is a lightweight HTTP router designed specifically for use with AWS Lambda functions consuming ApiGatewayProxyRequest events. It simplifies the process of routing HTTP requests to specific Lambda handlers within your Go application.

## Features

- **Simple Routing**: Golux provides a straightforward API for defining routes and associating them with Lambda handlers.
- **Integration with AWS Lambda and API Gateway**: Golux is tailored to seamlessly work with Lambda functions that consume ApiGatewayProxyRequest events, making it an ideal choice for building serverless applications on AWS.

## Installation

To use Golux in your Go application, simply install it using `go get`:

```bash
go get -u github.com/DuncanScu/golux
```

## Usage

```go
//main.go
package main

func main(){
	mux := golux.NewRouter()
	mux.Handle(http.MethodGet, "/test", testHandler)
	lambda.Start(mux.HandleRequest)
}

func testHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{StatusCode: http.StatusAccepted}, nil
}
```

## Contributing

Contributions to Golux are welcome! If you have suggestions, bug reports, or would like to contribute code, please open an issue or submit a pull request on GitHub.
