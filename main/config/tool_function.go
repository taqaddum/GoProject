package config

import (
	"github.com/knadh/koanf"
	"strconv"
)

var conf = struct {
	db  database `koanf:"database"`
	ser server   `koanf:"server"`
}{}

func parse(parser *koanf.Koanf) error {
	err := parser.Unmarshal("", &conf)
	return err
}

func GetDSN() string {
	return conf.db.dsn()
}

func GetDBMS() string {
	return conf.db.DBMS
}

func GetSaveDir() string {
	return conf.ser.SaveDir
}

func GetServerAddr() string {
	return conf.ser.Host + ":" + strconv.Itoa(int(conf.ser.Port))
}
