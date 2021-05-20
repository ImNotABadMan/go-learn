package main

import (
	"fmt"
	"github.com/techoner/gophp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"math/rand"
	"pms/pub_fetures/create"
	"pms/pub_fetures/tables"
	"strconv"
	"sync"
)

func main() {
	db := Connect()
	//createTask(db)
	task, taskProducts := getTask(db)
	sendToKafka(task, taskProducts)
}

func Connect() *gorm.DB {
	//dsn := "root:123@tcp(127.0.0.1:3306)/pms?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:crazy888@tcp(192.168.10.108:3306)/globaloutletcom?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ss_",
			SingularTable: true,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createTask(db *gorm.DB) {
	products := create.GetProducts(db, 73532, 73532)
	sites := create.GetSites(db)
	categories := create.GetSiteCategory(db)
	//fmt.Println(categories)
	//fmt.Println(sites)

	for _, product := range products {
		siteID := sites[rand.Intn(len(sites))].Id
		fmt.Println("Product: ", product.Product_code)
		if create.GetSiteProducts(db, product.ProductID, siteID) {
			log.Println("此产品在站点上已经刊登", siteID, " -- ", product.ProductID)
			continue
		}
		create.InsertTask(db, product, categories, siteID)
	}
}

func getTask(db *gorm.DB) (tables.Publication_task, tables.Publication_products) {
	var task tables.Publication_task
	var taskProducts tables.Publication_products

	syncGroup := sync.WaitGroup{}
	syncGroup.Add(2)
	//cond := sync.NewCond(&sync.RWMutex{})

	go func(taskIn *tables.Publication_task) {
		//cond.Signal()
		*taskIn = create.GetTasks(db, 149)
		fmt.Println("task:", taskIn)
		syncGroup.Done()
	}(&task)

	go func(taskProducts *tables.Publication_products) {
		//cond.Wait()
		*taskProducts = create.GetTaskProducts(db)
		//fmt.Println(taskProducts)
		fmt.Println("Name:", taskProducts.Name)
		fmt.Println("Third_product_id:", taskProducts.Third_product_id)
		syncGroup.Done()
	}(&taskProducts)
	syncGroup.Wait()

	return task, taskProducts
}

// 使用php 序列化发送到kafka中
func sendToKafka(task tables.Publication_task, taskProducts tables.Publication_products) {
	var phpMap = map[string]string{
		"task_id":    strconv.FormatInt(int64(task.Id), 10),
		"site_id":    strconv.FormatInt(int64(task.Site_id), 10),
		"admin_id":   strconv.FormatInt(int64(task.Admin_id), 10),
		"created_at": task.Created_at.String(),
	}
	fmt.Println(phpMap)

	phpSerialize, _ := gophp.Serialize(phpMap)

	fmt.Println(string(phpSerialize))

}
