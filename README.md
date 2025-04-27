# echoerror 🚨

[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/echoerror)](https://goreportcard.com/report/github.com/prongbang/echoerror)
[![Go Reference](https://pkg.go.dev/badge/github.com/prongbang/echoerror.svg)](https://pkg.go.dev/github.com/prongbang/echoerror)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/prongbang/echoerror.svg)](https://golang.org)

> Elegant error response handling for Echo web framework with built-in support for standard HTTP errors and custom error types.

## ✨ Features

- 🎯 **Simple Integration** - Easy to integrate with existing Echo applications
- 🔧 **Custom Error Types** - Define your own error types and response formats
- 🌐 **Standard HTTP Errors** - Built-in support for all standard HTTP status codes
- ⚡ **Lightweight** - Minimal dependencies and overhead
- 📦 **Type-Safe** - Go type-safe error handling
- 🔌 **Extensible** - Flexible configuration for custom needs

## 📦 Installation

```shell
go get github.com/prongbang/echoerror
```

## 🚀 Quick Start

### Basic Usage with Standard HTTP Status

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

## 🛠️ Advanced Usage

### Custom Error Types

Create and handle custom error types with specific response formats:

```go
package main

import (
    "net/http"
    "github.com/prongbang/goerror"
    "github.com/prongbang/echoerror"
    "github.com/labstack/echo/v4"
)

// Define custom error type
type CustomError struct {
    goerror.Body
}

func (c *CustomError) Error() string {
    return c.Message
}

func NewCustomError() error {
    return &CustomError{
        Body: goerror.Body{
            Code:    "CUS001",
            Message: "Custom error occurred",
        },
    }
}

// Define custom response handler
type customResponse struct{}

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
    
    // Configure custom response
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

## 📚 Error Types

### Standard HTTP Errors

Work with all standard HTTP error types from the `goerror` package:

```go
// 4xx Client Errors
goerror.NewBadRequest()
goerror.NewUnauthorized()
goerror.NewForbidden()
goerror.NewNotFound()
goerror.NewMethodNotAllowed()
goerror.NewConflict()
goerror.NewUnprocessableEntity()
goerror.NewTooManyRequests()

// 5xx Server Errors
goerror.NewInternalServerError()
goerror.NewNotImplemented()
goerror.NewBadGateway()
goerror.NewServiceUnavailable()
goerror.NewGatewayTimeout()
```

### Custom Error Structure

Define custom errors with additional fields:

```go
type ValidationError struct {
    goerror.Body
    Errors map[string]string `json:"errors"`
}

func NewValidationError(errors map[string]string) error {
    return &ValidationError{
        Body: goerror.Body{
            Code:    "VAL001",
            Message: "Validation failed",
        },
        Errors: errors,
    }
}
```

## 🔍 Examples

### API Error Handling

```go
app.POST("/users", func(c echo.Context) error {
    var user User
    if err := c.Bind(&user); err != nil {
        return response.With(c).Response(goerror.NewBadRequest("Invalid request body"))
    }
    
    if err := validator.Validate(user); err != nil {
        return response.With(c).Response(NewValidationError(err))
    }
    
    // Process user...
    return c.JSON(http.StatusCreated, user)
})
```

### Multiple Error Types

```go
type customResponse struct{}

func (c *customResponse) Response(ctx echo.Context, err error) error {
    switch resp := err.(type) {
    case *ValidationError:
        return ctx.JSON(http.StatusBadRequest, resp)
    case *AuthenticationError:
        return ctx.JSON(http.StatusUnauthorized, resp)
    case *ForbiddenError:
        return ctx.JSON(http.StatusForbidden, resp)
    default:
        return ctx.JSON(http.StatusInternalServerError, map[string]string{
            "code":    "INTERNAL_ERROR",
            "message": "An unexpected error occurred",
        })
    }
}
```

## ⚙️ Configuration Options

### echoerror.Config

| Option | Type | Description |
|--------|------|-------------|
| `Custom` | `*Custom` | Custom error response handler |

### Error Response Format

Standard response structure:

```json
{
    "code": "CUS001",
    "message": "Custom error message"
}
```

With additional fields:

```json
{
    "code": "VAL001",
    "message": "Validation failed",
    "errors": {
        "email": "Invalid email format",
        "password": "Password too short"
    }
}
```

## 🔧 Best Practices

1. **Use Standard Errors** - Prefer standard HTTP errors when possible
2. **Consistent Format** - Keep error response format consistent
3. **Meaningful Codes** - Use descriptive error codes
4. **Clear Messages** - Provide helpful error messages
5. **Status Codes** - Use appropriate HTTP status codes

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 💖 Support the Project

If you find this package helpful, please consider supporting it:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

## 🔗 Related Projects

- [goerror](https://github.com/prongbang/goerror) - Core error interface for Go
- [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework
- [fibererror](https://github.com/prongbang/fibererror) - Error handler for Fiber framework

---
