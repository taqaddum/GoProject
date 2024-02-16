package view

import (
	"GoProject/main/enum/opstatus"
)

type Response[R interface{ ~uint | ~int }] struct {
	Code    R      `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func StatusWith[T opstatus.Generic](status T, data ...any) Response[T] {
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

func Success(data ...any) Response[opstatus.Operate] {
	return StatusWith(opstatus.Ok, data...)
}

func Fail(data ...any) Response[opstatus.Operate] {
	return StatusWith(opstatus.Error, data...)
}
