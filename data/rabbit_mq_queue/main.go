package main

import (
	"data/rabbit_mq_queue/cp"
	"os"
	"sync"
)

func main() {
	syncWg := sync.WaitGroup{}
	syncWg.Add(1)
	argvs := os.Args[1:]

	if len(argvs) > 0 && argvs[0] == "send" {
		go func() {
			send()
			syncWg.Done()
		}()
	} else {

		go func() {
			receive()
			syncWg.Done()
		}()
	}

	syncWg.Wait()
}

func send() {
	cp.Publish()
}

func receive() {
	cp.Customer()
}
