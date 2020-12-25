package parser

import "strings"

// ParseHost is a way to normalize URL's
func ParseHost(url string) string {
	var result string
	result = strings.ToLower(url)
	result = strings.TrimPrefix(result, "https://")
	result = strings.TrimPrefix(result, "http://")
	result = strings.TrimPrefix(result, "ftp://")
	result = strings.TrimPrefix(result, "ws://")
	return result
}

// ParseDurationInDays is a way to determine how many days to certificate expires
func ParseDurationInDays(duration float64) int {
	floatDays := duration / 24
	parsedDays := int(floatDays)
	return parsedDays
}
