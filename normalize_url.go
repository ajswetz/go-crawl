package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {

	u, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	pathTrimmed := strings.TrimSuffix(u.Path, "/")

	normalizedURL := u.Host + pathTrimmed

	return normalizedURL, nil

}
