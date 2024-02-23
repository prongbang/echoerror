package echoerror

import (
	"fmt"
	"reflect"
)

type Body struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GetBody(err error) (Body, error) {
	errType := reflect.TypeOf(err)
	if errType.Kind() != reflect.Ptr {
		return Body{}, fmt.Errorf("error is not a pointer")
	}

	errElm := reflect.ValueOf(err).Elem()
	if errElm.Kind() != reflect.Struct {
		return Body{}, fmt.Errorf("error is not a struct")
	}

	bodyField := errElm.FieldByName("Body")
	if !bodyField.IsValid() {
		return Body{}, fmt.Errorf("body field not found in the struct")
	}

	body := Body{}
	codeField := bodyField.FieldByName("Code")
	if codeField.IsValid() {
		body.Code = fmt.Sprintf("%v", codeField.Interface())
	}

	msgField := bodyField.FieldByName("Message")
	if msgField.IsValid() {
		body.Message = fmt.Sprintf("%v", msgField.Interface())
	}

	return body, nil
}

func SetMessage(err interface{}, message string) {
	errType := reflect.TypeOf(err)
	if errType.Kind() != reflect.Ptr {
		return
	}

	errElm := reflect.ValueOf(err).Elem()
	if errElm.Kind() != reflect.Struct {
		return
	}

	bodyField := errElm.FieldByName("Body")
	if !bodyField.IsValid() {
		return
	}

	messageField := bodyField.FieldByName("Message")
	if !messageField.IsValid() || !messageField.CanSet() {
		return
	}
	messageField.SetString(message)
	return
}
