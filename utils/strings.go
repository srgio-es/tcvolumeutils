package utils

import "strings"

func RemoveStringLineEndings(st string) string {
	st = strings.TrimRight(st, " \r")

	return st
}