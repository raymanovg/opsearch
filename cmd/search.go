package main

import (
	"flag"
	"fmt"
	"net/url"
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

}
