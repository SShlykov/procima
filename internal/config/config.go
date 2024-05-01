package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Logger Logger `mapstructure:"logger" yaml:"logger"`
}
type Logger struct {
	Level string `mapstructure:"level" yaml:"level"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	return config, nil
}
