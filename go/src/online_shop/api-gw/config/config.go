package config

import "github.com/spf13/viper"

type Config struct {
	Port         string `mapstructure:"PORT"`
	ClientSvcUrl string `mapstructure:"CLIENT_SVC_URL"`
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
