package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/srgio-es/tcvolumeutils/model"
	"github.com/srgio-es/tcvolumeutils/utils"
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
		Run: createMissingRoot,
	}
)

func init() {
	rootCmd.AddCommand(createMissingCmd)

	createMissingCmd.Flags().StringVarP(&logFolder, "logs-folder", "f", "", "Specifies the location of the logs to be processed")
	createMissingCmd.MarkFlagDirname("logs-folder")

	createMissingCmd.Flags().StringVarP(&logFile, "log-file", "l", "", "Specifies the log file to be processed")
	createMissingCmd.MarkFlagFilename("log-file")
}

func createMissingRoot(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Parent().Flags().GetBool("verbose")
	output = out.VerboseOutput{Verbose: verbose}

	var toProcess []os.FileInfo

	switch {
	case logFolder != "":
		output.Printf("Log Folder: %s\n", logFolder)
		output.Println("")

		fmt.Println("You have entered a folder to be processed. This should not be stopped. Do you want to continue?\nPlease enter (y)es or (n)no")
		if askForConfirmation() {
			output.Println("")

			allFiles, err := ioutil.ReadDir(logFolder)
			if err != nil {
				fmt.Println("ERROR: The specified logs directory does not exist")
				fmt.Println("")
				cmd.Usage()
				os.Exit(1)
			}

			for _, file := range allFiles {
				if !file.IsDir() && strings.HasSuffix(file.Name(), "txt") {
					toProcess = append(toProcess, file)
				}
			}

			output.Printf("%d review_volumes log files marked to process", len(toProcess))

		} else {
			os.Exit(2)
		}

	case logFile != "":
		output.Printf("Log File: %s\n", logFile)
		output.Println("")

		fp, _ := filepath.Abs(logFile)
		s := strings.Split(fp, string(os.PathSeparator))

		f, err := os.Stat(logFile)
		if err != nil {
			fmt.Println("ERROR: The log file does not exist")
			fmt.Printf("%s\n", err.Error())
			cmd.Usage()
			os.Exit(1)
		}

		toProcess = append(toProcess, f)
		logFolder = strings.Join(s[:len(s)-1], string(os.PathSeparator))
	}

	data := processLogs(logFolder, toProcess)
	createMissingFiles(data)

}

func createMissingFiles(data map[string][]*model.MissingFile) {

	for volumeName := range data {
		fmt.Printf("\nCreating files for volume %s\n", volumeName)

		d := data[volumeName]

		for _, mf := range d {
			if utils.CheckDirExists(mf.FileLocation) {
				if !utils.CheckFileAlreadyExists(mf.FileLocation) {
					d := []byte("missing")
					err := ioutil.WriteFile(mf.FileLocation, d, 0755)
					if err != nil {
						fmt.Printf("%s\n", err.Error())
						os.Exit(3)
					}
					fmt.Printf("File %s for Dataset '%s' created succesfuly\n", mf.FileLocation, mf.DatasetName)
				} else {
					fmt.Printf("File %s exists in the destination folder. Skipping.\n", mf.FileLocation)
				}
			} else {
				fmt.Printf("Destination folder for %s doesn't exist. You must create it prior execution of this utility. Skipping.\n", mf.FileLocation)
			}
		}
	}
}

func checkArgs() error {
	if logFile != "" && logFolder != "" {
		return errors.New("log-file and log-folder arguments are mutually exclusive. Please specify only one of them.\n")
	} else if logFile == "" && logFolder == "" {
		return errors.New("Please specify at least one argument\n")
	}

	return nil
}
