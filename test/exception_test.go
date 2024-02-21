package test

import (
	"errors"
	"fmt"
	"testing"
)

func TestFileException(t *testing.T) {
	var ErrFileExist = errors.New("file exist")

	err := ErrFileExist
	res1 := errors.Is(err, ErrFileExist)
	res2 := errors.Is(errors.New(""), ErrFileExist)
	fmt.Println(res1)
	fmt.Println(res2)
}
