package cmd

import (
	"github.com/spf13/cobra"
)

var (
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

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "increase verbosity for commands")
}
