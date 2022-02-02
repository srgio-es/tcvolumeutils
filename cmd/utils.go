package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/srgio-es/tcvolumeutils/model"
	converter "github.com/srgio-es/tcvolumeutils/parser/line"
	"github.com/srgio-es/tcvolumeutils/utils"
)

func askForConfirmation() bool {
	var response string

	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("I'm sorry but I didn't get what you meant, please type (y)es or (n)o and then press enter:")
		return askForConfirmation()
	}
}

func processFile(location string, file string, volume string) []*model.MissingFile {

	var missingFiles []*model.MissingFile

	f, err := ioutil.ReadFile(location + string(os.PathSeparator) + file)
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

	// fmt.Println("")

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
