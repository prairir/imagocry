package encryptfile_test

import (
	"testing"

	"github.com/prairir/imacry/pkg/config"
	"github.com/prairir/imacry/pkg/encryptfile"
)

var enc = encryptfile.EncryptFile{}

func BenchmarkEncryptFile(b *testing.B) {
	config.Config.Password = "123456789012345678901234"
	tests := []struct {
		name     string
		filePath string
	}{
		{"2MB", "/root/file00"},
		{"4MB", "/root/file01"},
		{"8MB", "/root/file02"},
		{"16MB", "/root/file03"},
		{"32MB", "/root/file04"},
		{"64MB", "/root/file05"},
		{"128MB", "/root/file06"},
		{"256MB", "/root/file07"},
		{"512MB", "/root/file08"},
		{"1024MB", "/root/file09"},
		{"2056MB", "/root/file10"},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				enc.Do(test.filePath)
			}
		})
	}
}
