package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppName     string `mapstructure:"app_name" yaml:"app_name"`
	Host        string `mapstructure:"host" yaml:"host"`
	Connections int    `mapstructure:"connections" yaml:"connections"`
	Logger      Logger `mapstructure:"logger" yaml:"logger"`
}
type Logger struct {
	Level string `mapstructure:"level" yaml:"level"`
	Mode  string `mapstructure:"mode" yaml:"mode"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
