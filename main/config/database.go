package config

import (
	"fmt"
	"net/url"
)

type database struct {
	DBMS   string            `koanf:"dbms"`
	Host   string            `koanf:"host"`
	Port   uint16            `koanf:"port"`
	User   string            `koanf:"user"`
	DBName string            `koanf:"dbname"`
	Passwd string            `koanf:"passwd"`
	Params map[string]string `koanf:"params"`
}

// GetDSN
//
//	@Description: 获取数据库连接地址
//	@receiver db 数据库实例
//	@return string 连接地址
func (db database) dsn() string {
	var params = make(url.Values)
	for k, v := range db.Params {
		params.Add(k, v)
	}

	dsn := url.URL{
		Scheme:   db.DBMS,
		User:     url.UserPassword(db.User, db.Passwd),
		Host:     fmt.Sprintf("%s:%d", db.Host, db.Port),
		Path:     db.DBName,
		RawQuery: params.Encode(),
	}
	return dsn.String()
}
