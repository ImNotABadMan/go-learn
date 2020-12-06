package main

import (
	"demo/colly_demo/gl_colly"
	"fmt"
)

func main() {
	getGlCate()
}

func getGlCate() {
	cookies := gl_colly.Call()
	fmt.Println("login error", gl_colly.Login(cookies))
}
