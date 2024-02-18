package config

import (
	"github.com/knadh/koanf"
)

var conf = struct {
	DB  database `koanf:"database"`
	Sys system   `koanf:"system"`
	Ser server   `koanf:"server"`
}{}

func parse(parser *koanf.Koanf) error {
	err := parser.Unmarshal("", &conf)
	return err
}

func GetDSN() string {
	return conf.DB.dsn()
}

func GetDBMS() string {
	return conf.DB.DBMS
}
