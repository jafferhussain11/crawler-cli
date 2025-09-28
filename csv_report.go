package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	err = csvWriter.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"})
	if err != nil {
		fmt.Printf("Error writing header to csv file: %v", err)
		return err
	}

	for _, page := range pages {
		err = csvWriter.Write([]string{page.URL, page.H1,
			page.FirstParagraph,
			strings.Join(page.OutgoingLinks, ";"),
			strings.Join(page.ImageURLs, ";"),
		})
		if err != nil {
			fmt.Printf("Error writing a record to csv file: %v", err)
		}
	}

	return nil

}
