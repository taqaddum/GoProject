package test

import (
	"GoProject/main/view"
	"encoding/json"
	"fmt"
	"testing"
)

func TestResponse(t *testing.T) {
	success, _ := json.Marshal(view.Success("success"))
	fail, _ := json.Marshal(view.Failed())
	fmt.Println(string(success), string(fail))
}
