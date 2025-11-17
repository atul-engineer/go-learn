package config

type Config struct {
	DBHost string `envconfig:"DbHost"`
}