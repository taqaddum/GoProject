package view

import (
	"GoProject/main/enum/opstatus"
)

type Response[R interface{ ~uint | ~int }] struct {
	Code    R      `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func CodeWith[T opstatus.Generic](code T, data ...any) Response[T] {
	var temp any
	if len(data) > 0 {
		temp = data[0]
	}

	return Response[T]{
		Code:    code,
		Message: code.String(),
		Data:    temp,
	}
}

func Success(data ...any) Response[opstatus.Common] {
	return CodeWith(opstatus.Ok, data...)
}

func Failed(data ...any) Response[opstatus.Common] {
	return CodeWith(opstatus.Error, data...)
}
