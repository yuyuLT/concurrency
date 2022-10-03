package analysis

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

func Fetch(baseUrl *url.URL, doc *goquery.Document) (urls []string) {

	urls = make([]string, 0)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			reqUrl, err := baseUrl.Parse(href)
			if err == nil {
				if strings.Index(reqUrl.String(), "http") == 0 || strings.Index(reqUrl.String(), "/") == 0 {
					urls = append(urls, reqUrl.String())
				}
			}
		}
	})

	return
}
