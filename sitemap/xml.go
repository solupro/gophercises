package sitemap

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type urlset struct {
	Xmlns string `xml:"xmlns,attr"`
	URLs  []*URL `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

func NewSiteMap() *urlset {
	return &urlset{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  make([]*URL, 0),
	}
}

func (s *urlset) AppendURL(url string) {
	s.URLs = append(s.URLs, &URL{Loc: url})
}

func (s *urlset) SaveTo(path string) error {
	data, err := xml.MarshalIndent(*s, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, os.ModePerm)
}
