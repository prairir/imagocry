package cmd

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/spf13/viper"

	"github.com/prairir/imacry/pkg/config"
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
	viper.BindPFlag("cc-address", pflag.Lookup("cc-address"))

	pflags.StringP("password", "p", "", "encryption/decryption password.")
	viper.BindPFlag("password", pflag.Lookup("password"))

	pflags.StringP("base", "b", "", "base path to start encrypting")
	viper.BindPFlag("base", pflag.Lookup("base"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if base is empty, get the home directory
	if viper.GetString("base") == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.Set("base", home)
	}

	err := config.UnmarshalConfig()
	if err != nil {
		log.Fatalf("Couldnt unmarshal config: %s", err)
	}
}

func RunImacry(cmd *cobra.Command, args []string) {
	fmt.Println("hello")
	fmt.Printf("config: %#v\n", config.Config)
}
