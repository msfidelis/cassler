package parser

import "strings"

func ParseHost(url string) string {
	var result string
	result = strings.ToLower(url)
	result = strings.TrimPrefix(result, "https://")
	result = strings.TrimPrefix(result, "http://")
	result = strings.TrimPrefix(result, "ftp://")
	result = strings.TrimPrefix(result, "ws://")
	return result
}

func ParseDurationInDays(duration float64) int {
	floatDays := duration / 24
	parsedDays := int(floatDays)
	return parsedDays
}
