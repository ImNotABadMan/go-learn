package main

import (
	"fmt"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(quote.HelloV3())
}

/**
//vscode
// 安装go扩展
// 工具包和源码
//
//
//
git clone https://github.com/golang/tools.git
git clone https://github.com/golang/lint.git
git clone https://github.com/golang/mode.git
git clone https://github.com/golang/xerrors.git

// GOPATH 下执行
go get -u -v github.com/mdempsky/gocode
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/sqs/goreturns
go get -u -v github.com/go-delve/delve/cmd/dlv
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/cweill/gotests/...
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/josharian/impl
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
go get -u -v github.com/haya14busa/goplay/cmd/goplay
go get -u -v github.com/godoctor/godoctor

//go get -u -v github.com/stamblerre/gocode

go install github.com/mdempsky/gocode
go install github.com/ramya-rao-a/go-outline
go install github.com/acroca/go-symbols
go install github.com/rogpeppe/godef
go install github.com/sqs/goreturns
go install github.com/go-delve/delve/cmd/dlv
go install github.com/uudashr/gopkgs/cmd/gopkgs
go get github.com/cweill/gotests/...
go get github.com/fatih/gomodifytags
go get github.com/josharian/impl
go get github.com/davidrjenni/reftools/cmd/fillstruct
go get github.com/haya14busa/goplay/cmd/goplay
go get github.com/godoctor/godoctor

//go build -o $GOPATH/bin/gocode-gomod github.com/stamblerre/gocode
*/
