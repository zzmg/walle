package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gitlab.wallstcn.com/wscnbackend/ivankastd"
)

// Response API body
type Response struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func SuccessResponse(ctx echo.Context, data interface{}) error {
	var buffer bytes.Buffer
	m := json.NewEncoder(&buffer)
	if err := m.Encode(data); err != nil {
		return ErrorResponse(ctx, errors.New("Illegal JSON"))
	}
	return ctx.JSON(http.StatusOK, Payload(buffer.Bytes()))
}

func ErrorResponse(ctx echo.Context, err error) error {
	if err != nil {
		return ErrorResponseWithMessage(ctx, err.Error())
	}
	return ErrorResponseWithMessage(ctx, "")
}

func ErrorResponseWithMessage(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusOK, Payload([]byte("{}"), "50000", message))
}

func Payload(data json.RawMessage, fields ...string) *Response {
	res := &Response{
		Code:    ivankastd.ErrOK.Code,
		Message: ivankastd.ErrOK.Message,
		Data:    data,
	}

	length := len(fields)
	if length >= 1 {
		singleCode, _ := strconv.Atoi(fields[0])
		res.Code = int(singleCode)
	}
	if length >= 2 {
		res.Message = fields[1]
	}

	return res
}
