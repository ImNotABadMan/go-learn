package spider

type Request struct {
	Body string
	Urls []string
	*Request
}

type Fetcher interface {
	Fetch(url string)
}

func (request *Request) Add(subRequest *Request) {
	request.Request = subRequest
}

func fetch() {

}

func Crawl(request Request, depth int) {

}
