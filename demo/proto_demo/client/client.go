package main

import (
	"bytes"
	"demo/proto_demo/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"net/http"
)

func main() {
	w, err := http.Get("http://192.168.10.113:9999/hello")

	buf := bytes.Buffer{}

	_, err = buf.ReadFrom(w.Body)
	errors.Wrap(err, "read")
	fmt.Println(buf.String())

	//pb := proto.Buffer{}
	//pb.SetBuf(buf.Bytes())

	pb := protobuf.Hello{}
	err = proto.Unmarshal(buf.Bytes(), &pb)
	errors.Wrap(err, "pb unmarshal")

	fmt.Println(pb.String())
	fmt.Println(pb.GetName())
	fmt.Println(pb.GetText())
	fmt.Println(pb.Name)
	fmt.Println(pb.Text)
	fmt.Println()
}
