package main

import (
	"GoProject/main/component"
	"GoProject/main/mapper"
	"flag"
	"fmt"
)

var Version string

func init() {
	var isMigrate = *flag.Bool("migrate", true, "模型迁移，默认为true")
	if isMigrate {
		err := mapper.Migration(component.NewPostgres())
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	fmt.Println("当前版本号:", Version)
	app := NewApp()
	app.Start(":8088")
}
