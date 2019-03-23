package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/raymanovg/opsearch/matcher"
)

type (
	MatcherFlag struct {
		match matcher.Matcher
	}

	URLFlag struct {
		url *url.URL
	}
)

func (mf *MatcherFlag) String() string {
	if mf.match == nil {
		return ""
	}

	return mf.match.StringName()
}

func (mf *MatcherFlag) Set(value string) error {
	if match, exist := matcher.Matchers[value]; exist {
		mf.match = match
		return nil
	}

	return fmt.Errorf("There is no matcher %s", value)

}

func (uf *URLFlag) String() string {
	if uf.url == nil {
		return ""
	}

	return uf.url.String()
}

func (uf *URLFlag) Set(value string) error {
	parsedURL, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("Cannot parse url: %v", err)
	}

	uf.url = parsedURL

	return nil
}

var (
	URLArg        URLFlag
	matcherEngine MatcherFlag
)

func init() {
	flag.Var(&URLArg, "url", "Page url to search official page links")
	flag.Var(&matcherEngine, "matcher", "Name of matcher engine")
}

func main() {
	flag.Parse()

	pageURL := URLArg.url
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

	return matcherEngine.match.Match(document), nil
}

func loadPageContent(pageURL string) (io.ReadCloser, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
