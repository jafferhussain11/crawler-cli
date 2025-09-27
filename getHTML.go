package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return "", err
	}
	req.Header.Set("User-Agent", "CliCrawler/1.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return "", err
	}
	//dont forget to close
	defer resp.Body.Close()

	//this must be a string matcher as the value in the key is a large string that contains more than just text/html
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		fmt.Printf("%s is not HTML\n", rawURL)
		return "", err

	}
	if resp.StatusCode >= 400 {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return "", err
	}

	htmlString, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return "", err
	}

	return string(htmlString), nil
}
