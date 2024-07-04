<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [README](#readme)
    - [About the project](#about-the-project)
    - [Status](#status)
    - [Layout](#layout)
    - [Install](#install)
    - [Usage](#usage)
      - [Supported Methods](#supported-methods)
      - [Make Request](#make-request)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# README #

The easiest way to create request in go

### About the project ###

Request handler go library

### Status ###

Beta

### Layout ###

- __ap\-request\-go__
  - [LICENSE](LICENSE)
  - [README.md](README.md)
  - [go.mod](go.mod)
  - [go.sum](go.sum)
  - [main.go](main.go)
  - __pkg__
    - __app__
      - __auth__
        - [auth.go](pkg/app/auth/auth.go)
      - __domain__
        - [auth\_basic.go](pkg/app/domain/auth_basic.go)
        - [auth\_methods.go](pkg/app/domain/auth_methods.go)
        - [http\_methods.go](pkg/app/domain/http_methods.go)
        - [model.go](pkg/app/domain/model.go)
      - __request__
        - [request.go](pkg/app/request/request.go)
  - __tests__
    - [request\_test.go](tests/request_test.go)

### Install ###
`go get github.com/andresPirona/ap-request-go`

### Usage ###
#### Supported Methods ###
```

- GET
- POST

```


#### Make Request ###
```

package main

import (
"fmt"
ap_request_go "github.com/andresPirona/ap-request-go"
"github.com/andresPirona/ap-request-go/pkg/app/domain"
"net/http"
)

type EntityResponse struct {
ID   string `json:"id"`
Name string `json:"name"`
}

func main() {

	var testRequest = domain.Options{
		Method:  http.MethodGet,
		Url:     "https://6413050a3b710647375ca7be.mockapi.io/api/v1/users/1",
		Body:    nil,
		Auth:    nil,
		Headers: nil,
		Params:  nil,
	}

	entityResponse, response, err := ap_request_go.NewRequest[EntityResponse](testRequest)
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(response)
	fmt.Println(entityResponse)

}

```