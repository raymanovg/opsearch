package matcher

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type service struct {
	Name    string
	Domains []string
}

var Services = []service{
	service{
		Name: "Facebook",
		Domains: []string{
			"facebook.com",
		},
	},
	service{
		Name: "Twitter",
		Domains: []string{
			"twitter.com",
		},
	},
	service{
		Name: "VK",
		Domains: []string{
			"vk.com",
			"vkontakte.ru",
		},
	},
	service{
		Name: "Ok",
		Domains: []string{
			"ok.ru",
		},
	},
}

var Matchers map[string]Matcher

type Matcher interface {
	Match(document *goquery.Document) []*url.URL
	StringName() string
}

func init() {
	Matchers = make(map[string]Matcher)
}
