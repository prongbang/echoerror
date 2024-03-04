package echoerror

import (
	"github.com/labstack/echo/v4"
	"github.com/prongbang/goerror"
	"net/http"
)

type Config struct {
	Custom *Custom
	I18n   *I18n
}

type I18n struct {
	Enabled  bool
	Localize func(c echo.Context, code string) (string, error)
}

type Custom interface {
	Response(ctx echo.Context, err error) error
}

type Response interface {
	With(c echo.Context) HttpResponse
}

type HttpResponse interface {
	Response(err error) error
}

type response struct {
	Cus  *Custom
	I18n *I18n
}

type httpResponse struct {
	Ctx  echo.Context
	Cus  *Custom
	I18n *I18n
}

// With implements Response.
func (r *response) With(c echo.Context) HttpResponse {
	return &httpResponse{
		Ctx:  c,
		Cus:  r.Cus,
		I18n: r.I18n,
	}
}

// Response implements Response.
func (s *httpResponse) Response(err error) error {
	switch e := err.(type) {
	// Information
	case *goerror.Continue:
		return s.Ctx.JSON(http.StatusContinue, e)
	case *goerror.SwitchingProtocols:
		return s.Ctx.JSON(http.StatusSwitchingProtocols, e)
	case *goerror.Processing:
		return s.Ctx.JSON(http.StatusProcessing, e)
	case *goerror.EarlyHints:
		return s.Ctx.JSON(http.StatusEarlyHints, e)

	// Successful
	case *goerror.OK:
		return s.Ctx.JSON(http.StatusOK, e)
	case *goerror.Created:
		return s.Ctx.JSON(http.StatusCreated, e)
	case *goerror.Accepted:
		return s.Ctx.JSON(http.StatusAccepted, e)
	case *goerror.NonAuthoritativeInformation:
		return s.Ctx.JSON(http.StatusNonAuthoritativeInfo, e)
	case *goerror.NoContent:
		return s.Ctx.JSON(http.StatusNoContent, e)
	case *goerror.ResetContent:
		return s.Ctx.JSON(http.StatusResetContent, e)
	case *goerror.PartialContent:
		return s.Ctx.JSON(http.StatusPartialContent, e)
	case *goerror.MultiStatus:
		return s.Ctx.JSON(http.StatusMultiStatus, e)
	case *goerror.AlreadyReported:
		return s.Ctx.JSON(http.StatusAlreadyReported, e)
	case *goerror.IMUsed:
		return s.Ctx.JSON(http.StatusIMUsed, e)

	// Redirection
	case *goerror.MultipleChoices:
		return s.Ctx.JSON(http.StatusMultipleChoices, e)
	case *goerror.MovedPermanently:
		return s.Ctx.JSON(http.StatusMovedPermanently, e)
	case *goerror.Found:
		return s.Ctx.JSON(http.StatusFound, e)
	case *goerror.SeeOther:
		return s.Ctx.JSON(http.StatusSeeOther, e)
	case *goerror.NotModified:
		return s.Ctx.JSON(http.StatusNotModified, e)
	case *goerror.UseProxy:
		return s.Ctx.JSON(http.StatusUseProxy, e)
	case *goerror.TemporaryRedirect:
		return s.Ctx.JSON(http.StatusTemporaryRedirect, e)
	case *goerror.PermanentRedirect:
		return s.Ctx.JSON(http.StatusPermanentRedirect, e)

	// Client error
	case *goerror.BadRequest:
		return s.Ctx.JSON(http.StatusBadRequest, e)
	case *goerror.Unauthorized:
		return s.Ctx.JSON(http.StatusUnauthorized, e)
	case *goerror.PaymentRequired:
		return s.Ctx.JSON(http.StatusPaymentRequired, e)
	case *goerror.Forbidden:
		return s.Ctx.JSON(http.StatusForbidden, e)
	case *goerror.NotFound:
		return s.Ctx.JSON(http.StatusNotFound, e)
	case *goerror.MethodNotAllowed:
		return s.Ctx.JSON(http.StatusMethodNotAllowed, e)
	case *goerror.NotAcceptable:
		return s.Ctx.JSON(http.StatusNotAcceptable, e)
	case *goerror.ProxyAuthRequired:
		return s.Ctx.JSON(http.StatusProxyAuthRequired, e)
	case *goerror.RequestTimeout:
		return s.Ctx.JSON(http.StatusRequestTimeout, e)
	case *goerror.Conflict:
		return s.Ctx.JSON(http.StatusConflict, e)
	case *goerror.Gone:
		return s.Ctx.JSON(http.StatusGone, e)
	case *goerror.LengthRequired:
		return s.Ctx.JSON(http.StatusLengthRequired, e)
	case *goerror.PreconditionFailed:
		return s.Ctx.JSON(http.StatusPreconditionFailed, e)
	case *goerror.RequestEntityTooLarge:
		return s.Ctx.JSON(http.StatusRequestEntityTooLarge, e)
	case *goerror.RequestURITooLong:
		return s.Ctx.JSON(http.StatusRequestURITooLong, e)
	case *goerror.UnsupportedMediaType:
		return s.Ctx.JSON(http.StatusUnsupportedMediaType, e)
	case *goerror.RequestedRangeNotSatisfiable:
		return s.Ctx.JSON(http.StatusRequestedRangeNotSatisfiable, e)
	case *goerror.ExpectationFailed:
		return s.Ctx.JSON(http.StatusExpectationFailed, e)
	case *goerror.Teapot:
		return s.Ctx.JSON(http.StatusTeapot, e)
	case *goerror.MisdirectedRequest:
		return s.Ctx.JSON(http.StatusMisdirectedRequest, e)
	case *goerror.UnprocessableEntity:
		return s.Ctx.JSON(http.StatusUnprocessableEntity, e)
	case *goerror.Locked:
		return s.Ctx.JSON(http.StatusLocked, e)
	case *goerror.FailedDependency:
		return s.Ctx.JSON(http.StatusFailedDependency, e)
	case *goerror.TooEarly:
		return s.Ctx.JSON(http.StatusTooEarly, e)
	case *goerror.UpgradeRequired:
		return s.Ctx.JSON(http.StatusUpgradeRequired, e)
	case *goerror.PreconditionRequired:
		return s.Ctx.JSON(http.StatusPreconditionRequired, e)
	case *goerror.TooManyRequests:
		return s.Ctx.JSON(http.StatusTooManyRequests, e)
	case *goerror.RequestHeaderFieldsTooLarge:
		return s.Ctx.JSON(http.StatusRequestHeaderFieldsTooLarge, e)
	case *goerror.UnavailableForLegalReasons:
		return s.Ctx.JSON(http.StatusUnavailableForLegalReasons, e)

	// Server error
	case *goerror.InternalServerError:
		return s.Ctx.JSON(http.StatusInternalServerError, e)
	case *goerror.NotImplemented:
		return s.Ctx.JSON(http.StatusNotImplemented, e)
	case *goerror.BadGateway:
		return s.Ctx.JSON(http.StatusBadGateway, e)
	case *goerror.ServiceUnavailable:
		return s.Ctx.JSON(http.StatusServiceUnavailable, e)
	case *goerror.GatewayTimeout:
		return s.Ctx.JSON(http.StatusGatewayTimeout, e)
	case *goerror.HTTPVersionNotSupported:
		return s.Ctx.JSON(http.StatusHTTPVersionNotSupported, e)
	case *goerror.VariantAlsoNegotiates:
		return s.Ctx.JSON(http.StatusVariantAlsoNegotiates, e)
	case *goerror.InsufficientStorage:
		return s.Ctx.JSON(http.StatusInsufficientStorage, e)
	case *goerror.LoopDetected:
		return s.Ctx.JSON(http.StatusLoopDetected, e)
	case *goerror.NotExtended:
		return s.Ctx.JSON(http.StatusNotExtended, e)
	case *goerror.NetworkAuthenticationRequired:
		return s.Ctx.JSON(http.StatusNetworkAuthenticationRequired, e)

	// Other
	default:
		if s.Cus != nil {
			if s.I18n != nil && s.I18n.Enabled && s.I18n.Localize != nil {
				body, e1 := goerror.GetBody(err)
				if e1 == nil && body.Code != "" && body.Message == "" {
					if localize, e2 := s.I18n.Localize(s.Ctx, body.Code); e2 == nil {
						goerror.SetMessage(err, localize)
					}
				}
			}
			return (*s.Cus).Response(s.Ctx, err)
		}
		// Default response
		return s.Ctx.JSON(http.StatusBadRequest, goerror.NewBadRequest())
	}
}

func New(config ...*Config) Response {
	resp := &response{}
	if len(config) > 0 {
		cfg := config[0]
		resp.Cus = cfg.Custom
		resp.I18n = cfg.I18n
	}
	return resp
}
