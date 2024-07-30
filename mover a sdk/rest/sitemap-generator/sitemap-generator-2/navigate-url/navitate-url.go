package navigateurl

import (
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

var ListURLs = make(map[string]bool)

func NavigateUrl(navURL string, depth int) {
	if depth <= 0 {
		return
	}

	parsedURL, err := url.Parse(navURL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := GetBody(parsedURL.String())
	if err != nil {
		log.Fatal(err)
	}

	unscapedBody := html.UnescapeString(body)

	var urls = make([]string, 0)
	urls, err = ExtractURLs(parsedURL, unscapedBody)
	if err != nil {
		log.Fatal(err)
	}

	depth--
	for _, nextURL := range urls {
		if !checkForURL(nextURL) {
			ListURLs[nextURL] = true
			NavigateUrl(nextURL, depth)
		}
	}
}

func GetBody(getBodyURL string) (string, error) {
	resp, err := http.Get(getBodyURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	sbody := string(body)

	return sbody, nil

}

func ExtractURLs(gotURL *url.URL, content string) ([]string, error) {
	var err error
	var urls []string = make([]string, 0)
	var findURLs = regexp.MustCompile("<a.*?href=\"(.*?)\"")

	//extracts matches and captured submatches
	matches := findURLs.FindAllStringSubmatch(content, -1)

	for _, val := range matches {
		var linkUrl *url.URL

		linkUrl, err = url.Parse(val[1])
		if err != nil {
			return urls, err
		}

		if linkUrl.IsAbs() {
			urls = append(urls, linkUrl.String())
		} else {
			urls = append(urls, gotURL.Scheme+"://"+gotURL.Host+linkUrl.String())
		}
	}

	return urls, err
}

func checkForURL(link string) bool {
	for u := range ListURLs {
		if u == link {
			return true
		}
	}
	return false
}

func Worker(navURL string, depth int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	NavigateUrl(navURL, depth)
	m.Unlock()
}
