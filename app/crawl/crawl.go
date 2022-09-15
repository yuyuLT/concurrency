package crawl

import (
	"example.com/module/analysis"
	"example.com/module/typefile"
)

func Crawl(url string, depth int, ch *typefile.Channels) {
	defer func() { ch.Quit <- 0 }()

	Url := analysis.Analize(url)

	ch.Res <- typefile.Result{
		Url: url,
		Err: Url.Err,
		IsWordpress: Url.IsWordpress,
	}

	if Url.Err == nil {
		for _, value := range Url.Urls {
			ch.Req <- typefile.Request{
				Url: value,
				OriginUrl: url,
				Depth: depth - 1,
			}
		}
	}
}