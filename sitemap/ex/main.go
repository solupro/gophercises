package main

import (
	"flag"
	"github.com/solupro/gophercises/link"
	"github.com/solupro/gophercises/sitemap"
	urlLib "net/url"
)

var host string

func main() {
	var url string

	flag.StringVar(&url, "url", "", "sitemap for this url")
	flag.Parse()
	if url == "" {
		flag.PrintDefaults()
		return
	}

	urlInfo, err := urlLib.Parse(url)
	if err != nil {
		panic(err)
	}

	host = urlInfo.Host
	links := make(map[string]int)
	parse(url, links)

	s := sitemap.NewSiteMap()
	for k, _ := range links {
		s.AppendURL(k)
	}

	err = s.SaveTo("./sitemap.xml")
	if err != nil {
		panic(err)
	}
}

func parse(url string, links map[string]int) {
	if _, ok := links[url]; ok {
		//fmt.Println(ok, url)
		return
	}
	links[url] = 1

	//fmt.Println("parse:", url)
	urls, err := link.ParseHTML(url)
	if err != nil {
		return
	}

	for _, v := range urls {
		info, err := urlLib.Parse(v.Href)
		if err != nil {
			continue
		}

		if info.Host != host {
			continue
		}

		parse(v.Href, links)
	}

}
