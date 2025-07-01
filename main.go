package main

import (
	"fmt"
	"go-webcrawler/crawler")

func main() {
	startUrl := "https://example.com"
	maxDepth := 3

	fmt.Println("Starting crawler...")
	crawler.Crawl(startUrl, maxDepth)
}