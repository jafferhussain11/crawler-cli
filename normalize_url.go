package main

import (
	"net/url"
	"strings"
)

func normalizeURL(urlString string) (string, error) {

	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	path := parsedURL.Host + parsedURL.Path
	slashPos := strings.LastIndex(path, "/")
	lastIndex := len(path) - 1

	//we can use strings.TrimSuffix also
	if slashPos == lastIndex {
		path = path[:slashPos]
	}

	return strings.ToLower(path), nil
}
