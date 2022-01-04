package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	out "github.com/srgio-es/tcvolumeutils/utils/output"
)

var (
	createMissingCmd = &cobra.Command{
		Use:   "createmissing",
		Short: "Creates missing files with empty content found using review_volumes command. Useful when trying to clean volumes.",
		Long: `This command creates empty files with the same exact name and locations as the missing found files.
        This is helpful when administering volumes with commands such as dataset_cleanup or purge_datasets as they
        will fail in the event of a missing reference.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkArgs()
		},
		Run: createMissing,
	}
)

func init() {
	rootCmd.AddCommand(createMissingCmd)

	createMissingCmd.Flags().StringVarP(&logFolder, "logs-folder", "f", "", "Specifies the location of the logs to be processed")
	createMissingCmd.MarkFlagDirname("logs-folder")

	createMissingCmd.Flags().StringVarP(&logFile, "log-file", "l", "", "Specifies the log file to be processed")
	createMissingCmd.MarkFlagFilename("log-file")
}

func createMissing(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Parent().Flags().GetBool("verbose")
	output = out.VerboseOutput{Verbose: verbose}

	switch {
	case logFolder != "":
		output.Printf("Log Folder: %s\n", logFolder)
		output.Println("")

		fmt.Println("You have entered a folder to be processed. This should not be stopped. Do you want to continue?\nPlease enter (y)es or (n)no")
		if askForConfirmation() {

		}

	case logFile != "":
		output.Printf("Log File: %s\n", logFile)
		output.Println("")

	}

}
func checkArgs() error {
	if logFile != "" && logFolder != "" {
		return errors.New("log-file and log-folder arguments are mutually exclusive. Please specify only one of them.")
	}

	return nil
}
