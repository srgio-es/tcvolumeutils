package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/srgio-es/tcvolumeutils/model"
	converter "github.com/srgio-es/tcvolumeutils/parser/line"
	"github.com/srgio-es/tcvolumeutils/reporter"
	"github.com/srgio-es/tcvolumeutils/utils"
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

	collection map[string][]model.MissingFile
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

	data := processLogs(toProcess)

	reporter := reporter.NewExcelReporter(reportFile)
	reporter.GenerateMissingFilesReport(data)

	fmt.Println("")
	fmt.Printf("Report file generated: %s\n", reportFile)

}

func processLogs(files []os.FileInfo) map[string][]*model.MissingFile {

	var result = make(map[string][]*model.MissingFile)

	for _, file := range files {
		output.Printf("Processed file: %s\n", file.Name())
		volume := file.Name()[:len(file.Name())-4]
		missingfiles := processFile(logFolder, file.Name(), volume)

		result[volume] = missingfiles
	}

	return result

}

func processFile(location string, file string, volume string) []*model.MissingFile {

	var missingFiles []*model.MissingFile

	f, err := ioutil.ReadFile(logFolder + string(os.PathSeparator) + file)
	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(f), "Error accessing volume") {

		lines := getLines(f)

		if len(lines) > 0 {
			fmt.Printf("Volume %s has %d missing files\n", file, len(lines))
		} else {
			output.Printf("Volume %s has no missing files\n", file)
		}

		for _, line := range lines {
			output.Printf("Missing: %s\n", line)
			p := converter.MissingFileParser{Line: line, Volume: volume}
			missingFile := p.ParseLine(line)

			missingFiles = append(missingFiles, &missingFile)
		}

	} else {
		output.Printf("File %s has errrors and cannot be processed", file)
	}

	fmt.Println("")

	return missingFiles

}

func getLines(raw []byte) []string {

	b := strings.Index(string(raw), "Files missing from the OS file that are referenced by Teamcenter:")

	if b > 0 {
		raw = raw[b:]
	}

	s := strings.Index(string(raw), "\n")
	e := strings.Index(string(raw), "--------------------------------------------------------------------------------")

	if s > 0 && e > 0 {
		raw = raw[s:e]
	}

	return utils.RemoveLineEndingsFromSlice(utils.RemoveEmptyFromSlice(strings.Split(string(raw), "\n")))
}
