package config

import (
	"ikbs/lib/basic"
	"log"

	"github.com/spf13/viper"
)

var config *Config

func Init() {
	var err error
	config, err = initConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigType("yaml")

	rootPath := basic.GetRootPath()

	v.SetConfigFile(rootPath + "/config/config.yml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadConfig() *Config {
	return config
}
