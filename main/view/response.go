package view

import (
	"GoProject/main/enum/httpstatus"
)

type Response[R interface{ ~uint | ~int }] struct {
	Code    R      `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func StatusWith[T httpstatus.Generic](status T, data ...any) Response[T] {
	var temp any
	if len(data) > 0 {
		temp = data[0]
	}

	return Response[T]{
		Code:    status,
		Message: status.String(),
		Data:    temp,
	}
}

func Success(data ...any) Response[httpstatus.Common] {
	return StatusWith(httpstatus.Ok, data...)
}

func Fail() Response[httpstatus.Common] {
	return StatusWith(httpstatus.Error)
}
