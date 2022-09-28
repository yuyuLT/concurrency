package crawl

import (
	"example.com/module/analysis"
	"example.com/module/typefile"
)

func Crawl(url string, depth int, ch *typefile.Channels) {
	Url := analysis.Analize(url)

	ch.Res <- typefile.Result{
		Url: url,
		Err: Url.Err,
		IsWordpress: Url.IsWordpress,
	}

	if Url.Err == nil {
		for _, url := range Url.Urls {
			ch.Req <- typefile.Request{
				Url:   url,
				Depth: depth - 1,
			}
		}
	}
}