package uccrawler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	crawler "github.com/devpablocristo/crawler/internal/url-lister/crawler"
)

// HTMLParser is an interface representing the functionality needed for parsing HTML.
type HTMLParser interface {
	ParseLinks(body []byte) ([]string, error)
}

// Crawler is a type representing the crawling service.
type Crawler struct {
	baseURL     string
	visitedURLs []string
	visitedMux  sync.Mutex
	wg          sync.WaitGroup
	htmlParser  HTMLParser
}

// NewCrawler initializes a new web crawler with an injected HTML parser.
func NewCrawler(hp HTMLParser) crawler.CrawlerPort {
	return &Crawler{
		visitedURLs: make([]string, 0),
		htmlParser:  hp,
	}
}

// Crawl starts crawling from the initial URL and returns the found URLs.
func (w *Crawler) Crawl(ctx context.Context, startURL string) []string {
	w.wg.Add(1)
	w.baseURL = startURL
	w.crawlURL(ctx, w.baseURL)
	w.wg.Wait()
	return w.visitedURLs
}

// crawlURL is a private method to visit a web page.
func (w *Crawler) crawlURL(ctx context.Context, url string) {
	defer w.wg.Done()

	w.visitedMux.Lock()
	defer w.visitedMux.Unlock()

	// Check if the URL has already been visited
	for _, visitedURL := range w.visitedURLs {
		if visitedURL == url {
			return
		}
	}

	// Mark the current URL as visited
	w.visitedURLs = append(w.visitedURLs, url)

	fmt.Println("Visiting:", url)

	// Fetch the HTML content of the page
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("There was a problem getting the page:", err)
		return
	}
	defer resp.Body.Close()

	// Read the content of the response body into a []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("There was a problem reading the response body:", err)
		return
	}

	// Parse the HTML content using the injected HTMLParser
	links, err := w.htmlParser.ParseLinks(body)
	if err != nil {
		fmt.Println("There was a problem parsing the page:", err)
		return
	}

	// Process the extracted links
	for _, link := range links {
		absoluteURL := w.resolveURL(ctx, link, url)
		fmt.Println("    Found a link:", absoluteURL)

		// If the link is from the same domain, initiate crawling
		if w.isSameDomain(ctx, absoluteURL) {
			w.wg.Add(1)
			go w.crawlURL(ctx, absoluteURL)
		}
	}
}

// resolveURL is a private method to determine if a URL is relative or absolute and make it absolute.
func (w *Crawler) resolveURL(ctx context.Context, href, baseURL string) string {
	if strings.HasPrefix(href, "http") {
		// If the URL is already complete, leave it as is
		return href
	}
	if strings.HasPrefix(href, "/") {
		// If the URL is relative and starts with "/", make it complete with the page's base
		// and avoid duplicating "/page2" multiple times.
		return strings.TrimRight(w.baseURL, "/") + href
	}
	// If it doesn't start with "/", then it is part of the current URL
	return baseURL + "/" + href
}

// isSameDomain is a private method to determine if two URLs are part of the same web page.
func (w *Crawler) isSameDomain(ctx context.Context, url string) bool {
	return strings.HasPrefix(url, w.baseURL)
}
