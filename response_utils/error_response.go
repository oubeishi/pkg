package response_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type ResponseError struct {
	Code       int
	Message    string
	StatusCode int
	ERR        error
}

type MysqlError struct {
	mysql.MySQLError
}

func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

func UnWrapResponse(err error) *ResponseError {
	if v, ok := err.(*ResponseError); ok {
		return v
	}
	return nil
}

func ConvertToResponseError(a mysql.MySQLError) *ResponseError {
	return &ResponseError{
		Code:       500,
		Message:    a.Message,
		StatusCode: 200,
		ERR:        &MysqlError{a},
	}
}
func WrapResponse(err error, code, statusCode int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:       code,
		Message:    fmt.Sprintf(msg, args...),
		ERR:        err,
		StatusCode: statusCode,
	}
	return res
}

func Wrap400Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 400, 400, msg, args...)
}

func Wrap500Response(err error, msg string, args ...interface{}) error {
	return WrapResponse(err, 500, 500, msg, args...)
}

func New400Response(msg string, args ...interface{}) error {
	return NewErrorRes(400, 400, msg, args...)
}

func New500Response(msg string, args ...interface{}) error {
	return NewErrorRes(500, 500, msg, args...)
}
