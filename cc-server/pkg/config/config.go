package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	Password string `mapstructure:"password"`
	Port     uint   `mapstructure:"port"`
}

var Config config

// UnmarshalConfig: Unmarshals internal viper registery into a config struct
func UnmarshalConfig() error {
	err := viper.Unmarshal(&Config)
	if err != nil {
		return fmt.Errorf("Config Error: %w", err)
	}
	return nil
}
