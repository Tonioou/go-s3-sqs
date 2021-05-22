package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AwsConfig AwsConfiguration
}

type AwsConfiguration struct {
	AccessKey string
	SecretKey string
	Region    string
}

func NewConfig() {
	_, err := loadConfig("./resources")
	if err != nil {
		fmt.Errorf("Cannot found config file", err)
		panic("Cannot found config file")
	}
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

	config := GetConfiguration()
	return *config, nil
}

func GetConfiguration() *Config {
	return &Config{
		AwsConfig: AwsConfiguration{
			AccessKey: viper.GetString("AWS.AccessKey"),
			SecretKey: viper.GetString("AWS.SecretKey"),
			Region:    viper.GetString("AWS.Region"),
		},
	}
}
