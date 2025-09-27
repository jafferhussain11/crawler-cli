package main

import (
	"fmt"
	"net/url"
)

//Here's my pseudocode:
//
//-Make sure the rawCurrentURL is on the same domain as the rawBaseURL.
//-If it's not, just return. We don't want to crawl the entire internet, just the domain in question.

//Get a normalized version of the rawCurrentURL.
//If the pages map already has an entry for the normalized version of the current URL, just increment the count
//and be done, we've already crawled this page.
//Otherwise, add an entry to the pages map for the normalized version of the current URL, and set the count to 1.

//Get the HTML from the current URL,
//and add a print statement so you can watch your crawler in real-time.
//Assuming all went well with the request, get all the URLs from the response body HTML
//Recursively crawl each URL on the page

//Be careful testing this! Be sure to add print statements so you can see what your crawler is doing,
//and kill it with ctrl+c if it's stuck in a loop.
//If you make too many spammy requests to a website (including my blog) you could get your IP address blocked.

//Call crawlPage in the main function instead of getHTML.

//When it's complete, print the keys and values of the pages map to the console.
//Test your program by running it against a small site (10–50 pages, like my blog).
//When you're satisfied that everything is working, you can move on.

//TODO : rawBaseURL = {string} "https://wikipedia.org"
//rawCurrentURL = {string} "https://en.wikipedia.org/"
//pages = {map[string]int}

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL: %s\n", err)
	}

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

	if _, ok := pages[normalizedCurrentURL]; ok {
		pages[normalizedCurrentURL]++
		return
	}
	pages[normalizedCurrentURL] = 1

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

	urls, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Printf("error getting URLs from HTML: %s\n , will not proceed with further crawling", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}

}
