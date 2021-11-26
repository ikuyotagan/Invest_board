package apiserver

import "os"

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
	CustomDatabseURl string `env:"CUSTOM_DATABASE_URL" env-default:""`
	CustomBindAddr string `env:"CUSTOM_BIND_ADDRESS" env-default:""`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		CustomDatabseURl: os.Getenv("CUSTOM_DATABASE_URL"),
		CustomBindAddr: os.Getenv("CUSTOM_BIND_ADDRESS"),
	}
}
