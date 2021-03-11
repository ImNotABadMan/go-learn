package main

import (
	"demo/chrome_demo/gl_csv_import"
	"fmt"
	"os"
	"sync"
)

func main() {
	//click_demo.RunChromeClick()
	//gl_csv_import.OpenChrome()

	//gl_csv_import.ReStartCsvQueue()
	path := gl_csv_import.GetCsvPath("/gl_csv_import/csv.json")
	fmt.Println(path)
	os.Exit(0)

	var ws = sync.WaitGroup{}
	ws.Add(2)

	// 多个浏览器
	go func() {
		gl_csv_import.OpenChrome("binz", "binz123", "/gl_csv_import/csv.json")
		ws.Done()
	}()

	go func() {
		gl_csv_import.OpenChrome("crazyman", "test123", "/gl_csv_import/csv.json")
		ws.Done()
	}()

	//for	i := 0; i < 2; i++{
	//	go func() {
	//		gl_csv_import.OpenChrome("crazyman", "test123")
	//		ws.Done()
	//	}()
	//}

	ws.Wait()
}
