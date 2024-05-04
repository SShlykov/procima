package config

import (
	"github.com/spf13/viper"
)

type ProcessorConfig struct {
	LargestSideLimit int `mapstructure:"largest_side_limit" yaml:"largest_side_limit"`
}

func LoadProcessorConfig(configPath string) (*ServerConfig, error) {
	viper.SetConfigName("processor")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &ServerConfig{}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
