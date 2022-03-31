package mail

type Config struct {
	Domain string `koanf:"domain"`
	APIKEY string `koanf:"api_key"`
}
