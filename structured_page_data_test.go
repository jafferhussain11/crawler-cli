package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
    </body></html>`

	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  PageData
	}{
		{
			name:      "Test successful PageData extraction",
			inputURL:  inputURL,
			inputBody: inputBody,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
				ImageURLs:      []string{"https://blog.boot.dev/image1.jpg"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := extractPageData(inputBody, inputURL)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("actual = %v\nwant %v", actual, test.expected)
			}
		})
	}
}

//structured_page_data_test.go:41:
//actual = {https://blog.boot.dev Test Title This is the first paragraph. [https://blog.boot.dev/link1] [https://blog.boot.dev/link1]}
//want {https://blog.boot.dev Test Title This is the first paragraph. [https://blog.boot.dev/link1] [https://blog.boot.dev/image1.jpg]}
