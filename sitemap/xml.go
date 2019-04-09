package sitemap

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Sitemap struct {
	Urlset xml.Name `xml:"urlset"`
	Xmlns  string   `xml:"xmlns,attr"`
	URLs   []*URL   `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

func NewSiteMap() *Sitemap {
	return &Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  make([]*URL, 0),
	}
}

func (s *Sitemap) AppendURL(url string) {
	s.URLs = append(s.URLs, &URL{Loc: url})
}

func (s *Sitemap) SaveTo(path string) error {
	data, err := xml.Marshal(*s)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, os.ModePerm)
}
