package echoerror_test

import "testing"

import (
	"github.com/prongbang/echoerror"
)

type customErr struct {
	echoerror.Body
}

func (u *customErr) Error() string {
	return u.Message
}

func newCustomErr() error {
	return &customErr{
		echoerror.Body{
			Code:    "XXX",
			Message: "YYY",
			Data:    nil,
		},
	}
}

func BenchmarkGetBody(b *testing.B) {
	err := newCustomErr()

	for i := 0; i < b.N; i++ {
		_, _ = echoerror.GetBody(err)
	}
}

func TestGetBody(t *testing.T) {
	err := newCustomErr()

	actual, _ := echoerror.GetBody(err)

	if actual.Code != "XXX" || actual.Message != "YYY" {
		t.Error("Error:", actual)
	}
}
