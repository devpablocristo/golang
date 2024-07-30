package xmlfile_test

import (
	"os"
	"testing"

	nu "github.com/devpablocristo/sitemap-generator/navigate-url"
	xml "github.com/devpablocristo/sitemap-generator/xml-file"
)

func TestCreateXML(t *testing.T) {

	nu.ListURLs = make(map[string]bool)
	nu.ListURLs = map[string]bool{
		"http://localhost:8081/index.html":     true,
		"http://localhost:8081/a1-depth1.html": true,
	}

	got, _ := xml.CreateXML("testfile.xml")
	data, _ := os.ReadFile("testfile.xml")
	want := string(data)

	if got != want {
		t.Errorf("\ngot %q \n\n want %q", got, want)
	}
}
