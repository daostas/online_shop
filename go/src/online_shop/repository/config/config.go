package config

type Config struct {
	DbAddr string
}

func NewConfig() *Config {
	return &Config{
		DbAddr: "postgresql://daostas:St_031028As@185.102.75.212:5432/online_shop",
	}
}
