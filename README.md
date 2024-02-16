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
	"github.com/prongbang/echoerror"
	"github.com/labstack/echo/v4"
)

func login() error {
	return echoerror.NewUnauthorized()
}

func main() {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		err := login()
		return echoerror.New(c).Response(err)
	})

	app.Logger.Fatal(app.Start(":1323"))
}
```

- Use by Custom Code

```go
package main

import (
	"http"
	"github.com/prongbang/echoerror"
	"github.com/labstack/echo/v4"
)

type CustomError struct {
	echoerror.Body
}

// Error implements error.
func (c *CustomError) Error() string {
	return c.Message
}

func NewCustomError() error {
	return &CustomError{
		Body: echoerror.Body{
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

func custom() error {
	return NewCustomError()
}

func main() {
	app := echo.New()

	customResp := NewCustomResponse()

	app.GET("/", func(c echo.Context) error {
		err := custom()
		return echoerror.New(c).Custom(customResp).Response(err)
	})

	app.Logger.Fatal(app.Start(":1323"))
}
```