package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {

	syGroup := sync.WaitGroup{}
	syGroup.Add(2)

	var url string

	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	log.Println("url: ", url)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		log.Fatal(response.Status)
	}

	buf := bytes.Buffer{}

	buf.ReadFrom(response.Body)

	urlSlice := strings.Split(buf.String(), "\r\n")

	sliLen := len(urlSlice)
	fmt.Println("url len: ", sliLen)

	var (
		onceLen = sliLen / 2
		times   = 0
	)

	for start := 0; start < sliLen; start += onceLen {
		go func(start int) {
			for i := start; i < start+onceLen; i++ {
				urlResponse, err := http.Get(urlSlice[i])
				if err != nil || urlResponse.StatusCode != 200 {
					fmt.Println("error index: ", i, urlSlice[i], urlResponse.Status, err)
				}
				fmt.Println("index: ", i, urlSlice[i], urlResponse.Status, err)

				time.Sleep(time.Millisecond * 100)
			}

			syGroup.Done()
		}(start)

		times++
	}

	fmt.Println("times: ", times)

	syGroup.Wait()

}
