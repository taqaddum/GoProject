package component

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"log"
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
