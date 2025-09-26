package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(inputBody string, inputURL string) PageData {

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Printf("Error parsing URL: %s\n", inputURL)
		return PageData{}
	}

	H1 := getH1FromHTML(inputBody)

	firstParagraph := getFirstParagraphFromHTML(inputBody)

	outgoingLinks, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		fmt.Printf("Error fetching outgoing links: %q\n", err)
		return PageData{}
	}

	imageURLs, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		fmt.Printf("Error fetching image URLs: %q\n", err)
		return PageData{}
	}

	return PageData{
		URL:            inputURL,
		H1:             H1,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageURLs,
	}
}
