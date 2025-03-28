package main

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	if !strings.Contains(inputURL, "https://") && !strings.Contains(inputURL, "http://") {
		return "", errors.New("invalid url: no http protocol")
	}

	URL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	normURL, _ := url.JoinPath(URL.Host, URL.Path)

	return strings.ToLower(strings.TrimRight(normURL, "/")), nil
}
