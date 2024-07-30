package crawler

import (
	"context"
)

//go:generate mockgen -source=./ports.go -destination=./mocks/htmlparser_adapter/htmlparser_adapter_mock.go -package=mocks
type HTMLParserPort interface {
	ParseLinks(body []byte) ([]string, error)
}

type CrawlerPort interface {
	Crawl(context.Context, string) []string
}
