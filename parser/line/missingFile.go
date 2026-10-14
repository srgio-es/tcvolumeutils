package line

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/srgio-es/tcvolumeutils/model"
)

var (
	datasetNameRe  = regexp.MustCompile("'(.*?)'")
	datasetUIDRe   = regexp.MustCompile("<+([^<>]+?)>+")
	siteRe         = regexp.MustCompile("site:<(.*?)>")
	versionRe      = regexp.MustCompile("version:([0-9]+)")
	modifiedDateRe = regexp.MustCompile(`lmd:(.+?)\)`)
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

// findGroup returns the first capture group of re in the line, or "" when
// there is no match. Some lines (e.g. "unreferenced ImanFile <<UID>> references ...")
// lack the dataset name, site, version and lmd fields.
func (p *MissingFileParser) findGroup(re *regexp.Regexp) string {
	m := re.FindStringSubmatch(p.Line)
	if m == nil {
		return ""
	}
	return m[1]
}

func (p *MissingFileParser) parseDatasetName() string {
	return p.findGroup(datasetNameRe)
}

func (p *MissingFileParser) parseDatasetUID() string {
	return p.findGroup(datasetUIDRe)
}

func (p *MissingFileParser) parseSite() string {
	return p.findGroup(siteRe)
}

func (p *MissingFileParser) parseVersion() int64 {
	str := p.findGroup(versionRe)
	if str == "" {
		return 0
	}

	v, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return  v
}

func (p *MissingFileParser) parseModifiedDate() time.Time {
	str := p.findGroup(modifiedDateRe)
	if str == "" {
		return time.Time{}
	}

	layout := "06/01/02 15:04:05"

	date, err := time.Parse(layout, str)

	if err!=nil {
		log.Fatal(err)
	}

	return date
}

func (p *MissingFileParser) parseFileLocation() string {
	return p.Line[strings.Index(p.Line, "references ")+11:]
}
