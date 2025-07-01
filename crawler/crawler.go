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

func extractLinks(resp *http.Response) []string {
	links := []string{}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return links
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type = html.ElementNode && n.Data == "a" {
			for _, atr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links
}