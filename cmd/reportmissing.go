package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/srgio-es/tcvolumeutils/reporter"
	out "github.com/srgio-es/tcvolumeutils/utils/output"
)

// reportmissingCmd represents the reportmissing command
var (
	reportmissingCmd = &cobra.Command{
		Use:   "reportmissing",
		Short: "Extracts missing OS files from review_volumes logs",
		Long: `This command extracts missing OS files info from the logs generated by the review_volumes command.

The information is processed and writed into a XLXS file to improve readiness.
If the XLSX file specified contains data, the file is updated appending the new values.`,
		Run: reportMissing,
	}

	output out.VerboseOutput
)

func init() {
	rootCmd.AddCommand(reportmissingCmd)

	reportmissingCmd.Flags().StringVarP(&logFolder, "logs-folder", "f", "", "Specifies the location of the logs to be processed")
	reportmissingCmd.MarkFlagDirname("logs-folder")
	reportmissingCmd.MarkFlagRequired("logs-folder")

	reportmissingCmd.Flags().StringVarP(&reportFile, "report", "r", "volumes-report.xlsx", `Specifies the path to the XLSX file to populate with the results.`)
	reportmissingCmd.MarkFlagFilename("report")
}

func reportMissing(cmd *cobra.Command, args []string) {
	verbose, _ := cmd.Parent().Flags().GetBool("verbose")
	output = out.VerboseOutput{Verbose: verbose}

	fmt.Println("reportmissing process begings")
	fmt.Println("")

	output.Printf("Log Folder: %s\n", logFolder)
	output.Println("")

	allFiles, err := ioutil.ReadDir(logFolder)
	if err != nil {
		fmt.Println("ERROR: The specified logs directory does not exist")
		fmt.Println("")
		cmd.Usage()
		os.Exit(1)
	}

	var toProcess []os.FileInfo
	for _, file := range allFiles {
		if !file.IsDir() && strings.HasSuffix(file.Name(), "txt") {
			toProcess = append(toProcess, file)
		}
	}

	output.Printf("%d review_volumes log files marked to process", len(toProcess))

	data := processLogs(logFolder, toProcess)

	reporter := reporter.NewExcelReporter(reportFile)
	reporter.GenerateMissingFilesReport(data)

	fmt.Println("")
	fmt.Printf("Report file generated: %s\n", reportFile)

}
