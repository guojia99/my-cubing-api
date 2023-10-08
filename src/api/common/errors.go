package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Msg        string `json:"Msg"`
	Code       int    `json:"Code"`
	StatusCode int    `json:"StatusCode"`
}

func Error(ctx *gin.Context, statusCode int, code int, err any) {
	var msg string
	switch data := err.(type) {
	case string:
		msg = data
	case error:
		msg = data.Error()
	default:
		msg = fmt.Sprintf("%+v", msg)
	}

	ctx.JSON(statusCode, ErrorMessage{
		Msg:        msg,
		Code:       code,
		StatusCode: statusCode,
	})
}
