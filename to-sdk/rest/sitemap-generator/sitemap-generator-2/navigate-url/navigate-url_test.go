package navigateurl_test

import (
	"reflect"
	"testing"

	nu "github.com/devpablocristo/sitemap-generator/navigate-url"
)

func TestNavigateUrl(t *testing.T) {

	t.Run("Test http://localhost:8081/ depth 0", func(t *testing.T) {
		depth := 0
		depth++

		startURL := "http://localhost:8081/index.html"

		nu.ListURLs = make(map[string]bool)
		nu.ListURLs[startURL] = true

		nu.NavigateUrl(startURL, depth)

		got := nu.ListURLs
		want := map[string]bool{
			"http://localhost:8081/b1-depth1.html": true,
			"http://localhost:8081/c1-depth1.html": true,
			"http://localhost:8081/index.html":     true,
			"http://localhost:8081/a1-depth1.html": true,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v \n\n want %v", got, want)
		}
	})

	t.Run("Test http://localhost:8081 depth 4", func(t *testing.T) {
		depth := 4
		depth++

		startURL := "http://localhost:8081/index.html"
		nu.ListURLs = make(map[string]bool)
		nu.ListURLs[startURL] = true

		nu.NavigateUrl(startURL, depth)

		got := nu.ListURLs
		want := map[string]bool{
			"http://localhost:8081/index.html":     true,
			"http://localhost:8081/b1-depth1.html": true,
			"http://localhost:8081/b2-depth2.html": true,
			"http://localhost:8081/b4-depth3.html": true,
			"http://localhost:8081/c2-depth2.html": true,
			"http://localhost:8081/b3-depth2.html": true,
			"http://localhost:8081/c1-depth1.html": true,
			"http://localhost:8081/c3-depth3.html": true,
			"http://localhost:8081/a1-depth1.html": true,
			"http://localhost:8081/a2-depth2.html": true,
			"http://localhost:8081/a3-depth2.html": true,
			"http://localhost:8081/c4-depth4.html": true,
			"http://localhost:8081/c5-depth5.html": true,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v \n\n want %v", got, want)
		}
	})

	t.Run("Test https://www.example.com depth 0", func(t *testing.T) {
		depth := 0
		depth++

		startURL := "https://www.example.com"
		nu.ListURLs = make(map[string]bool)
		nu.ListURLs[startURL] = true

		nu.NavigateUrl(startURL, depth)

		got := nu.ListURLs
		want := map[string]bool{
			"https://www.example.com":              true,
			"https://www.iana.org/domains/example": true,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v \n\n want %v", got, want)
		}
	})

}
