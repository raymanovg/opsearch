package flag

import (
	"fmt"
	"net/url"
)

type URLFlag struct {
	URL *url.URL
}

func (uf *URLFlag) String() string {
	if uf.URL == nil {
		return ""
	}

	return uf.URL.String()
}

func (uf *URLFlag) Set(value string) error {
	parsedURL, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("Cannot parse url: %v", err)
	}

	uf.URL = parsedURL

	return nil
}
