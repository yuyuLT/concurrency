package analysis

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

func CheckWp(baseUrl *url.URL, doc *goquery.Document) (isWp bool) {

	isWp = false

	doc.Find("link").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			reqUrl, err := baseUrl.Parse(href)
			if err == nil {
				if strings.Contains(reqUrl.String(), "wp-content") {
					isWp = true
				}
			}
		}
	})

	return
}
