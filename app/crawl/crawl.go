package crawl

import (
	"example.com/module/analysis"
	"example.com/module/typefile"
)

func Crawl(url string, depth int, ch *typefile.Channels) {
	defer func() { ch.Quit <- 0 }()

	urls, err := analysis.Analize(url)

	ch.Res <- typefile.Result{
		Url: url,
		Err: err,
	}

	if err == nil {
		for _, url := range urls {
			ch.Req <- typefile.Request{
				Url:   url,
				Depth: depth - 1,
			}
		}
	}
}