package main

import (
	"example.com/module/crawl"
	"example.com/module/typefile"
	"fmt"
	"log"
	"net/url"
	"os"
)

const crawlerDepthDefault = 3

var crawlerDepth int

func main() {

	startUrl := ""
	if crawlerDepth < 1 {
		crawlerDepth = crawlerDepthDefault
	}

	chs := typefile.NewChannels()
	urlMap := make(map[string]bool)
	hostMap := make(map[string]bool)
	wordPressMap := make(map[string]bool)

	chs.Req <- typefile.Request{
		Url:   startUrl,
		Depth: crawlerDepth,
	}

	wc := 0

	done := false
	for !done {
		select {
		case res := <-chs.Res:
			if res.Err == nil {
				fmt.Printf("Success %s || %d\n", res.Url, wc)
			} else {
				fmt.Fprintf(os.Stderr, "Error %s\n%v\n", res.Url, res.Err)
			}

			if res.IsWordpress {
				wordPressMap[res.Url] = true
			}
		case req := <-chs.Req:
			if req.Depth == 0 {
				break
			}

			u, err := url.Parse(req.Url)
			if err != nil {
				log.Fatal(err)
			}

			if urlMap[req.Url] {
				// 取得済み
				break
			}
			if hostMap[u.Host] {
				// 取得済み
			} else {
				fmt.Printf("%04d | Host = %s\n", len(hostMap), u.Host)
			}

			hostMap[u.Host] = true
			urlMap[req.Url] = true

			wc++
			go crawl.Crawl(req.Url, req.Depth, chs)
		case <-chs.Quit:
			wc--
		}

		if wc == 0 && len(chs.Req) == 0 {
			done = true
		}
	}

	fmt.Println("------------------")
	for key := range wordPressMap {
		fmt.Println(key)
	}
}
