package gl_colly

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
)

//var endPoint = "http://vm.globaloutlet-backend.com:8011"
var endPoint = "http://v2.globaloutlet-backend.com:8011"

var token string

var headers *http.Header

var loginColly *colly.Collector

//func Call() map[string]string {
func Call() []*http.Cookie {
	collect := colly.NewCollector()
	cookieCh := make(chan string, 2)
	//resMap := make(map[string]string)

	collect.OnRequest(func(request *colly.Request) {

		fmt.Println(request.Body)
		fmt.Println("Headers", request.Headers)
		fmt.Println("url", request.URL)

		fmt.Println()

	})

	collect.OnResponse(func(response *colly.Response) {
		//fmt.Println("Response Body", string(response.Body))
		fmt.Println("Response Headers", response.Headers)
		headers = response.Headers
		fmt.Println()
		cookies := response.Headers.Values("Set-Cookie")
		//for _, cookie := range cookies {
		//	fmt.Println(len(cookie))
		//	if len(cookie) > 0 {
		//		cookieCh <- cookie
		//	}
		//}
		close(cookieCh)
		fmt.Println("Response Set-Cookie", cookies)

		fmt.Println()
		fmt.Println("Response status code", response.StatusCode)
	})

	collect.OnHTML("input[name=_token]", func(element *colly.HTMLElement) {
		token = element.Attr("value")
		fmt.Println("now ", token)
	})

	fmt.Println("visit error:", collect.Visit(endPoint+"/login"))

	//select {
	//case cookie := <-cookieCh:
	//	resMap["Cookie"] = cookie
	//default:
	//}

	loginColly = collect.Clone()

	fmt.Println("test")

	return collect.Cookies(endPoint)
}

func GetCate(cookies map[string]string) {

}

func Login(cookies []*http.Cookie) error {
	//loginColly := colly.NewCollector()

	//loginColly.SetCookies(url, cookies)

	fmt.Println(loginColly.String())

	// _token=d8WBHJkyPHY6gOnECU0sAaB5EN3YvtkojsEUQ4sA&username=binz&password=binz123&remember=1

	data := map[string]string{
		"_token":   token,
		"username": "binz",
		"password": "binz123",
		"remeber":  "1",
	}
	//loginColly = loginColly.Clone()

	//form := url.Values{}
	//for k, v := range data {
	//	form.Add(k, v)
	//}

	buf := bytes.Buffer{}

	for key, value := range data {
		buf.WriteString(key + "=" + value + "&")
	}
	buf.Truncate(len(buf.String()) - 1)
	//buf1 := strings.NewReader(form.Encode())
	err := loginColly.Request("POST", endPoint+"/login", &buf, nil, nil)

	//loginColly = loginColly.Clone()

	fmt.Println("headers ----------------------- ")
	//fmt.Println(headers)
	//fmt.Println()
	//cookie := headers.Get("Set-Cookie")
	//headers.Set("Cookie", cookie)
	//headers.Del("Set-Cookie")
	//fmt.Println(headers)
	fmt.Println("headers ----------------------- ")

	//err := loginColly.Post(url+"/login", data)

	//if err != nil {
	//	fmt.Println("login", err)
	//}

	loginColly.OnRequest(func(request *colly.Request) {

		//request.Headers.Set("User-Agent", "colly - https://github.com/gocolly/colly")
		//request.Headers.Del("Set-Cookie")

		fmt.Println(request.Body)
		fmt.Println("Headers", request.Headers)
		fmt.Println("url", request.URL)
		fmt.Println("type", request.Method)

		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()

	})

	loginColly.OnResponse(func(response *colly.Response) {
		fmt.Println("login response")
		fmt.Println(string(response.Body))
		fmt.Println(response.Headers)
		fmt.Println(response.StatusCode)
	})

	loginColly.OnHTML("input[name=_token]", func(element *colly.HTMLElement) {
		//token = element.Attr("value")
		//fmt.Println("now ", token)
	})

	err = loginColly.Visit(endPoint + "/product/csv/index")

	loginColly.Clone()

	fmt.Println(cookies)

	return err
}
