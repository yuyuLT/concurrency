package main

import (
	"fmt"
	"log"
	"os"
	"net/url"
	"example.com/module/crawl"
	"example.com/module/typefile"
	"time"
)


const crawlerDepthDefault = 2

var crawlerDepth int

func main() {

	now := time.Now() 

	startUrl := "https://www.creatures.co.jp/"
	if crawlerDepth < 1 {
		crawlerDepth = crawlerDepthDefault
	}

	chs := typefile.NewChannels()
	urlMap := make(map[string]string)
	hostMap := make(map[string]bool)
	wordPressMap := make(map[string]bool)

	chs.Req <- typefile.Request{
		Url:   startUrl,
		Depth: crawlerDepth,
	}

	wc := 1

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
				wc --
				if wc == 0 {
					done = true
				}
				break
			}

			u, err := url.Parse(req.Url)
			if err != nil {
				log.Fatal(err)
			}

			_, ok := urlMap[req.Url]
			if ok {
				// 取得済み
				break
			}
			if hostMap[u.Host] {
				// 取得済み
			} else {
				fmt.Printf("%04d | Host = %s\n", len(hostMap), u.Host)
			}

			hostMap[u.Host] = true
			urlMap[req.Url] = req.OriginUrl
			go crawl.Crawl(req.Url, req.Depth, chs, &wc)
		case <-chs.Quit:
			wc--
			if wc == 0 {
				done = true
			}
		}
	}

	fmt.Println("------------------")
	for key,value := range urlMap {
		_, ok := wordPressMap[key]
		if ok {
			fmt.Print(key)
			fmt.Print("  ")
			fmt.Println(value)
		}
	}

	fmt.Printf("経過: %vms\n", time.Since(now).Milliseconds()) 
}