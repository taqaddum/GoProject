package config

var conf = struct {
	DB  database `koanf:"database"`
	Sys system   `koanf:"system"`
	Ser server   `koanf:"server"`
}{}

func GetDSN() string {
	return conf.DB.dsn()
}

func GetDBMS() string {
	return conf.DB.DBMS
}
