package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type URLFlag struct {
	url *url.URL
}

func (f *URLFlag) String() string {
	if f.url == nil {
		return ""
	}

	return f.url.String()
}

func (f *URLFlag) Set(value string) error {
	parsedURL, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("Cannot parse url: %v", err)
	}

	f.url = parsedURL

	return nil
}

var arg URLFlag

func init() {
	flag.Var(&arg, "url", "Page url to search official page links")
}

func main() {
	flag.Parse()

	pageURL := arg.url
	fmt.Printf("Looging for op links on page %s", pageURL.String())
	links, err := search(pageURL)
	if err != nil {
		fmt.Printf("Could not find op links: %v", err)
		return
	}

	for _, link := range links {
		fmt.Printf("\n -- %s", link.String())
	}
}

func search(pageURL *url.URL) ([]*url.URL, error) {
	contentBody, err := loadPageContent(pageURL.String())
	if err != nil {
		return nil, err
	}

	defer contentBody.Close()

	document, err := goquery.NewDocumentFromReader(contentBody)
	if err != nil {
		return nil, err
	}

	return scrapLinks(document), nil
}

func loadPageContent(pageURL string) (io.ReadCloser, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func scrapLinks(document *goquery.Document) []*url.URL {
	var links []*url.URL
	document.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exist := s.Attr("href")
		if !exist {
			return
		}

		parsedLink, err := url.Parse(link)
		if err != nil {
			return
		}

		links = append(links, parsedLink)
	})

	return links
}
