package engine

import (
	"fmt"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/raymanovg/opsearch/matcher"
)

type domainMatcher struct{}

var engine domainMatcher

func (dm *domainMatcher) Match(document *goquery.Document) []*url.URL {
	var links []*url.URL
	for _, service := range matcher.Services {
		fmt.Printf("Loking for links for %s \n", service.Name)
		for _, domain := range service.Domains {
			pattern := fmt.Sprintf("a[href*=\"%s\"]", domain)
			document.Find(pattern).Each(func(i int, s *goquery.Selection) {
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
		}
	}

	return links
}

func (dm *domainMatcher) StringName() string {
	return "domainMatcher"
}

func init() {
	engine = domainMatcher{}
	matcher.Matchers[engine.StringName()] = &engine
}
