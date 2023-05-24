package utils

import (
	"strings"
)

func GenerateURL(URL string, path string) string {
	var sb strings.Builder
	sb.WriteString(URL)
	sb.WriteString("/")
	sb.WriteString(path)
	s := sb.String()

	return s
}

func ValidateURL(URL string) string {
	if strings.Contains(URL, "https://") || strings.Contains(URL, "http://") {
		return URL
	}

	return "http://" + URL
}
