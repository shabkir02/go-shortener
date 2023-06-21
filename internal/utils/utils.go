package utils

import (
	"compress/gzip"
	"io"
	"net/http"
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

func HandleReadBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	// переменная reader будет равна r.Body или *gzip.Reader
	var reader io.Reader

	if r.Header.Get(`Content-Encoding`) == `gzip` {
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return []byte{}, err
		}
		reader = gz
		defer gz.Close()
	} else {
		reader = r.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return []byte{}, err
	}

	return body, nil
}
