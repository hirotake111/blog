package config

type Config struct {
	Host string
	Port string
}

func NewConfig(host, port string) *Config {
	return &Config{
		Host: host,
		Port: port,
	}
}
