package engine

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/raymanovg/opsearch/matcher"
)

type domainMatcher struct{}

var engine domainMatcher

func (dm *domainMatcher) Match(pageURL *url.URL, document *goquery.Document) []*url.URL {
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

func (dm *domainMatcher) genereateSearchVariants(pageURL *url.URL) []string {
	hostname := pageURL.Hostname()
	if strings.Contains(hostname, "www.") {
		hostname = strings.Replace(hostname, "www.", "", 1)
	}

	partsOfHostname := strings.Split(hostname, ".")
	var variants []string
	variants = append(variants, hostname)
	for _, partOne := range partsOfHostname {
		for _, partTwo := range partsOfHostname {
			if partOne == partTwo {
				continue
			}

			variantOne := partOne + partTwo
			variantTwo := partTwo + partOne

			exist := false
			for _, builtVariant := range variants {
				if builtVariant == variantOne {
					exist = true
					break
				}
			}

			if !exist {
				variants = append(variants, variantOne)
			}

			exist = false
			for _, builtVariant := range variants {
				if builtVariant == variantOne {
					exist = true
					break
				}
			}

			if !exist {
				variants = append(variants, variantTwo)
			}
		}
	}

	return variants
}

func init() {
	engine = domainMatcher{}
	matcher.Matchers[engine.StringName()] = &engine
}
