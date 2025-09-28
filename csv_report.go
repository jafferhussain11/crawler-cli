package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
)

func writeCSVReport(pages *sync.Map, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	// Write header
	err = csvWriter.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"})
	if err != nil {
		fmt.Printf("Error writing header to csv file: %v", err)
		return err
	}

	// Use Range to iterate over sync.Map
	pages.Range(func(key, value interface{}) bool {
		// Type assert the value to PageData
		page, ok := value.(PageData)
		if !ok {
			fmt.Printf("Warning: invalid page data for key %v\n", key)
			return true // Continue iteration
		}

		// Write the record
		err := csvWriter.Write([]string{
			page.URL,
			page.H1,
			page.FirstParagraph,
			strings.Join(page.OutgoingLinks, ";"),
			strings.Join(page.ImageURLs, ";"),
		})
		if err != nil {
			fmt.Printf("Error writing record to csv file: %v\n", err)
			// You might want to return false here to stop iteration on error
			// return false
		}

		return true // Continue to next item
	})

	return nil
}
