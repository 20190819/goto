package helpers

import "strings"

func StrTrimSpace(str string) string {
	return strings.Trim(str, " \n")
}
