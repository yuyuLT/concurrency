package analysis

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"regexp"
)

func CheckWp(baseUrl *url.URL, doc *goquery.Document) (isWp bool) {

	isWp = false

	doc.Find("link").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			reqUrl, err := baseUrl.Parse(href)
			if err == nil {
				r := regexp.MustCompile(`wp-content.*css`)
				if r.MatchString(reqUrl.String()){
					isWp = true
				}
			}
		}
	})
	
	return
}