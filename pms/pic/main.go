package main

import (
	"crypto/tls"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"pms/pic/tables"
	"sync"
)

var baseUrl string = "https://admin.globaloutlet.com/products_pictures/"
var basePath string = "/home/ubuntu/globaloutlet_backend_v2/public/products_pictures/"

func main() {
	var (
		procssCount int64 = 0
		limit       int64 = 2000
		productID   int64 = 56414
		//productID int64 = 155
		syncGroup = sync.WaitGroup{}
	)

	db := Connect()
	count := tables.CountPic(db, productID)
	times := math.Ceil(float64(count / limit))

	syncGroup.Add(int(times))

	fmt.Println("count", count)

	// 获取productIDs 组合
	for procssCount < count {
		procssCount += limit
		rows := tables.GetAll(db, limit, productID)
		fmt.Println("len rows", len(rows), "now product", procssCount, "ProductID", productID)
		productID = int64(rows[len(rows)-1].ProductID)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpCli := &http.Client{
			Transport:     tr,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		}

		go func() {
			for _, row := range rows {
				fmt.Println("url:", baseUrl+row.Filename)
				download(httpCli, row.Filename)
				fmt.Println("url:", baseUrl+row.Thumbnail)
				download(httpCli, row.Thumbnail)
				fmt.Println("url:", baseUrl+row.Smallthumbnail)
				download(httpCli, row.Smallthumbnail)
				fmt.Println("url:", baseUrl+row.Newthumbnail)
				download(httpCli, row.Newthumbnail)
			}
			syncGroup.Done()
		}()

	}

	syncGroup.Wait()

	fmt.Println("Finish")
	//file, err := os.Create(basePath + "/test/test")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(file.Name())

}

func download(httpCli *http.Client, filePath string) {
	// 跳过证书验证

	res, err := httpCli.Get(baseUrl + filePath)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer httpCli.CloseIdleConnections()
	absPath := basePath + filePath
	dir := path.Dir(absPath)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create(absPath)
	if err != nil {
		fmt.Println("create file err ")
		fmt.Println(err)
	} else {
		fmt.Println("create", file.Name())
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	// 需要关闭 res.body
	defer res.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

}

func Connect() *gorm.DB {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/pms?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:crazy888@tcp(192.168.10.108:3306)/globaloutletcom?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
