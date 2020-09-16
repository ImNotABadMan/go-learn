package main

import (
	"fmt"
	"net/http"
)

func handler() {
	fmt.Fprint()
}

func main() {
	http.ListenAndServe("8080", handler)
}
