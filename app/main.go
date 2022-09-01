package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Request struct {
	url   string
	depth int
}

type Result struct {
	err error
	url string
}

type Channels struct {
	req  chan Request
	res  chan Result
	quit chan int
}

func NewChannels() *Channels {
	return &Channels{
		req:  make(chan Request, 10),
		res:  make(chan Result, 10),
		quit: make(chan int, 10),
	}
}

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

func Crawl(url string, depth int, ch *Channels) {
	defer func() { ch.quit <- 0 }()

	urls, err := Fetch(url)

	ch.res <- Result{
		url: url,
		err: err,
	}

	if err == nil {
		for _, url := range urls {
			ch.req <- Request{
				url:   url,
				depth: depth - 1,
			}
		}
	}
}

const crawlerDepthDefault = 5

var crawlerDepth int

func main() {
	/*
		flag.IntVar(&crawlerDepth, "depth", crawlerDepthDefault, "クロールする深さを指定。")
		flag.Parse()

		if len(flag.Args()) < 1 {
			fmt.Fprintln(os.Stderr, "URLを指定してください。")
			os.Exit(1)
		}
		startUrl := flag.Arg(0)
	*/
	startUrl := "https://www.e2e-inc.co.jp"
	if crawlerDepth < 1 {
		crawlerDepth = crawlerDepthDefault
	}

	chs := NewChannels()
	urlMap := make(map[string]bool)
	hostMap := make(map[string]bool)

	chs.req <- Request{
		url:   startUrl,
		depth: crawlerDepth,
	}

	wc := 0

	done := false
	for !done {
		select {
		case res := <-chs.res:
			if res.err == nil {
				fmt.Printf("Success %s || %d\n", res.url, wc)
			} else {
				fmt.Fprintf(os.Stderr, "Error %s\n%v\n", res.url, res.err)
			}
		case req := <-chs.req:
			if req.depth == 0 {
				break
			}

			u, err := url.Parse(req.url)
			if err != nil {
				log.Fatal(err)
			}

			//if urlMap[u.Host] {
			if urlMap[req.url] {
				// 取得済み
				break
			}
			if hostMap[u.Host] {
				// 取得済み
			} else {
				fmt.Printf("%04d | Host = %s\n", len(hostMap), u.Host)
			}

			hostMap[u.Host] = true
			urlMap[req.url] = true

			wc++
			go Crawl(req.url, req.depth, chs)
		case <-chs.quit:
			wc--
			if wc == 0 {
				done = true
			}
		}
	}
}