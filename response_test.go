package echoerror_test

import (
	"github.com/labstack/echo/v4"
	"github.com/prongbang/echoerror"
	"github.com/prongbang/goerror"
	"net/http"
	"net/http/httptest"
	"testing"
)

var response = echoerror.New()

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
			Code: "CUS001",
		},
	}
}

type customResponse struct {
}

// Response implements response.Custom.
func (c *customResponse) Response(ctx echo.Context, err error) error {
	switch e := err.(type) {
	case *CustomError:
		return ctx.JSON(http.StatusBadRequest, e)
	}
	return nil
}

func NewCustomResponse() echoerror.Custom {
	return &customResponse{}
}

func TestNewCustomError(t *testing.T) {
	app := echo.New()

	customResp := NewCustomResponse()
	res := echoerror.New(&echoerror.Config{
		Custom: &customResp,
	})

	handler := func(c echo.Context) error {
		return res.With(c).Response(NewCustomError())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusBadRequest {
		t.Error("Error", resp.Code)
	}
}

func TestNewUseProxy(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewUseProxy())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusUseProxy {
		t.Error("Error", resp.Code)
	}
}

func TestNewUnauthorized(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewUnauthorized())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusUnauthorized {
		t.Error("Error", resp.Code)
	}
}

func TestNewTemporaryRedirect(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewTemporaryRedirect())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusTemporaryRedirect {
		t.Error("Error", resp.Code)
	}
}

func TestNewNotFound(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewNotFound())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusNotFound {
		t.Error("Error", resp.Code)
	}
}

func TestNewSwitchingProtocols(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewSwitchingProtocols())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusSwitchingProtocols {
		t.Error("Error", resp.Code)
	}
}

func TestNewSeeOther(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewSeeOther())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusSeeOther {
		t.Error("Error", resp.Code)
	}
}

func TestNewResetContent(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewResetContent())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusResetContent {
		t.Error("Error", resp.Code)
	}
}

func TestNewRequestTimeout(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewRequestTimeout())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusRequestTimeout {
		t.Error("Error", resp.Code)
	}
}

func TestNewProxyAuthRequired(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewProxyAuthRequired())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusProxyAuthRequired {
		t.Error("Error", resp.Code)
	}
}

func TestNewProcessing(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewProcessing())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusProcessing {
		t.Error("Error", resp.Code)
	}
}

func TestNewPermanentRedirect(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewPermanentRedirect())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusPermanentRedirect {
		t.Error("Error", resp.Code)
	}
}

func TestNewPaymentRequired(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewPaymentRequired())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusPaymentRequired {
		t.Error("Error", resp.Code)
	}
}

func TestNewPartialContent(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewPartialContent())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusPartialContent {
		t.Error("Error", resp.Code)
	}
}

func TestNewOK(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewOK(nil))
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusOK {
		t.Error("Error", resp.Code)
	}
}

func TestNewNotModified(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewNotModified())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusNotModified {
		t.Error("Error", resp.Code)
	}
}

func TestNewNotAcceptable(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewNotAcceptable())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusNotAcceptable {
		t.Error("Error", resp.Code)
	}
}

func TestNewNonAuthoritativeInformation(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewNonAuthoritativeInformation())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusNonAuthoritativeInfo {
		t.Error("Error", resp.Code)
	}
}

func TestNewNoContent(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewNoContent())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusNoContent {
		t.Error("Error", resp.Code)
	}
}

func TestNewMultipleChoices(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewMultipleChoices())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusMultipleChoices {
		t.Error("Error", resp.Code)
	}
}

func TestNewMultiJSON(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewMultiStatus())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusMultiStatus {
		t.Error("Error", resp.Code)
	}
}

func TestNewMovedPermanently(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewMovedPermanently())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusMovedPermanently {
		t.Error("Error", resp.Code)
	}
}

func TestNewMethodNotAllowed(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewMethodNotAllowed())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusMethodNotAllowed {
		t.Error("Error", resp.Code)
	}
}

func TestNewIMUsed(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewIMUsed())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusIMUsed {
		t.Error("Error", resp.Code)
	}
}

func TestNewFound(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewFound())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusFound {
		t.Error("Error", resp.Code)
	}
}

func TestNewForbidden(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewForbidden())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusForbidden {
		t.Error("Error", resp.Code)
	}
}

func TestNewEarlyHints(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewEarlyHints())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusEarlyHints {
		t.Error("Error", resp.Code)
	}
}

func TestNewCreated(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewCreated(nil))
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusCreated {
		t.Error("Error", resp.Code)
	}
}

func TestNewContinue(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewContinue())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusContinue {
		t.Error("Error", resp.Code)
	}
}

func TestNewConflict(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewConflict())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusConflict {
		t.Error("Error", resp.Code)
	}
}

func TestNewBadRequest(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewBadRequest())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusBadRequest {
		t.Error("Error", resp.Code)
	}
}

func TestNewAlreadyReported(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewAlreadyReported())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusAlreadyReported {
		t.Error("Error", resp.Code)
	}
}

func TestNewAccepted(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewAccepted())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusAccepted {
		t.Error("Error", resp.Code)
	}
}

func TestNewGone(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewGone())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusGone {
		t.Error("Error", resp.Code)
	}
}

func TestNewLengthRequired(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewLengthRequired())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusLengthRequired {
		t.Error("Error", resp.Code)
	}
}

func TestNewPreconditionFailed(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewPreconditionFailed())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusPreconditionFailed {
		t.Error("Error", resp.Code)
	}
}

func TestNewRequestEntityTooLarge(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewRequestEntityTooLarge())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusRequestEntityTooLarge {
		t.Error("Error", resp.Code)
	}
}

func TestNewRequestURITooLong(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewRequestURITooLong())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusRequestURITooLong {
		t.Error("Error", resp.Code)
	}
}

func TestNewUnsupportedMediaType(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewUnsupportedMediaType())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusUnsupportedMediaType {
		t.Error("Error", resp.Code)
	}
}

func TestNewRequestedRangeNotSatisfiable(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewRequestedRangeNotSatisfiable())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusRequestedRangeNotSatisfiable {
		t.Error("Error", resp.Code)
	}
}

func TestNewExpectationFailed(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewExpectationFailed())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusExpectationFailed {
		t.Error("Error", resp.Code)
	}
}

func TestNewTeapot(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewTeapot())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusTeapot {
		t.Error("Error", resp.Code)
	}
}

func TestNewMisdirectedRequest(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewMisdirectedRequest())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusMisdirectedRequest {
		t.Error("Error", resp.Code)
	}
}

func TestNewUnprocessableEntity(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewUnprocessableEntity())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusUnprocessableEntity {
		t.Error("Error", resp.Code)
	}
}

func TestNewLocked(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewLocked())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusLocked {
		t.Error("Error", resp.Code)
	}
}

func TestNewFailedDependency(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewFailedDependency())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusFailedDependency {
		t.Error("Error", resp.Code)
	}
}

func TestNewTooEarly(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewTooEarly())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusTooEarly {
		t.Error("Error", resp.Code)
	}
}

func TestNewUpgradeRequired(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewUpgradeRequired())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusUpgradeRequired {
		t.Error("Error", resp.Code)
	}
}

func TestNewPreconditionRequired(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewPreconditionRequired())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusPreconditionRequired {
		t.Error("Error", resp.Code)
	}
}

func TestNewTooManyRequests(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewTooManyRequests())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusTooManyRequests {
		t.Error("Error", resp.Code)
	}
}

func TestNewRequestHeaderFieldsTooLarge(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewRequestHeaderFieldsTooLarge())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusRequestHeaderFieldsTooLarge {
		t.Error("Error", resp.Code)
	}
}

func TestNewUnavailableForLegalReasons(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewUnavailableForLegalReasons())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusUnavailableForLegalReasons {
		t.Error("Error", resp.Code)
	}
}

func TestNewInternalServerError(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewInternalServerError())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusInternalServerError {
		t.Error("Error", resp.Code)
	}
}

func TestNewNotImplemented(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewNotImplemented())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusNotImplemented {
		t.Error("Error", resp.Code)
	}
}

func TestNewBadGateway(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewBadGateway())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusBadGateway {
		t.Error("Error", resp.Code)
	}
}

func TestNewServiceUnavailable(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewServiceUnavailable())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusServiceUnavailable {
		t.Error("Error", resp.Code)
	}
}

func TestNewGatewayTimeout(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewGatewayTimeout())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusGatewayTimeout {
		t.Error("Error", resp.Code)
	}
}

func TestNewHTTPVersionNotSupported(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewHTTPVersionNotSupported())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusHTTPVersionNotSupported {
		t.Error("Error", resp.Code)
	}
}

func TestNewVariantAlsoNegotiates(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewVariantAlsoNegotiates())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusVariantAlsoNegotiates {
		t.Error("Error", resp.Code)
	}
}

func TestNewInsufficientStorage(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewInsufficientStorage())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusInsufficientStorage {
		t.Error("Error", resp.Code)
	}
}

func TestNewLoopDetected(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewLoopDetected())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusLoopDetected {
		t.Error("Error", resp.Code)
	}
}

func TestNewNotExtended(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewNotExtended())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusNotExtended {
		t.Error("Error", resp.Code)
	}
}

func TestNewNetworkAuthenticationRequired(t *testing.T) {
	app := echo.New()

	handler := func(c echo.Context) error {
		return response.With(c).Response(goerror.NewNetworkAuthenticationRequired())
	}
	app.GET("/test", handler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := app.NewContext(req, resp)

	_ = handler(c)

	if resp.Code != http.StatusNetworkAuthenticationRequired {
		t.Error("Error", resp.Code)
	}
}
