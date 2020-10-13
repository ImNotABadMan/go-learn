package main

import (
	"fmt"
	"net/http"
)

func getLaravel(w http.ResponseWriter, req *http.Request) {
	//html := api_get.Get("http://192.168.10.113:8011" + req.URL.String())
	fmt.Println(req.URL)
	_, err := fmt.Fprint(w, "test")
	if err != nil {
		panic(err)
	}

}

func main() {
	http.HandleFunc("/", getLaravel)
	err := http.ListenAndServe(":8222", nil)
	fmt.Println("Listen 8222 ready...")
	if err != nil {
		panic(err)
	}
}
