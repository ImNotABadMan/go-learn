package main

import (
	//api_get "api.get"
	"net/http"
)
import "fmt"

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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
