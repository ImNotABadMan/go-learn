package main

import (
	"demo/proto_demo/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func main() {
	//sig := make(chan os.Signal)
	//signal.Notify(sig, syscall.SIGINT)
	log.Println("start Server :9999")
	http.HandleFunc("/hello", handle)

	err := http.ListenAndServe(":9999", nil)

	log.Println("start Server")

	if err != nil {
		log.Fatal("start http fail", err)
	}

	//select {
	//case <-sig:
	//	log.Println("close http")
	//	os.Exit(0)
	//}

}

func handle(w http.ResponseWriter, r *http.Request) {
	h := protobuf.Hello{
		Name: "Server Hello",
		Text: "Test",
	}
	//pb := proto.Buffer{}
	pb, err := proto.Marshal(&h)
	errors.Wrap(err, "pb encode")

	fmt.Println(r.RemoteAddr)

	_, err = w.Write(pb)
	//w.Write([]byte(r.RemoteAddr))
	errors.Wrap(err, "server write")
}
