package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
)

func main() {
	var (
		shopifySessionID = "0776bbe3c66bdd1195e4f0dc6873182e"
		url              = "https://reeneefashion.myshopify.com/admin/products/6639277768865"
		inputClass       = "Polaris-TextField__Input_30ock"
		headers          = http.Header{}
	)

	collect := colly.NewCollector()

	collect.OnResponse(func(response *colly.Response) {
		fmt.Println("Response Body", string(response.Body))
		fmt.Println("Response Headers", response.Headers)
		//for _, cookie := range cookies {
		//	fmt.Println(len(cookie))
		//	if len(cookie) > 0 {
		//		cookieCh <- cookie
		//	}
		//}
		fmt.Println()
		fmt.Println("Response status code", response.StatusCode)
	})

	collect.OnHTML("."+inputClass, func(element *colly.HTMLElement) {
		fmt.Println(element.Attr("value"))
	})

	headers.Set("cookie", "_secure_admin_session_id_csrf="+shopifySessionID)
	err := collect.Request("GET", url, nil, nil, headers)
	if err != nil {
		log.Fatal(err)
	}

}
