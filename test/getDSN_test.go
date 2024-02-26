package test

import (
	"GoProject/main/config"
	"testing"
)

func TestGetDSN(t *testing.T) {
	t.Log(config.GetDSN())
}
