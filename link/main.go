package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./ex2.html")
	if err != nil {
		panic(err)
	}

	node, err := html.Parse(f)
	if err != nil {
		panic(err)
	}

	list := []Link{}
	list = findA(node, list)

	fmt.Printf("%+v", list)
}

func findA(node *html.Node, result []Link) []Link {
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "a" && c.Type == html.ElementNode {
			link := Link{}

			if c.FirstChild != nil && c.FirstChild.Type == html.TextNode {
				link.Text = strings.TrimSpace(c.FirstChild.Data)
			}

			for _, v := range c.Attr {
				if v.Key == "href" {
					link.Href = v.Val
					break
				}
			}

			if link.Href != "#" {
				result = append(result, link)
			}

		}

		result = findA(c, result)
	}

	return result
}

type Link struct {
	Href string
	Text string
}
