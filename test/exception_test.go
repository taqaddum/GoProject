package test

import (
	"GoProject/main/enum/exception"
	"errors"
	"fmt"
	"testing"
)

func TestFileException(t *testing.T) {
	err := exception.ErrFileExist
	res1 := errors.Is(err, exception.ErrFileExist)
	res2 := errors.Is(errors.New(""), exception.ErrFileExist)
	fmt.Println(res1)
	fmt.Println(res2)
}
