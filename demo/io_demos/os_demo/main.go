package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var b []byte = make([]byte, 100)

	dir, _ := os.Getwd()

	fd, _ := os.OpenFile(dir+"/kill-listen.sh", os.O_RDONLY, 0777)

	defer fd.Close()

	fmt.Println(fd)

	for true {
		// 行缓冲
		if num, err := fd.Read(b); err != nil || num < 1 {
			log.Println(err)
			break
		}

		fmt.Println(string(b))
	}

	uInt := fd.Fd()
	fmt.Println(uInt)
	var fileMode = 0777
	fmt.Println(os.FileMode(fileMode).Perm())
	fmt.Println("rxw", 1<<2|1<<1|1)
	fmt.Println("rw-", 1<<2|1<<1|0)
	fmt.Println("r-w", 1<<2|0<<1|1)
	fmt.Println("r--", 1<<2|0<<1|0)
	fmt.Println("-wx", 0<<2|1<<1|1)
	fmt.Println("-w-", 0<<2|1<<1|0)
	fmt.Println("--x", 0<<2|0<<1|1)
	fmt.Println("---", 0<<2|0<<1|0)
	bin := strconv.FormatInt(6, 2)
	netmask := strconv.FormatInt(240, 2)
	fmt.Println(bin)
	fmt.Println(netmask, "-len-", len(netmask))

	fmt.Printf("%3d\n", 1)

}
