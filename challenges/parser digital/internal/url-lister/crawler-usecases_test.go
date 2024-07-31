package uccrawler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"

	uccrawler "github.com/devpablocristo/crawler/internal/url-lister"
	crawler "github.com/devpablocristo/crawler/internal/url-lister/crawler"
)

func TestWebCrawler_CrawlURL(t *testing.T) {
	// Create a test server that returns a simple HTML
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<!DOCTYPE html>
		<html>
		<body>
			<a href="/page1">Page 1</a>
			<a href="/page2">Page 2</a>
		</body>
		</html>`))
	}))
	defer ts.Close()

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name         string
		args         args
		startURL     string
		expectedURLs []string
		wantedError  bool
	}{
		{
			name:     "Happy Path",
			startURL: ts.URL,
			args: args{
				ctx: context.Background(),
			},
			expectedURLs: []string{
				ts.URL,
				ts.URL + "/page1",
				ts.URL + "/page2",
			},
			wantedError: false,
		},
		{
			name:     "Not enough links",
			startURL: ts.URL,
			args: args{
				ctx: context.Background(),
			},
			expectedURLs: []string{
				ts.URL,
				ts.URL + "/page1",
			},
			wantedError: true,
		},
		{
			name:     "Too many links",
			startURL: ts.URL,
			args: args{
				ctx: context.Background(),
			},
			expectedURLs: []string{
				ts.URL,
				ts.URL + "/page1",
				ts.URL + "/page2",
				ts.URL + "/page3",
			},
			wantedError: true,
		},
	}
	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			htmlParser := crawler.NewHTMLParser()
			crawlerSrv := uccrawler.NewCrawler(htmlParser)

			visitedURLs := crawlerSrv.Crawl(tt.args.ctx, tt.startURL)

			sort.Strings(visitedURLs)
			sort.Strings(tt.expectedURLs)

			if !tt.wantedError && !reflect.DeepEqual(visitedURLs, tt.expectedURLs) {
				t.Errorf("Visited URLs do not match the expected ones. \nExpected: %v\nGot: %v", tt.expectedURLs, visitedURLs)
			}
			if tt.wantedError && reflect.DeepEqual(visitedURLs, tt.expectedURLs) {
				t.Errorf("Visited URLs do not match the expected ones. \nExpected: %v\nGot: %v", tt.expectedURLs, visitedURLs)
			}
		})
	}
}
