package cmd

import (
	"github.com/apex/log"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	"github.com/prairir/imacry/cc-server/pkg/config"
	"github.com/prairir/imacry/cc-server/web"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "imacry-cc-server",
	Short: "command & control server for imacry",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: root,
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

	pflags.StringVar(&cfgFile, "config", "", "config file (default is /etc/imacry-cc-server/cc-server.yaml)")

	pflags.StringP("password", "p", "", "AES password (must be length of 16 + 8n where n is some number over 0)")
	viper.BindPFlag("password", pflags.Lookup("password"))

	pflags.Uint("port", 80, "network port to run on")
	viper.BindPFlag("port", pflags.Lookup("port"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in default directory with name "cc-server" (without extension).
		viper.AddConfigPath("/etc/imacry-cc-server")
		viper.SetConfigType("yaml")
		viper.SetConfigName("cc-server")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// read the config in
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Couldnt read config file: %s", err)
	}

	err := config.UnmarshalConfig()
	if err != nil {
		log.Fatalf("Couldnt unmarshal config: %s", err)
	}

	// making sure it is proper size
	if (len(config.Config.Password)-16)/8 < 1 {
		log.Fatalf("Password is length: %d .It requires length 16 + 8n where n is over 0", len(config.Config.Password))
	}
}

func root(cmd *cobra.Command, args []string) {
	web.Run()

}
