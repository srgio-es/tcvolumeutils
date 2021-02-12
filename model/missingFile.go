package model

import "time"


type MissingFile struct {
	Volume			string
	DatasetName 	string
	UID				string
	Site			string
	Version			int64
	ModifiedDate	time.Time
	FileLocation	string
}