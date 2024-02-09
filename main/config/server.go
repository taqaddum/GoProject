package config

type server struct {
	host  string `koanf:"Host"`
	port  uint16 `koanf:"Port"`
	salve bool   `koanf:"slave"`
}
