package config

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type config struct {
	Password      string        `mapstructure:"password"`
	Port          string        `mapstructure:"port"`
	TriggerTime   time.Time     `mapstructure:"trigger-time"`
	TriggerLength time.Duration `mapstructure:"trigger-length"`
}

var Config config

// UnmarshalConfig: Unmarshals internal viper registery into a config struct
func UnmarshalConfig() error {
	//
	err := viper.Unmarshal(&Config, func(m *mapstructure.DecoderConfig) {
		m.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeHookFunc(time.Kitchen),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		)
	})
	if err != nil {
		return fmt.Errorf("Config Error: %w", err)
	}
	return nil
}
