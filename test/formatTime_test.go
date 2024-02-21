package test

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	fmt.Println(time.Now().Format("20060102"))
}
