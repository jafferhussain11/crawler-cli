package main

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
)

//TODO : rawBaseURL = {string} "https://wikipedia.org"
//rawCurrentURL = {string} "https://en.wikipedia.org/" returns as different domain
//pages = {map[string]int}

type config struct {
	pages              sync.Map
	maxPages           int
	pageCount          atomic.Int32
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {

	if cfg.pageCount.Load() >= int32(cfg.maxPages) {
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

	if _, loaded := cfg.pages.LoadOrStore(normalizedCurrentURL, pageData); loaded {
		return
	}

	cfg.pageCount.Add(1)

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
