package test

import (
	"GoProject/main/config"
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseFile(t *testing.T) {
	fmt.Println(config.GetDSN())
	var jsonStruct = struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
		Sex  string `json:"sex"`
		Desc string `json:"desc"`
		Addr string `json:"addr"`
		Tel  string `json:"tel"`
		Mail string `json:"mail"`
		Url  string `json:"url"`
		Id   string `json:"id"`
		Pwd  string `json:"pwd"`
	}{
		Name: "张三",
		Sex:  "男",
	}
	res, _ := json.Marshal(jsonStruct)
	fmt.Println(string(res))
}
