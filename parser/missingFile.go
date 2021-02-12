package parser

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/srgio-es/tcvolumeutils/model"
)

type MissingFileParser struct {
	Line 	string
	Volume 	string
}

func (p *MissingFileParser) ParseLine(line string) model.MissingFile {
	missingFile := model.MissingFile{
		DatasetName: p.parseDatasetName(),
		UID: p.parseDatasetUID(),
		Site: p.parseSite(),
		Version: p.parseVersion(),
		ModifiedDate: p.parseModifiedDate(),
		FileLocation: p.parseFileLocation(),
		Volume: p.Volume,
	}
	return missingFile
}

func (p *MissingFileParser) parseDatasetName() string {
	re := regexp.MustCompile("'.*?'")
	str := re.FindString(p.Line)
	return str[1:len(str)-1]
}

func (p *MissingFileParser) parseDatasetUID() string {
	re := regexp.MustCompile("<.*?>")
	str := re.FindString(p.Line)
	return str[1:len(str)-1]
}

func (p *MissingFileParser) parseSite() string {
	re := regexp.MustCompile("site:<.*?>")
	str := re.FindString(p.Line)
	return str[6:len(str)-1]
}

func (p *MissingFileParser) parseVersion() int64 {
	re := regexp.MustCompile("version:([0-9]{1,})")
	str := re.FindString(p.Line)
	v, err := strconv.ParseInt(str[8:], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return  v
}

func (p *MissingFileParser) parseModifiedDate() time.Time {
	re := regexp.MustCompile("lmd:.{1,}\\)")
	str := re.FindString(p.Line)

	layout := "02/01/06 15:04:05"

	date, err := time.Parse(layout, str[4:len(str)-1])

	if err!=nil {
		log.Fatal(err)
	}

	return date
}

func (p *MissingFileParser) parseFileLocation() string {
	return p.Line[strings.Index(p.Line, "references ")+11:]
}
