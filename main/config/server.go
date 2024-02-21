package config

type server struct {
	Host    string `koanf:"host"`
	Port    uint16 `koanf:"port"`
	SaveDir string `koanf:"savedir"`
}
