package fetch

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

func Fetch(u string) (urls []string, err error) {
	baseUrl, err := url.Parse(u)
	if err != nil {
		return
	}

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

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