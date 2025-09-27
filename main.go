package main

import (
	"fmt"
	"os"
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

	//call getHTML

	//htmlString, err := getHTML(baseURL)
	//if err != nil {
	//	fmt.Printf("error getting HTML: %s\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println(htmlString)

	//crawl
	crawlRes := make(map[string]int)

	crawlPage(baseURL, baseURL, crawlRes)

	fmt.Println(crawlRes)

}
