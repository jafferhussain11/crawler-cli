package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	fmt.Println("Welcome to the CrAwLeR....")

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]

	fmt.Printf("starting crawl of: %s\n", baseURL)

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("error parsing base URL: %s\n", err)
		os.Exit(1)
	}

	cfg := config{
		pages:              make(map[string]PageData),
		mu:                 &sync.Mutex{},
		wg:                 &sync.WaitGroup{},
		concurrencyControl: make(chan struct{}, 7), //limit to 7
		baseURL:            parsedURL,
	}

	cfg.crawlPage(baseURL)
	cfg.wg.Wait()
	fmt.Println("Crawling done, Result below :")
	for k, page := range cfg.pages {
		//print key and value
		fmt.Printf("%s: %s\n", k, page)
	}

}
