package cmd

import (
	"errors"
	"fmt"
	"os"

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

	err := checkArgs()
	if err != nil {
		fmt.Println(err.Error() + "\n\n")
		cmd.Usage()
		os.Exit(1)
	}

}
func checkArgs() error {
	if logFile != "" && logFolder != "" {
		return errors.New("log-file and log-folder arguments are mutually exclusive. Please specify only one of them")
	}

	return nil
}
