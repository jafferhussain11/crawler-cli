package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

	result := []string{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		if val, exists := s.Attr("href"); exists {
			parsedURL, err := url.Parse(val)
			if err != nil {
				fmt.Printf("Couldn't parse URL: %v", err)
			}

			absoluteURL := baseURL.ResolveReference(parsedURL)
			absoluteURLString := absoluteURL.String()
			result = append(result, absoluteURLString)
		}

	})

	return result, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	result := []string{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		if val, exists := s.Attr("src"); exists {

			parsedURL, err := url.Parse(val)
			if err != nil {
				fmt.Printf("Couldn't parse URL: %v", err)
			}

			absoluteURL := baseURL.ResolveReference(parsedURL)
			absoluteURLString := absoluteURL.String()
			result = append(result, absoluteURLString)
		}
	})

	return result, nil
}
