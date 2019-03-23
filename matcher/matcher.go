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

type domainMatcher struct{}

func (dm *domainMatcher) Match(document *goquery.Document) []*url.URL {
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

func (dm *domainMatcher) StringName() string {
	return "DomainMatcher"
}

func init() {
	Matchers = make(map[string]Matcher)
	matcher := &domainMatcher{}
	Matchers[matcher.StringName()] = matcher
}