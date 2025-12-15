package config

import (
	"ikbs/lib/basic"

	"github.com/spf13/viper"
)

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigType("yaml")

	rootPath, err := basic.GetRootPath()
	if err != nil {
		return nil, err
	}
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
