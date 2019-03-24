package engine

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"runtime"

	"github.com/PuerkitoBio/goquery"
	"github.com/raymanovg/opsearch/matcher"
)

type (
	domainMatcher struct{}
	Config        struct {
		Settings []struct {
			Sf      string   `json:"sf"`
			Domains []string `json:"domains"`
		}
	}
)

var config Config
var engine domainMatcher

func (dm *domainMatcher) Match(document *goquery.Document) []*url.URL {
	var links []*url.URL
	for _, setting := range config.Settings {
		fmt.Printf("Loking for links for %s \n", setting.Sf)
		for _, domain := range setting.Domains {
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

func loadConfig() {
	_, filename, _, _ := runtime.Caller(1)
	configFileName := path.Join(path.Dir(filename), engine.StringName()+"Config.json")
	file, err := os.Open(configFileName)
	if err != nil {
		fmt.Printf("Cannot open config file: %v", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Printf("Cannot decode config: %v", err)
		return
	}
}

func init() {
	engine = domainMatcher{}
	matcher.Matchers[engine.StringName()] = &engine
	loadConfig()
}
