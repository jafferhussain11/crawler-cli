package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("Welcome to the CrAwLeR....")

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]
	maxPages, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("invalid max pages")
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("invalid max concurrency")
		os.Exit(1)
	}

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
		concurrencyControl: make(chan struct{}, maxConcurrency), //limit to 7
		baseURL:            parsedURL,
		maxPages:           maxPages,
	}

	cfg.crawlPage(baseURL)
	cfg.wg.Wait()
	fmt.Println("Crawling done, Result below :")
	for k, page := range cfg.pages {
		//print key and value
		fmt.Printf("%s: %s\n", k, page)
	}

}
