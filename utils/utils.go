package utils

import "strings"

func GetLines(input string) []string {
	return strings.Split(input, "\n\n")
}

