package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string

	logFolder  string
	logFile    string
	reportFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tcvolumeutils",
	Short: "Teamcenter Volume management utils",
	Long:  "Teamcenter Volume management utils is a set of tools that will help Teamcenter administrator to perform some TC volumes management tedious tasks in an easier way.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tcvolumeutils.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "increase verbosity for commands")
}

// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".tcvolumeutils" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".tcvolumeutils")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}
// }
