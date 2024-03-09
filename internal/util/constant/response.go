package constant

import "fmt"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

func (r Response) Error() string {
	return r.Message
}

func Success(msg string, data any) error {
	return Response{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}

func Created(msg string) error {
	return Response{
		Code:    201,
		Message: msg,
	}
}

func NotFound(msg string) error {
	return Response{
		Code:    404,
		Message: msg,
	}
}

func BadRequest(msg string) error {
	return Response{
		Code:    400,
		Message: msg,
	}
}

func InternalServerError(msg string) error {
	return Response{
		Code:    500,
		Message: msg,
	}
}

func Deleted(msg string, data any) error {
	return Response{
		Code:    204,
		Message: fmt.Sprintf("%s %s", msg, data),
	}
}
