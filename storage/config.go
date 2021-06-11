package storage

type Config struct {
	//DB connection string
	DatabaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
