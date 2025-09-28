package main

import (
	"fmt"
	"net/url"
	"sync"
)

//TODO : rawBaseURL = {string} "https://wikipedia.org"
//rawCurrentURL = {string} "https://en.wikipedia.org/" returns as different domain
//pages = {map[string]int}

type config struct {
	pages              map[string]PageData
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {

	if cfg.getLenOfPagesMap(cfg.pages) >= cfg.maxPages {
		return
	}

	baseURL := cfg.baseURL

	baseDomain := baseURL.Hostname()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current URL: %s\n", err)
	}

	currentDomain := currentURL.Hostname()

	if baseDomain != currentDomain {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing current URL: %s\n", err)
	}

	if cfg.updatePageVisit(normalizedCurrentURL) {
		return
	}

	//Get the HTML from the current URL,
	//and add a print statement so you can watch your crawler in real-time.
	//Assuming all went well with the request, get all the URLs from the response body HTML
	//Recursively crawl each URL on the page

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %s\n", err)
		return
	}
	fmt.Println("Successfully fetched " + rawCurrentURL + " page, processing...")

	pageData := extractPageData(html, rawCurrentURL)

	cfg.mu.Lock()
	cfg.pages[normalizedCurrentURL] = pageData
	cfg.mu.Unlock()

	//we can spawn new go routines
	for _, url := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go func() {
			defer cfg.wg.Done()
			cfg.concurrencyControl <- struct{}{}
			cfg.crawlPage(url)
			<-cfg.concurrencyControl
		}()
	}

}

func (cfg *config) getLenOfPagesMap(pages map[string]PageData) int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(pages)

}

func (cfg *config) updatePageVisit(normalizedCurrentURL string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normalizedCurrentURL]; ok {
		return true
	}
	return false
}
