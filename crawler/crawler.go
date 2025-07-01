package crawler

import ("fmt"
"net/http"
"golang.org/x/net/html")

func Crawl(url string, depth int) {
	if depth <= 0 {
		return
	}

	fmt.Println("Crawling:", url)


	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	links := extractLinks(resp)

	for _, link := range links {
		fmt.Println("Found link:", link)
		Crawl(link, depth-1)
	}
}

func extractLinks(resp *http.Response) {
	fmt.Println("Extracting links from the response body")
}