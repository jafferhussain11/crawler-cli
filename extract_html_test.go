package main

import "testing"

func TestGetH1FromHTMLBasic(t *testing.T) {

	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "return empty string if no H1 tag",
			html:     "<html><body>Test Title</body></html>",
			expected: "",
		},
		{
			name:     "return Title if H1 tag present",
			html:     "<html><body><h1>Test Title</h1></body></html>",
			expected: "Test Title",
		},
		{
			name:     "return Empty if Input URL is empty",
			html:     "",
			expected: "",
		},
		{
			name:     "return Empty if no Html Tag",
			html:     "Test Title",
			expected: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.html)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {

	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "return empty string if no paragraph",
			html:     "<html><body>Test Title</body></html>",
			expected: "",
		},
		{
			name: "return Paragraph if paragraph present",
			html: `<html><body>
					<p>Outside paragraph.</p>
					<main>
					<p>Main paragraph.</p>
					<p>2nd Main paragraph.</p>
					</main>
					</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "return first Paragraph if main Not present",
			html: `<html><body>
					<p>Outside paragraph.</p>
					<p>Main paragraph.</p>
					<p>2nd Main paragraph.</p>
					</body></html>`,
			expected: "Outside paragraph.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.html)
			if actual != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}
