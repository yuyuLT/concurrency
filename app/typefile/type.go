package typefile

type Result struct {
	Err error
	Url string
}

type Request struct {
	Url   string
	Depth int
}

type Channels struct {
	Req  chan Request
	Res  chan Result
	Quit chan int
}

type UrlStruct struct{
	Urls []string
	Err error
	IsWordpress bool
}

func NewChannels() *Channels {
	return &Channels{
		Req:  make(chan Request, 10),
		Res:  make(chan Result, 10),
		Quit: make(chan int, 10),
	}
}