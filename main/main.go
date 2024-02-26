package main

import (
	"GoProject/main/component"
	"GoProject/main/mapper"
	"flag"
	"fmt"
)

var Version string

func main() {
	var autoMigrate = *flag.Bool("migrate", true, "数据结构迁移，默认为true")
	flag.Parse()

	if autoMigrate {
		mapper.Migration(component.NewPostgres())
	}

	fmt.Println("当前版本号:", Version)

	app := NewApp()
	app.Start(":8088")
}
