package echoerror

import (
	"net/http"
)

type UseProxy struct {
	Body
}

// Error implements error.
func (u *UseProxy) Error() string {
	return u.Message
}

func NewUseProxy() error {
	return &UseProxy{
		Body: Body{
			Code:    CodeUseProxy,
			Message: http.StatusText(http.StatusUseProxy),
		},
	}
}

type Unauthorized struct {
	Body
}

func (e *Unauthorized) Error() string {
	return e.Message
}

func NewUnauthorized() error {
	return &Unauthorized{
		Body: Body{
			Code:    CodeUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
		},
	}
}

type TemporaryRedirect struct {
	Body
}

// Error implements error.
func (t *TemporaryRedirect) Error() string {
	return t.Message
}

func NewTemporaryRedirect() error {
	return &TemporaryRedirect{
		Body: Body{
			Code:    CodeTemporaryRedirect,
			Message: http.StatusText(http.StatusTemporaryRedirect),
		},
	}
}

type NotFound struct {
	Body
}

// Error implements error.
func (t *NotFound) Error() string {
	return t.Message
}

func NewNotFound() error {
	return &NotFound{
		Body: Body{
			Code:    CodeNotFound,
			Message: http.StatusText(http.StatusNotFound),
		},
	}
}

type SwitchingProtocols struct {
	Body
}

// Error implements error.
func (s *SwitchingProtocols) Error() string {
	return s.Message
}

func NewSwitchingProtocols() error {
	return &SwitchingProtocols{
		Body: Body{
			Code:    CodeSwitchingProtocols,
			Message: http.StatusText(http.StatusSwitchingProtocols),
		},
	}
}

type SeeOther struct {
	Body
}

// Error implements error.
func (s *SeeOther) Error() string {
	return s.Message
}

func NewSeeOther() error {
	return &SeeOther{
		Body: Body{
			Code:    CodeSeeOther,
			Message: http.StatusText(http.StatusSeeOther),
		},
	}
}

type ResetContent struct {
	Body
}

// Error implements error.
func (r *ResetContent) Error() string {
	return r.Message
}

func NewResetContent() error {
	return &ResetContent{
		Body: Body{
			Code:    CodeResetContent,
			Message: http.StatusText(http.StatusResetContent),
		},
	}
}

type RequestTimeout struct {
	Body
}

// Error implements error.
func (r *RequestTimeout) Error() string {
	return r.Message
}

func NewRequestTimeout() error {
	return &RequestTimeout{
		Body: Body{
			Code:    CodeRequestTimeout,
			Message: http.StatusText(http.StatusRequestTimeout),
		},
	}
}

type ProxyAuthRequired struct {
	Body
}

// Error implements error.
func (p *ProxyAuthRequired) Error() string {
	return p.Message
}

func NewProxyAuthRequired() error {
	return &ProxyAuthRequired{
		Body: Body{
			Code:    CodeProxyAuthRequired,
			Message: http.StatusText(http.StatusProxyAuthRequired),
		},
	}
}

type Processing struct {
	Body
}

// Error implements error.
func (p *Processing) Error() string {
	return p.Message
}

func NewProcessing() error {
	return &Processing{
		Body: Body{
			Code:    CodeProcessing,
			Message: http.StatusText(http.StatusProcessing),
		},
	}
}

type PermanentRedirect struct {
	Body
}

// Error implements error.
func (p *PermanentRedirect) Error() string {
	return p.Message
}

func NewPermanentRedirect() error {
	return &PermanentRedirect{
		Body: Body{
			Code:    CodePermanentRedirect,
			Message: http.StatusText(http.StatusPermanentRedirect),
		},
	}
}

type PaymentRequired struct {
	Body
}

// Error implements error.
func (p *PaymentRequired) Error() string {
	return p.Message
}

func NewPaymentRequired() error {
	return &PaymentRequired{
		Body: Body{
			Code:    CodePaymentRequired,
			Message: http.StatusText(http.StatusPaymentRequired),
		},
	}
}

type PartialContent struct {
	Body
}

// Error implements error.
func (p *PartialContent) Error() string {
	return p.Message
}

func NewPartialContent() error {
	return &PartialContent{
		Body: Body{
			Code:    CodePartialContent,
			Message: http.StatusText(http.StatusPartialContent),
		},
	}
}

type OK struct {
	Body
}

func (e *OK) Error() string {
	return e.Message
}

func NewOK(data any) error {
	return &OK{
		Body: Body{
			Code:    CodeOK,
			Data:    data,
			Message: http.StatusText(http.StatusOK),
		},
	}
}

type NotModified struct {
	Body
}

// Error implements error.
func (n *NotModified) Error() string {
	return n.Message
}

func NewNotModified() error {
	return &NotModified{
		Body: Body{
			Code:    CodeNotModified,
			Message: http.StatusText(http.StatusNotModified),
		},
	}
}

type NotAcceptable struct {
	Body
}

// Error implements error.
func (n *NotAcceptable) Error() string {
	return n.Message
}

func NewNotAcceptable() error {
	return &NotAcceptable{
		Body: Body{
			Code:    CodeNotAcceptable,
			Message: http.StatusText(http.StatusNotAcceptable),
		},
	}
}

type NonAuthoritativeInformation struct {
	Body
}

// Error implements error.
func (n *NonAuthoritativeInformation) Error() string {
	return n.Message
}

func NewNonAuthoritativeInformation() error {
	return &NonAuthoritativeInformation{
		Body: Body{
			Code:    CodeNonAuthoritativeInformation,
			Message: http.StatusText(http.StatusNonAuthoritativeInfo),
		},
	}
}

type NoContent struct {
	Body
}

// Error implements error.
func (n *NoContent) Error() string {
	return n.Message
}

func NewNoContent() error {
	return &NoContent{
		Body: Body{
			Code:    CodeNoContent,
			Message: http.StatusText(http.StatusNoContent),
		},
	}
}

type MultipleChoices struct {
	Body
}

// Error implements error.
func (m *MultipleChoices) Error() string {
	return m.Message
}

func NewMultipleChoices() error {
	return &MultipleChoices{
		Body: Body{
			Code:    CodeMultipleChoices,
			Message: http.StatusText(http.StatusMultipleChoices),
		},
	}
}

type MultiStatus struct {
	Body
}

// Error implements error.
func (m *MultiStatus) Error() string {
	return m.Message
}

func NewMultiStatus() error {
	return &MultiStatus{
		Body: Body{
			Code:    CodeMultiStatus,
			Message: http.StatusText(http.StatusMultiStatus),
		},
	}
}

type MovedPermanently struct {
	Body
}

// Error implements error.
func (m *MovedPermanently) Error() string {
	return m.Message
}

func NewMovedPermanently() error {
	return &MovedPermanently{
		Body: Body{
			Code:    CodeMovedPermanently,
			Message: http.StatusText(http.StatusMovedPermanently),
		},
	}
}

type MethodNotAllowed struct {
	Body
}

// Error implements error.
func (m *MethodNotAllowed) Error() string {
	return m.Message
}

func NewMethodNotAllowed() error {
	return &MethodNotAllowed{
		Body: Body{
			Code:    CodeMethodNotAllowed,
			Message: http.StatusText(http.StatusMethodNotAllowed),
		},
	}
}

type IMUsed struct {
	Body
}

// Error implements error.
func (i *IMUsed) Error() string {
	return i.Message
}

func NewIMUsed() error {
	return &IMUsed{
		Body: Body{
			Code:    CodeIMUsed,
			Message: http.StatusText(http.StatusIMUsed),
		},
	}
}

type Found struct {
	Body
}

// Error implements error.
func (f *Found) Error() string {
	return f.Message
}

func NewFound() error {
	return &Found{
		Body: Body{
			Code:    CodeFound,
			Message: http.StatusText(http.StatusFound),
		},
	}
}

type Forbidden struct {
	Body
}

// Error implements error.
func (f *Forbidden) Error() string {
	return f.Message
}

func NewForbidden() error {
	return &Forbidden{
		Body: Body{
			Code:    CodeForbidden,
			Message: http.StatusText(http.StatusForbidden),
		},
	}
}

type EarlyHints struct {
	Body
}

// Error implements error.
func (e *EarlyHints) Error() string {
	return e.Message
}

func NewEarlyHints() error {
	return &EarlyHints{
		Body: Body{
			Code:    CodeEarlyHints,
			Message: http.StatusText(http.StatusEarlyHints),
		},
	}
}

type Created struct {
	Body
}

// Error implements error.
func (c *Created) Error() string {
	return c.Message
}

func NewCreated(data any) error {
	return &Created{
		Body: Body{
			Code:    CodeCreated,
			Message: http.StatusText(http.StatusCreated),
			Data:    data,
		},
	}
}

type Continue struct {
	Body
}

// Error implements error.
func (c *Continue) Error() string {
	return c.Message
}

func NewContinue() error {
	return &Continue{
		Body: Body{
			Code:    CodeContinue,
			Message: http.StatusText(http.StatusContinue),
		},
	}
}

type Conflict struct {
	Body
}

// Error implements error.
func (c *Conflict) Error() string {
	return c.Message
}

func NewConflict() error {
	return &Conflict{
		Body: Body{
			Code:    CodeConflict,
			Message: http.StatusText(http.StatusConflict),
		},
	}
}

type BadRequest struct {
	Body
}

// Error implements error.
func (b *BadRequest) Error() string {
	return b.Message
}

func NewBadRequest(message ...string) error {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = http.StatusText(http.StatusBadRequest)
	}
	return &BadRequest{
		Body: Body{
			Code:    CodeBadRequest,
			Message: msg,
			Data:    nil,
		},
	}
}

type AlreadyReported struct {
	Body
}

// Error implements error.
func (a *AlreadyReported) Error() string {
	return a.Message
}

func NewAlreadyReported() error {
	return &AlreadyReported{
		Body: Body{
			Code:    CodeAlreadyReported,
			Message: http.StatusText(http.StatusAlreadyReported),
		},
	}
}

type Accepted struct {
	Body
}

// Error implements error.
func (a *Accepted) Error() string {
	return a.Message
}

func NewAccepted() error {
	return &Accepted{
		Body: Body{
			Code:    CodeAccepted,
			Message: http.StatusText(http.StatusAccepted),
		},
	}
}

type Gone struct {
	Body
}

// Error implements error.
func (a *Gone) Error() string {
	return a.Message
}

func NewGone() error {
	return &Gone{
		Body: Body{
			Code:    CodeGone,
			Message: http.StatusText(http.StatusGone),
		},
	}
}

type LengthRequired struct {
	Body
}

// Error implements error.
func (a *LengthRequired) Error() string {
	return a.Message
}

func NewLengthRequired() error {
	return &LengthRequired{
		Body: Body{
			Code:    CodeLengthRequired,
			Message: http.StatusText(http.StatusLengthRequired),
		},
	}
}

type PreconditionFailed struct {
	Body
}

// Error implements error.
func (a *PreconditionFailed) Error() string {
	return a.Message
}

func NewPreconditionFailed() error {
	return &PreconditionFailed{
		Body: Body{
			Code:    CodePreconditionFailed,
			Message: http.StatusText(http.StatusPreconditionFailed),
		},
	}
}

type RequestEntityTooLarge struct {
	Body
}

// Error implements error.
func (a *RequestEntityTooLarge) Error() string {
	return a.Message
}

func NewRequestEntityTooLarge() error {
	return &RequestEntityTooLarge{
		Body: Body{
			Code:    CodeRequestEntityTooLarge,
			Message: http.StatusText(http.StatusRequestEntityTooLarge),
		},
	}
}

type RequestURITooLong struct {
	Body
}

// Error implements error.
func (a *RequestURITooLong) Error() string {
	return a.Message
}

func NewRequestURITooLong() error {
	return &RequestURITooLong{
		Body: Body{
			Code:    CodeRequestURITooLong,
			Message: http.StatusText(http.StatusRequestURITooLong),
		},
	}
}

type UnsupportedMediaType struct {
	Body
}

// Error implements error.
func (a *UnsupportedMediaType) Error() string {
	return a.Message
}

func NewUnsupportedMediaType() error {
	return &UnsupportedMediaType{
		Body: Body{
			Code:    CodeUnsupportedMediaType,
			Message: http.StatusText(http.StatusUnsupportedMediaType),
		},
	}
}

type RequestedRangeNotSatisfiable struct {
	Body
}

// Error implements error.
func (a *RequestedRangeNotSatisfiable) Error() string {
	return a.Message
}

func NewRequestedRangeNotSatisfiable() error {
	return &RequestedRangeNotSatisfiable{
		Body: Body{
			Code:    CodeRequestedRangeNotSatisfiable,
			Message: http.StatusText(http.StatusRequestedRangeNotSatisfiable),
		},
	}
}

type ExpectationFailed struct {
	Body
}

// Error implements error.
func (a *ExpectationFailed) Error() string {
	return a.Message
}

func NewExpectationFailed() error {
	return &ExpectationFailed{
		Body: Body{
			Code:    CodeExpectationFailed,
			Message: http.StatusText(http.StatusExpectationFailed),
		},
	}
}

type Teapot struct {
	Body
}

// Error implements error.
func (a *Teapot) Error() string {
	return a.Message
}

func NewTeapot() error {
	return &Teapot{
		Body: Body{
			Code:    CodeTeapot,
			Message: http.StatusText(http.StatusTeapot),
		},
	}
}

type MisdirectedRequest struct {
	Body
}

// Error implements error.
func (a *MisdirectedRequest) Error() string {
	return a.Message
}

func NewMisdirectedRequest() error {
	return &MisdirectedRequest{
		Body: Body{
			Code:    CodeMisdirectedRequest,
			Message: http.StatusText(http.StatusMisdirectedRequest),
		},
	}
}

type UnprocessableEntity struct {
	Body
}

// Error implements error.
func (a *UnprocessableEntity) Error() string {
	return a.Message
}

func NewUnprocessableEntity() error {
	return &UnprocessableEntity{
		Body: Body{
			Code:    CodeUnprocessableEntity,
			Message: http.StatusText(http.StatusUnprocessableEntity),
		},
	}
}

type Locked struct {
	Body
}

// Error implements error.
func (a *Locked) Error() string {
	return a.Message
}

func NewLocked() error {
	return &Locked{
		Body: Body{
			Code:    CodeLocked,
			Message: http.StatusText(http.StatusLocked),
		},
	}
}

type FailedDependency struct {
	Body
}

// Error implements error.
func (a *FailedDependency) Error() string {
	return a.Message
}

func NewFailedDependency() error {
	return &FailedDependency{
		Body: Body{
			Code:    CodeFailedDependency,
			Message: http.StatusText(http.StatusFailedDependency),
		},
	}
}

type TooEarly struct {
	Body
}

// Error implements error.
func (a *TooEarly) Error() string {
	return a.Message
}

func NewTooEarly() error {
	return &TooEarly{
		Body: Body{
			Code:    CodeTooEarly,
			Message: http.StatusText(http.StatusTooEarly),
		},
	}
}

type UpgradeRequired struct {
	Body
}

// Error implements error.
func (a *UpgradeRequired) Error() string {
	return a.Message
}

func NewUpgradeRequired() error {
	return &UpgradeRequired{
		Body: Body{
			Code:    CodeUpgradeRequired,
			Message: http.StatusText(http.StatusUpgradeRequired),
		},
	}
}

type PreconditionRequired struct {
	Body
}

// Error implements error.
func (a *PreconditionRequired) Error() string {
	return a.Message
}

func NewPreconditionRequired() error {
	return &PreconditionRequired{
		Body: Body{
			Code:    CodePreconditionRequired,
			Message: http.StatusText(http.StatusPreconditionRequired),
		},
	}
}

type TooManyRequests struct {
	Body
}

// Error implements error.
func (a *TooManyRequests) Error() string {
	return a.Message
}

func NewTooManyRequests() error {
	return &TooManyRequests{
		Body: Body{
			Code:    CodeTooManyRequests,
			Message: http.StatusText(http.StatusTooManyRequests),
		},
	}
}

type RequestHeaderFieldsTooLarge struct {
	Body
}

// Error implements error.
func (a *RequestHeaderFieldsTooLarge) Error() string {
	return a.Message
}

func NewRequestHeaderFieldsTooLarge() error {
	return &RequestHeaderFieldsTooLarge{
		Body: Body{
			Code:    CodeRequestHeaderFieldsTooLarge,
			Message: http.StatusText(http.StatusRequestHeaderFieldsTooLarge),
		},
	}
}

type UnavailableForLegalReasons struct {
	Body
}

// Error implements error.
func (a *UnavailableForLegalReasons) Error() string {
	return a.Message
}

func NewUnavailableForLegalReasons() error {
	return &UnavailableForLegalReasons{
		Body: Body{
			Code:    CodeUnavailableForLegalReasons,
			Message: http.StatusText(http.StatusUnavailableForLegalReasons),
		},
	}
}

type InternalServerError struct {
	Body
}

// Error implements error.
func (a *InternalServerError) Error() string {
	return a.Message
}

func NewInternalServerError() error {
	return &InternalServerError{
		Body: Body{
			Code:    CodeInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		},
	}
}

type NotImplemented struct {
	Body
}

// Error implements error.
func (a *NotImplemented) Error() string {
	return a.Message
}

func NewNotImplemented() error {
	return &NotImplemented{
		Body: Body{
			Code:    CodeNotImplemented,
			Message: http.StatusText(http.StatusNotImplemented),
		},
	}
}

type BadGateway struct {
	Body
}

// Error implements error.
func (a *BadGateway) Error() string {
	return a.Message
}

func NewBadGateway() error {
	return &BadGateway{
		Body: Body{
			Code:    CodeBadGateway,
			Message: http.StatusText(http.StatusBadGateway),
		},
	}
}

type ServiceUnavailable struct {
	Body
}

// Error implements error.
func (a *ServiceUnavailable) Error() string {
	return a.Message
}

func NewServiceUnavailable() error {
	return &ServiceUnavailable{
		Body: Body{
			Code:    CodeServiceUnavailable,
			Message: http.StatusText(http.StatusServiceUnavailable),
		},
	}
}

type GatewayTimeout struct {
	Body
}

// Error implements error.
func (a *GatewayTimeout) Error() string {
	return a.Message
}

func NewGatewayTimeout() error {
	return &GatewayTimeout{
		Body: Body{
			Code:    CodeGatewayTimeout,
			Message: http.StatusText(http.StatusGatewayTimeout),
		},
	}
}

type HTTPVersionNotSupported struct {
	Body
}

// Error implements error.
func (a *HTTPVersionNotSupported) Error() string {
	return a.Message
}

func NewHTTPVersionNotSupported() error {
	return &HTTPVersionNotSupported{
		Body: Body{
			Code:    CodeHTTPVersionNotSupported,
			Message: http.StatusText(http.StatusHTTPVersionNotSupported),
		},
	}
}

type VariantAlsoNegotiates struct {
	Body
}

// Error implements error.
func (a *VariantAlsoNegotiates) Error() string {
	return a.Message
}

func NewVariantAlsoNegotiates() error {
	return &VariantAlsoNegotiates{
		Body: Body{
			Code:    CodeVariantAlsoNegotiates,
			Message: http.StatusText(http.StatusVariantAlsoNegotiates),
		},
	}
}

type InsufficientStorage struct {
	Body
}

// Error implements error.
func (a *InsufficientStorage) Error() string {
	return a.Message
}

func NewInsufficientStorage() error {
	return &InsufficientStorage{
		Body: Body{
			Code:    CodeInsufficientStorage,
			Message: http.StatusText(http.StatusInsufficientStorage),
		},
	}
}

type LoopDetected struct {
	Body
}

// Error implements error.
func (a *LoopDetected) Error() string {
	return a.Message
}

func NewLoopDetected() error {
	return &LoopDetected{
		Body: Body{
			Code:    CodeLoopDetected,
			Message: http.StatusText(http.StatusLoopDetected),
		},
	}
}

type NotExtended struct {
	Body
}

// Error implements error.
func (a *NotExtended) Error() string {
	return a.Message
}

func NewNotExtended() error {
	return &NotExtended{
		Body: Body{
			Code:    CodeNotExtended,
			Message: http.StatusText(http.StatusNotExtended),
		},
	}
}

type NetworkAuthenticationRequired struct {
	Body
}

// Error implements error.
func (a *NetworkAuthenticationRequired) Error() string {
	return a.Message
}

func NewNetworkAuthenticationRequired() error {
	return &NetworkAuthenticationRequired{
		Body: Body{
			Code:    CodeNetworkAuthenticationRequired,
			Message: http.StatusText(http.StatusNetworkAuthenticationRequired),
		},
	}
}
