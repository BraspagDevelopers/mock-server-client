![License](https://img.shields.io/github/license/BraspagDevelopers/mock-server-client.svg?style=flat-square)
![Tag](https://img.shields.io/github/tag/BraspagDevelopers/mock-server-client.svg?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/BraspagDevelopers/mock-server-client?style-flat-square)](https://goreportcard.com/report/github.com/BraspagDevelopers/mock-server-client)

Mock Server Client - a Go package
=============================

Client for the [mock-server.com](https://mock-server.com) product.

The focus of this library is to use within integration tests in order to verify if a specific request was made in a mocked dependency (the mock-server).

It is still work in progress.


## Installation

Using `go get`
```go
go get -u github.com/BraspagDevelopers/mock-server-client
```

Using `go.mod` file
```go
require github.com/BraspagDevelopers/mock-server-client
```


## Usage

Import the package
```go
import (
    "github.com/BraspagDevelopers/mock-server-client"
)
```

### Verifying a request

The folowing verifies if the mock-server received requests matching the following filters:
* The method was `POST`
* The endpoint was `/api/categories`
* Having a header `Environment: Development`
* The body was a JSON and the body field `name` had the value `"Tools"`
* The request was made once and only once
```go
ms := mockserver.NewClient("localhost", 8080)
err := ms.Verify(
    mockserver.RequestMatcher{
        Method: http.MethodPost,
        Path:   "/api/categories"}.
        WithHeader("Environment", "Development").
        WithJSONFields(map[string]interface{}{
            "name": "Tools",
        }),
    mockserver.Once())
```

### Clearing requests
You can clear the mock-server requests in the log that matches an specific `RequestMatcher` using the method `MockClient.Verify`.
```go
err := ms.Verify(mockserver.RequestMatcher{
    Method: http.MethodPost,
    Path:   "/api/categories"}.
    WithHeader("Environment", "Development").
    WithJSONFields(map[string]interface{}{
        "name": "Tools",
    }))
```
Notice that the cardinality is not specified in this request for this method. All matching requests will be erased from the log.

### Verify and clear at the same time
You can also verify and clear matching requests at once by using the method `MockClient.VerifyAndClear()`.
