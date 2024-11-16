package utils

import (
	"strings"
)

func SplitAndRemoveWhitespace(stringToSplit *string) []string {
	stringSlice := strings.Split(*stringToSplit, ",")

	for key, value := range stringSlice {
		stringSlice[key] = strings.TrimSpace(value)
	}

	return stringSlice
}
