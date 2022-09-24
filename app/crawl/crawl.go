package crawl

import (
	"example.com/module/analysis"
	"example.com/module/typefile"
	"fmt"
)

func Crawl(url string, depth int, ch *typefile.Channels, wc *int) {
	defer func() { ch.Quit <- 0 }()

	Url := analysis.Analize(url)

	ch.Res <- typefile.Result{
		Url: url,
		Err: Url.Err,
		IsWordpress: Url.IsWordpress,
	}

	fmt.Println(wc)
	fmt.Println(*wc)

	if Url.Err == nil {
		for _, value := range Url.Urls {
			*wc ++
			fmt.Println("クロール ",*wc)
			ch.Req <- typefile.Request{
				Url: value,
				OriginUrl: url,
				Depth: depth - 1,
			}
		}
	}
}