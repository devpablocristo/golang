package xmlfile

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"time"

	nu "github.com/devpablocristo/sitemap-generator/navigate-url"
)

type Url struct {
	XMLName      xml.Name `xml:"url"`
	Loc          string   `xml:"loc"`
	LastModified string   `xml:"lastmod"`
}

func CreateXML(dir string) error {
	var cont *Url
	var xmlstring []byte
	var err error
	var t time.Time = time.Now()
	var body string

	for u := range nu.ListURLs {
		cont = &Url{
			Loc:          u,
			LastModified: t.Format("2006-01-02"),
		}

		xmlstring, err = xml.MarshalIndent(cont, " ", "  ")
		if err != nil {
			return err
		}

		body += string(xmlstring) + "\n"
	}

	result := xml.Header +
		"<urlset xmlns='http://www.sitemaps.org/schemas/sitemap/0.9'>" +
		body +
		"</urlset>"

	err = ioutil.WriteFile(dir, []byte(result), 0776)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
