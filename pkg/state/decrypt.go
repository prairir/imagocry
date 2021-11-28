package state

import (
	"fmt"

	"github.com/prairir/imacry/pkg/config"
	"github.com/prairir/imacry/pkg/decryptfile"
	"github.com/prairir/imacry/pkg/walk"
)

// state.Decrypt: decrypts file system starting at config.Config.Base
// params: the next state
// returns: error
func Decrypt(nextState config.State) error {
	df := decryptfile.DecryptFile{}
	// Decrypt from the base file path walking through all files on the system
	err := walk.Walk(config.Config.Base, df)
	if err != nil {
		return fmt.Errorf("state.Decrypt error: %w", err)
	}
	config.Config.State = nextState
	return nil
}
