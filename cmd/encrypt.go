package cmd

import (
	"fmt"

	"github.com/apex/log"
	"github.com/prairir/imacry/pkg/config"
	"github.com/prairir/imacry/pkg/encryptfile"
	"github.com/prairir/imacry/pkg/walk"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt files",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: RunEncrypt,
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RunEncrypt(cmd *cobra.Command, args []string) {
	// set the state first thing
	config.Config.State = config.EncryptState
	fmt.Printf("config: %#v\n", config.Config)

	fmt.Println("encrypt run")
	ef := encryptfile.EncryptFile{}
	err := walk.Walk(config.Config.Base, ef)
	if err != nil {
		log.Log.Fatalf("Fatal error: %s", err)
	}
}
