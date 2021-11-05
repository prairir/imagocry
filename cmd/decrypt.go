package cmd

import (
	"fmt"

	"github.com/apex/log"
	"github.com/prairir/imacry/pkg/config"
	"github.com/prairir/imacry/pkg/decryptfile"
	"github.com/prairir/imacry/pkg/walk"
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
	fmt.Println("decrypt run")
	// initialize the decrypt file struct to be passed to the file walker
	df := decryptfile.DecryptFile{}
	// Decrypt from the base file path walking through all files on the system
	err := walk.Walk(config.Config.Base, df)
	if err != nil {
		log.Log.Fatalf("Fatal error: %s", err)
	}
}
