package api_get

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	html, er := ioutil.ReadAll(res.Body)

	if er != nil {
		panic(er)
	}

	return string(html)
}
