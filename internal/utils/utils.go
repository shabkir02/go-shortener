package utils

import "strings"

func GenerateURL(host string, path string) string {
	var sb strings.Builder
	sb.WriteString(host)
	sb.WriteString("/")
	sb.WriteString(path)
	s := sb.String()

	if strings.Contains(s, "https://") || strings.Contains(s, "http://") {
		return s
	} else {
		return "http://" + s
	}
}
