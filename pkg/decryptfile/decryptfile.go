package decryptfile

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"

	"github.com/prairir/imacry/pkg/config"
)

type DecryptFile struct{}

func (df DecryptFile) Do(filePath string) error {
	// get the password from config
	password := config.Config.Password
	// Check to make sure the password is not empty.
	if password == "" {
		return fmt.Errorf("decryptfile.Do error: No password in config.Config")
	}

	// Create a cipherblock from the password
	cipherBlock, err := aes.NewCipher([]byte(password))
	if err != nil {
		return fmt.Errorf("decryptfile.Do error: %w", err)
	}

	// Open the input file
	infile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("decryptfile.Do error: %w", err)
	}

	// Get the permissions of the input file.
	inPerms, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("decryptfile.Do error: %w", err)
	}

	// Open the output file, this is a new file which will overwrite the original path later.
	outfile, err := os.OpenFile(filePath+".dec", os.O_WRONLY|os.O_CREATE, inPerms.Mode())
	if err != nil {
		return fmt.Errorf("decryptfile.Do error: %w", err)
	}

	outfileBuf := bufio.NewWriterSize(outfile, 10240)

	// Get the initialization vector from the beginning of the input file.
	iv := make([]byte, cipherBlock.BlockSize())
	if _, err := infile.Read(iv); err != nil {
		return fmt.Errorf("decryptfile.Do error: %w", err)
	}

	buf := make([]byte, 1024)
	stream := cipher.NewCTR(cipherBlock, iv)

	for {
		n, err := infile.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			_, err := outfileBuf.Write(buf[:n])
			if err != nil {
				return fmt.Errorf("decryptfile.Do error: %w", err)
			}

		}

		// EOF has been reached
		if err == io.EOF {
			err := outfileBuf.Flush()
			if err != nil {
				return fmt.Errorf("decryptfile.Do error: %w", err)
			}
			break
		}

		// Any error other than EOF
		if err != nil {
			return fmt.Errorf("decryptfile.Do error: %w", err)
		}
	}

	// Close input file.
	err = infile.Close()
	if err != nil {
		return fmt.Errorf("decryptfile.Do error: %w", err)
	}

	// Close output file.
	err = outfile.Close()
	if err != nil {
		return fmt.Errorf("decryptfile.Do error: %w", err)
	}

	// Move the file into the original filePath.
	err = os.Rename(filePath+".dec", filePath)
	if err != nil {
		return fmt.Errorf("decryptfile.Do error: %w", err)
	}

	return nil
}
