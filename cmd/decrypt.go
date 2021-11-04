package cmd

import (
	"crypto/aes"
	"fmt"

	"github.com/apex/log"
	"github.com/prairir.imacry/pkg/decryptfile"
	"github.com/prairir/imacry/pkg/config"
	"github.com/prairir/imacry/pkg/walk"

	//"github.com/prairir/imacry/cmd/walk"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt files",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: RunDecrypt,
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RunDecrypt(cmd *cobra.Command, args []string) {
	// set the state to decrypt
	config.Config.State = config.DecryptState
	fmt.Printf("config: %#v\n", config.Config)
	df := decryptfile.DecryptFile{}
	err := walk.Walk(config.Config.Base, df)
	if err != nil {
		log.Log.Fatalf("Fatal error: %s", err)
	}
	/*
		fileData, err := os.ReadFile("/home/testFile.txt")
		if err != nil {
			fmt.Errorf("Decrypt Error: %w", err)
		}
		//fmt.Printf("%s\n", string(fileData))
		decryptedData := DecryptAES([]byte("thisis32bitlongpassphraseimusing"), fileData)
		os.WriteFile("/home/testFile.txt", decryptedData, 0644)
		//fmt.Printf("password: %#v\n", config.Config.Password)
		fmt.Println("decrypt complete")
	*/
}

func DecryptAES(key []byte, cipherText []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plainText := make([]byte, len(cipherText))
	cipher.Decrypt(plainText, cipherText)

	s := string(plainText[:])
	fmt.Println("DECRYPTED:", s)
	return plainText
}
