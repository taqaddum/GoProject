package config

type system struct {
	host string `koanf:"Host"`
	port uint16 `koanf:"Port"`
}
