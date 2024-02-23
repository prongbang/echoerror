package echoerror

import (
	"github.com/labstack/echo/v4"
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
	case *Continue:
		return s.Ctx.JSON(http.StatusContinue, e)
	case *SwitchingProtocols:
		return s.Ctx.JSON(http.StatusSwitchingProtocols, e)
	case *Processing:
		return s.Ctx.JSON(http.StatusProcessing, e)
	case *EarlyHints:
		return s.Ctx.JSON(http.StatusEarlyHints, e)

	// Successful
	case *OK:
		return s.Ctx.JSON(http.StatusOK, e)
	case *Created:
		return s.Ctx.JSON(http.StatusCreated, e)
	case *Accepted:
		return s.Ctx.JSON(http.StatusAccepted, e)
	case *NonAuthoritativeInformation:
		return s.Ctx.JSON(http.StatusNonAuthoritativeInfo, e)
	case *NoContent:
		return s.Ctx.JSON(http.StatusNoContent, e)
	case *ResetContent:
		return s.Ctx.JSON(http.StatusResetContent, e)
	case *PartialContent:
		return s.Ctx.JSON(http.StatusPartialContent, e)
	case *MultiStatus:
		return s.Ctx.JSON(http.StatusMultiStatus, e)
	case *AlreadyReported:
		return s.Ctx.JSON(http.StatusAlreadyReported, e)
	case *IMUsed:
		return s.Ctx.JSON(http.StatusIMUsed, e)

	// Redirection
	case *MultipleChoices:
		return s.Ctx.JSON(http.StatusMultipleChoices, e)
	case *MovedPermanently:
		return s.Ctx.JSON(http.StatusMovedPermanently, e)
	case *Found:
		return s.Ctx.JSON(http.StatusFound, e)
	case *SeeOther:
		return s.Ctx.JSON(http.StatusSeeOther, e)
	case *NotModified:
		return s.Ctx.JSON(http.StatusNotModified, e)
	case *UseProxy:
		return s.Ctx.JSON(http.StatusUseProxy, e)
	case *TemporaryRedirect:
		return s.Ctx.JSON(http.StatusTemporaryRedirect, e)
	case *PermanentRedirect:
		return s.Ctx.JSON(http.StatusPermanentRedirect, e)

	// Client error
	case *BadRequest:
		return s.Ctx.JSON(http.StatusBadRequest, e)
	case *Unauthorized:
		return s.Ctx.JSON(http.StatusUnauthorized, e)
	case *PaymentRequired:
		return s.Ctx.JSON(http.StatusPaymentRequired, e)
	case *Forbidden:
		return s.Ctx.JSON(http.StatusForbidden, e)
	case *NotFound:
		return s.Ctx.JSON(http.StatusNotFound, e)
	case *MethodNotAllowed:
		return s.Ctx.JSON(http.StatusMethodNotAllowed, e)
	case *NotAcceptable:
		return s.Ctx.JSON(http.StatusNotAcceptable, e)
	case *ProxyAuthRequired:
		return s.Ctx.JSON(http.StatusProxyAuthRequired, e)
	case *RequestTimeout:
		return s.Ctx.JSON(http.StatusRequestTimeout, e)
	case *Conflict:
		return s.Ctx.JSON(http.StatusConflict, e)
	case *Gone:
		return s.Ctx.JSON(http.StatusGone, e)
	case *LengthRequired:
		return s.Ctx.JSON(http.StatusLengthRequired, e)
	case *PreconditionFailed:
		return s.Ctx.JSON(http.StatusPreconditionFailed, e)
	case *RequestEntityTooLarge:
		return s.Ctx.JSON(http.StatusRequestEntityTooLarge, e)
	case *RequestURITooLong:
		return s.Ctx.JSON(http.StatusRequestURITooLong, e)
	case *UnsupportedMediaType:
		return s.Ctx.JSON(http.StatusUnsupportedMediaType, e)
	case *RequestedRangeNotSatisfiable:
		return s.Ctx.JSON(http.StatusRequestedRangeNotSatisfiable, e)
	case *ExpectationFailed:
		return s.Ctx.JSON(http.StatusExpectationFailed, e)
	case *Teapot:
		return s.Ctx.JSON(http.StatusTeapot, e)
	case *MisdirectedRequest:
		return s.Ctx.JSON(http.StatusMisdirectedRequest, e)
	case *UnprocessableEntity:
		return s.Ctx.JSON(http.StatusUnprocessableEntity, e)
	case *Locked:
		return s.Ctx.JSON(http.StatusLocked, e)
	case *FailedDependency:
		return s.Ctx.JSON(http.StatusFailedDependency, e)
	case *TooEarly:
		return s.Ctx.JSON(http.StatusTooEarly, e)
	case *UpgradeRequired:
		return s.Ctx.JSON(http.StatusUpgradeRequired, e)
	case *PreconditionRequired:
		return s.Ctx.JSON(http.StatusPreconditionRequired, e)
	case *TooManyRequests:
		return s.Ctx.JSON(http.StatusTooManyRequests, e)
	case *RequestHeaderFieldsTooLarge:
		return s.Ctx.JSON(http.StatusRequestHeaderFieldsTooLarge, e)
	case *UnavailableForLegalReasons:
		return s.Ctx.JSON(http.StatusUnavailableForLegalReasons, e)

	// Server error
	case *InternalServerError:
		return s.Ctx.JSON(http.StatusInternalServerError, e)
	case *NotImplemented:
		return s.Ctx.JSON(http.StatusNotImplemented, e)
	case *BadGateway:
		return s.Ctx.JSON(http.StatusBadGateway, e)
	case *ServiceUnavailable:
		return s.Ctx.JSON(http.StatusServiceUnavailable, e)
	case *GatewayTimeout:
		return s.Ctx.JSON(http.StatusGatewayTimeout, e)
	case *HTTPVersionNotSupported:
		return s.Ctx.JSON(http.StatusHTTPVersionNotSupported, e)
	case *VariantAlsoNegotiates:
		return s.Ctx.JSON(http.StatusVariantAlsoNegotiates, e)
	case *InsufficientStorage:
		return s.Ctx.JSON(http.StatusInsufficientStorage, e)
	case *LoopDetected:
		return s.Ctx.JSON(http.StatusLoopDetected, e)
	case *NotExtended:
		return s.Ctx.JSON(http.StatusNotExtended, e)
	case *NetworkAuthenticationRequired:
		return s.Ctx.JSON(http.StatusNetworkAuthenticationRequired, e)

	// Other
	default:
		if s.Cus != nil {
			if s.I18n != nil && s.I18n.Enabled && s.I18n.Localize != nil {
				body, e1 := GetBody(err)
				if e1 == nil && body.Code != "" && body.Message == "" {
					if localize, e2 := s.I18n.Localize(s.Ctx, body.Code); e2 == nil {
						SetMessage(err, localize)
					}
				}
			}
			return (*s.Cus).Response(s.Ctx, err)
		}
		// Default response
		return s.Ctx.JSON(http.StatusBadRequest, NewBadRequest())
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
