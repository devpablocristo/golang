package crawler

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type HTMLParser struct{}

func NewHTMLParser() HTMLParserPort {
	return &HTMLParser{}
}

// ParseLinks implements the HTMLParser interface for PuerkitoBio/goquery.
func (p *HTMLParser) ParseLinks(body []byte) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	var links []string

	// Find all anchor tags and extract the href attribute
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, href)
		}
	})

	return links, nil
}
