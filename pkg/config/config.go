package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type State int

const (
	EncryptState State = iota
	WaitState
	DecryptState
)

// global config struct
type config struct {
	State    State
	Address  string `mapstructure:"cc-address"`
	Password string `mapstructure:"password"`
	Base     string `mapstructure:"base"`
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
