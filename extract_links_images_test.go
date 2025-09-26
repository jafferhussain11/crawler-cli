package main

import (
	"net/url"
	"reflect"
	"testing"
)

// ensure :
// 1.In your tests, make sure that:
//relative URLs are converted to absolute URLs

//you find all the <a> tags in a body of HTML

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://blog.boot.dev"

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	tests := []struct {
		name     string
		htmlBody string
		baseURL  *url.URL
		expected []string
	}{
		{
			name:     "Test If result has all absolute URLs",
			htmlBody: `<html><body><a href="https://blog.boot.dev"><span>Boot.dev</span></a></body></html>`,
			baseURL:  baseURL,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name: "Test If result has all <a> Tags correctly collected",
			htmlBody: `<html><body>
  						 <a href="https://blog.boot.dev"><span>Boot.dev</span></a>
  						 <a href="/about">About</a>
						 </body></html>`,
			baseURL:  baseURL,
			expected: []string{"https://blog.boot.dev", "https://blog.boot.dev/about"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.htmlBody, tc.baseURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", tc.name, err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}

		})
	}
}
