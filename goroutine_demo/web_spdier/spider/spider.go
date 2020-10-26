package spider

type Request struct {
	Body string
	Urls []string
	*Request
}

type Fetcher interface {
	Fetch()
}

func (request *Request) Add(subRequest *Request) {
	request.Request = subRequest
}

func fetch() {

}
