package matcher

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

var Matchers map[string]Matcher

type Matcher interface {
	Match(document *goquery.Document) []*url.URL
	StringName() string
}

func init() {
	Matchers = make(map[string]Matcher)
}
