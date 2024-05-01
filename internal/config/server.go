package config

import (
	"github.com/spf13/viper"
	"time"
)

type ServerConfig struct {
	Addr string `mapstructure:"addr" yaml:"addr"`

	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout" yaml:"read_header_timeout"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout" yaml:"read_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout" yaml:"write_timeout"`
	IddleTimeout      time.Duration `mapstructure:"iddle_timeout" yaml:"iddle_timeout"`

	Timeout time.Duration `mapstructure:"timeout" yaml:"timeout"`

	CorsEnabled    bool `mapstructure:"cors_enabled" yaml:"cors_enabled"`
	SwaggerEnabled bool `mapstructure:"swagger_enabled" yaml:"swagger_enabled"`
}

func LoadServerConfig(configPath string) (*ServerConfig, error) {
	viper.SetConfigName("server")
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
