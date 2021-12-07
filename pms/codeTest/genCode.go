package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

func main() {
	var (
		goCount = 5
	)
	wg := sync.WaitGroup{}
	wg.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func(index int) {
			var stdout bytes.Buffer
			cdErr := os.Chdir("/home/ubuntu/globaloutlet_backend_v2")
			if cdErr != nil {
				log.Fatal(cdErr)
			}
			shell := exec.Command("php", "artisan", "test", "--group=code-gen")
			shell.Stdout = &stdout
			err := shell.Run()

			//fmt.Println(string(index) + stdout.String())
			ioutil.WriteFile(
				"/home/ubuntu/go/code/src/pms/codeTest/code"+strconv.FormatInt(int64(index), 10)+".txt",
				stdout.Bytes(), 0644)

			if err != nil {
				log.Fatal(err)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}
