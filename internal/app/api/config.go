package api

//General instance for API server of REST application

type Config struct {
	//Port
	BindArr     string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Config {
	return &Config{
		BindArr:     "8080",
		LoggerLevel: "debug",
	}
}
