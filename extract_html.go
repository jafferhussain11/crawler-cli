package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {

	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	h1Text := document.Find("H1").Text()
	return h1Text
}

func getFirstParagraphFromHTML(html string) string {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	documentWithMain := document.Find("main")

	if documentWithMain.Text() == "" {
		return document.Find("p").First().Text()
	}
	return documentWithMain.Find("p").First().Text()
}
