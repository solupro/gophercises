package link

import (
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

func ParseHTML(url string) ([]Link, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return Parse(resp.Body)
}

func Parse(reader io.Reader) ([]Link, error) {
	node, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	return findA(node, []Link{}), nil
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
