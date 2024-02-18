package config

import (
	"GoProject/resource"
	"flag"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
)

var onceKoan sync.Once
var koan *koanf.Koanf

func NewKoanf(path string) *koanf.Koanf {
	onceKoan.Do(func() {
		koan = koanf.New(".")
	})

	err := koan.Load(file.Provider(path), toml.Parser())
	if err != nil {
		log.Fatalf("koanf初始化错误:%s", err.Error())
	}
	return koan
}

// 导包时自动初始化conf变量
func init() {
	var dir, _ = os.Getwd()
	var path = *flag.String("config-path", filepath.Join(dir, "settings.toml"), "指定配置文件，若忽略使用默认配置")

	//  @Description: 匿名函数实现配置解析
	//  @param data 内嵌配置
	//  @return func(data []byte)
	func(data []byte) {
		//程序所在目录生成默认配置文件
		if _, err := os.Stat(path); err != nil {
			slog.Warn("即将生成配置文件，请修改后重启程序！！！")
			//内嵌配置序列化并生成settings.toml文件
			if err := os.WriteFile(path, data, 0664); err != nil {
				log.Fatalf("配置文件生成失败：%s", err.Error())
			}
			os.Exit(0)
		}
		//解析settings.toml文件
		if err := parse(NewKoanf(path)); err != nil {
			log.Fatalf("配置文件解析错误 %s", err.Error())
		}
	}(resource.Settings) //执行程序
}
