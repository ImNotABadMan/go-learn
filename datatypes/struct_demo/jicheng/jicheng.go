package jicheng

import (
	"fmt"
)

type Log struct {
	msg string
}

type JiCheng struct {
	name string
	Log
}

func (str *Log) AddLog(name string) {
	//logger := log.Logger{}
	fmt.Println("test Log")

	str.msg = "Log " + name

	fmt.Println(str)
}

func (str *JiCheng) Add(name string) {
	str.name = name
	str.AddLog(name)
}
