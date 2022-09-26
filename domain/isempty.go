package domain

import "strings"

func CheckEmpty(s string) bool {
	return strings.ReplaceAll(s, "\\n", "") == ""
}
