package walk_test

import (
	"testing"

	"github.com/prairir/imacry/pkg/config"
	"github.com/prairir/imacry/pkg/decryptfile"
	"github.com/prairir/imacry/pkg/encryptfile"
	"github.com/prairir/imacry/pkg/walk"
)

func BenchmarkWalk(b *testing.B) {
	config.Config.Password = "123456789012345678901234"

	tests := []struct {
		name string
		fun  walk.FileAction
	}{
		{"encrypt", encryptfile.EncryptFile{}},
		{"decrypt", decryptfile.DecryptFile{}},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				walk.Walk("/root/", test.fun)
			}
		})
	}
}
