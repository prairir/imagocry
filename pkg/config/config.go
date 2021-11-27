package config

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
)

type State int

const (
	InitState State = iota
	EncryptState
	WaitState
	DecryptState
	ExitState
)

// global config struct
type config struct {
	// current state the machine is on
	State State

	// command and control server address
	Address string `mapstructure:"cc-address"`

	// encrypt/decrypt password
	Password string `mapstructure:"password"`

	// base path to start encrypt
	Base string `mapstructure:"base"`

	// connection with cc server
	Conn *websocket.Conn

	// decrypt signal
	// the wait state requires this
	Signal chan struct{}

	// if heartbeat has an error
	// propegate it into a channel
	HBError chan error
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
