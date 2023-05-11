package utils

import "strings"

func GenerateURL(host string, path string) string {
	var sb strings.Builder
	sb.WriteString(host)
	sb.WriteString("/")
	sb.WriteString(path)

	return sb.String()
}
