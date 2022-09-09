package analysis

import (
	"github.com/PuerkitoBio/goquery"
	"example.com/module/typefile"
	"net/http"
	"net/url"
	"fmt"
	"os"
)

func Analize(u string) (Url typefile.UrlStruct) {

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

	urls := Fetch(baseUrl,doc)
	if(CheckWp(baseUrl,doc)){
		fmt.Println(baseUrl)
		os.Exit(0)
	}

	Url = typefile.UrlStruct{urls,err}

	return
}