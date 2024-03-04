# echoerror

### Install

```shell
go get github.com/prongbang/echoerror
```

### How to use

- Use by Standard HTTP Status Code

```go
package main

import (
	"github.com/prongbang/goerror"
	"github.com/prongbang/echoerror"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	response := echoerror.New()
	
	app.GET("/", func(c echo.Context) error {
		return response.With(c).Response(goerror.NewUnauthorized())
	})

	app.Logger.Fatal(app.Start(":1323"))
}
```

- Use by Custom Code

```go
package main

import (
	"http"
	"github.com/prongbang/goerror"
	"github.com/prongbang/echoerror"
	"github.com/labstack/echo/v4"
)

type CustomError struct {
	goerror.Body
}

// Error implements error.
func (c *CustomError) Error() string {
	return c.Message
}

func NewCustomError() error {
	return &CustomError{
		Body: goerror.Body{
			Code:    "CUS001",
			Message: "Custom 001",
		},
	}
}

type customResponse struct {
}

// Response implements response.Custom.
func (c *customResponse) Response(ctx echo.Context, err error) error {
	switch resp := err.(type) {
	case *CustomError:
		return ctx.JSON(http.StatusBadRequest, resp)
	}
	return nil
}

func NewCustomResponse() echoerror.Custom {
	return &customResponse{}
}

func main() {
	app := echo.New()

	customResp := NewCustomResponse()
	response := echoerror.New(&echoerror.Config{
		Custom: &customResp,
    })

	app.GET("/", func(c echo.Context) error {
		return response.With(c).Response(NewCustomError())
	})

	app.Logger.Fatal(app.Start(":1323"))
}
```