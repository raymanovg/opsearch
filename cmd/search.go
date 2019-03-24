package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	cmdFlag "github.com/raymanovg/opsearch/cmd/flag"
	_ "github.com/raymanovg/opsearch/matcher/engine"
)

var (
	pageURLArg       cmdFlag.URLFlag
	matcherEngineArg cmdFlag.MatcherFlag
)

func init() {
	flag.Var(&pageURLArg, "url", "Page url to search official page links")
	flag.Var(&matcherEngineArg, "matcher", "Name of matcher engine")
}

func main() {
	flag.Parse()

	pageURL := pageURLArg.URL
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

	return matcherEngineArg.Engine.Match(document), nil
}

func loadPageContent(pageURL string) (io.ReadCloser, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
