package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://vm.globaloutlet-backend.com:8011/category")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	fmt.Println(res.Header)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	buffer := bytes.Buffer{}
	fmt.Println(buffer.ReadFrom(res.Body))

	for _, html := range doc.Nodes {
		fmt.Println(html.Data)
		fmt.Println(html.Attr)
	}
	// Find the review items
	doc.Find("input").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		s.SetHtml("123")
		//title := s.Find("i").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band, title)
		fmt.Printf("Review %d: %s - %s\n", i, s.Text())
		//fmt.Printf("Review %d: %s - %s\n", i, band)
	})

	doc.Find("button").Each(func(i int, selection *goquery.Selection) {
		fmt.Printf("button", selection.Text())
	})
}

func main() {
	ExampleScrape()
}
