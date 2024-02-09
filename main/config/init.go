package config

import (
	"GoProject/main/component"
	"GoProject/resource"
	"flag"
	"github.com/knadh/koanf"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

var conf = struct {
	DB  database `koanf:"database"`
	Sys system   `koanf:"system"`
	Ser server   `koanf:"server"`
}{}

// 导包时自动初始化conf变量
func init() {
	var dir, _ = os.Getwd()
	var path = *flag.String("config-path", filepath.Join(dir, "settings.toml"), "指定配置文件，若忽略使用默认配置")

	//  @Description: 匿名函数实现配置解析
	//  @param data 内嵌配置
	//  @return func(data []byte)
	func(data []byte, parser *koanf.Koanf) {
		//程序所在目录生成默认配置文件
		if _, err := os.Stat(path); err != nil {
			slog.Warn("即将生成配置文件，请修改后重启程序！！！")
			//内嵌配置序列化并生成settings.toml文件
			if err := os.WriteFile(path, data, 0664); err != nil {
				log.Fatalf("配置文件生成失败：%s", err.Error())
			}
			os.Exit(0)
		}
		//解析settings.toml文件并反序列化到结构体
		if err := parser.Unmarshal("", &conf); err != nil {
			log.Fatalf("配置文件解析错误 %s", err.Error())
		}
	}(resource.Settings, component.NewKoanf(path)) //执行程序
}

func GetDSN() string {
	return conf.DB.dsn()
}

func GetDBMS() string {
	return conf.DB.DBMS
}
