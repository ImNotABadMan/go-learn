package main

import (
	"demo/chrome_demo/gl_csv_import"
	"sync"
)

func main() {
	//click_demo.RunChromeClick()
	//gl_csv_import.OpenChrome()

	//gl_csv_import.ReStartCsvQueue()
	var ws = sync.WaitGroup{}
	ws.Add(2)

	// 多个浏览器
	go func() {
		gl_csv_import.OpenChrome("binz", "binz123", "/gl_csv_import/csv.json")
		ws.Done()
	}()

	go func() {
		gl_csv_import.OpenChrome("crazyman", "test123", "/gl_csv_import/csv1.json")
		ws.Done()
	}()

	//for	i := 0; i < 2; i++{
	//	go func() {
	//		gl_csv_import.OpenChrome("crazyman", "test123")
	//		ws.Done()
	//	}()
	//}

	//for i := 0; i < 2; i++ {
	//	go func() {
	//		gl_csv_import.OpenChrome("crazyman", "test123")
	//		ws.Done()
	//	}()
	//}

	ws.Wait()
}
