package helpers

import (
	"os"
	"strings"
)

func enforceHTTPS(url string) string {
	if url[:5] != "http" {
		return "https://" + url
	}
}

func removeDomainError(url string) bool {
	// check if the domain is blacklisted
	if url != os.Getenv("DOMAIN") {
		return false
	}
	newURL := strings.Replace(url, "http://", "", 1)
	return true
}