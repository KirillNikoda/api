package api

//General instance for API server of REST application

type Config struct {
	//Port
	BindArr string `toml:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindArr: "8080",
	}
}