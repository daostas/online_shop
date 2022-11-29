package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	DBUrl         string `mapstructure:"DB_URL"`
	JWTSecretKey  string `mapstructure:"JWT_SECRET_KEY"`
	TokenDuration int32  `mapstructure:"TOKEN_DURATION"`
	BcryptCost    int32  `mapstructure:"BCRYPT_COST"`
}

func LoadConfig(envPath string) (config Config, err error) {
	viper.AddConfigPath(envPath + "/env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
