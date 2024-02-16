package test

import (
	"GoProject/main/config"
	"fmt"
	"testing"
)

func TestGetDSN(t *testing.T) {
	fmt.Print(config.GetDSN())
}
