package main

import (
	"bytes"
	"demo/proto_demo/protobuf/from_php"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
)

func main() {
	buf := bytes.Buffer{}

	r, err := http.Get("http://192.168.10.113:18009/proto")

	if err != nil {
		log.Fatal(err)
	}

	if _, err = buf.ReadFrom(r.Body); err != nil {
		log.Println(err)
	}

	helloFromPHP := from_php.Hello{}

	body := proto.Unmarshal(buf.Bytes(), &helloFromPHP)

	fmt.Println(body)
	fmt.Println(helloFromPHP.GetName())
	fmt.Println(helloFromPHP)
	//fmt.Println(helloFromPHP.GetText())
	fmt.Println(helloFromPHP.GetDes())
	fmt.Println(helloFromPHP.String())
}
