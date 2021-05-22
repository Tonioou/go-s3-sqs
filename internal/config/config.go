package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AccessKey string
	SecretKey string
}

func NewConfig() *Config {
	config, err := loadConfig("./resources")
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
		//panic("Cannot found config file")
	}
	return &config
}

func loadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
		return Config{}, err
	}

	config := Config{
		AccessKey: viper.GetString("AWS.AccessKey"),
		SecretKey: viper.GetString("AWS.SecretKey"),
	}
	return config, err
}
