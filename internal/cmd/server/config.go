package server

type Config struct {
	Address string `koanf:"address"`
	Route   string `koanf:"route"`
}
