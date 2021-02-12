package parser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

	var testString = "Dataset 'X391 GWM Part by Part v2 20-06-19.xlsx' <FtIAQL3SIDhs9A>(site:<local> version:54 lmd:19/06/20 15:46:08) references \\\\PLM_DATA\\VOLS\\West_Midlands\\west_midlands_north__5b068187\\x391_g_exc_lt609us830v1m.xlsx"

	var p = MissingFileParser{Line: testString, Volume: "SF-TOL"}

func  Test_parseDatasetName(t *testing.T) {
	res := p.parseDatasetName()
	assert.Equal(t, "X391 GWM Part by Part v2 20-06-19.xlsx", res)
}

func Test_parseDatasetUID(t *testing.T) {
	res := p.parseDatasetUID()
	assert.Equal(t, "FtIAQL3SIDhs9A", res)
}

func Test_parseSite(t *testing.T) {
	res := p.parseSite()
	assert.Equal(t, "local", res)
}

func Test_ParseVersion(t *testing.T) {
	res := p.parseVersion()
	assert.EqualValues(t, 54, res)
}

func Test_ParseModifiedDate(t *testing.T) {
	testDate := time.Date(2020, 6, 19, 15, 46, 8, 0, time.UTC)
	res := p.parseModifiedDate()
	assert.Equal(t, testDate, res)
}

func Test_ParseFileLocation(t *testing.T) {
	test := "\\\\PLM_DATA\\VOLS\\West_Midlands\\west_midlands_north__5b068187\\x391_g_exc_lt609us830v1m.xlsx"
	res := p.parseFileLocation()
	assert.Equal(t, test, res)

	p.Line = "Dataset 'M8E2-5G054-A-INS-01/G013' <m9DARl9UIDhs9A>(site:<local> version:1 lmd:20/11/20 13:00:39) references D:\\PLM_DATA\\VOLS\\SF-UK-AYC\\jatkinson_uk_5b1e6255\\m8e_cat_c640gd88tq0e6.catpart"
	test = "D:\\PLM_DATA\\VOLS\\SF-UK-AYC\\jatkinson_uk_5b1e6255\\m8e_cat_c640gd88tq0e6.catpart"

	res = p.parseFileLocation()
	assert.Equal(t, test, res)
}

func Test_Volume(t *testing.T) {
	assert.Equal(t, "SF-TOL", p.Volume)
}