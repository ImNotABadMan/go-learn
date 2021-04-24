package main

import (
	"fmt"
	"github.com/techoner/gophp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pms/pub_fetures/create"
	"pms/pub_fetures/tables"
	"strconv"
	"sync"
)

func main() {
	var task tables.Task
	var taskProducts tables.TaskProduct
	db := Connect()
	syncGroup := sync.WaitGroup{}
	syncGroup.Add(2)
	//cond := sync.NewCond(&sync.RWMutex{})

	go func(taskIn tables.Task) {
		//cond.Signal()
		taskIn = create.GetTasks(db)
		fmt.Println(task)
		syncGroup.Done()
	}(task)

	go func(taskProducts tables.TaskProduct) {
		//cond.Wait()
		taskProducts = create.GetTaskProducts(db)
		fmt.Println(taskProducts)
		fmt.Println("Name:", taskProducts.Name)
		fmt.Println("Third_product_id:", taskProducts.Third_product_id)
		syncGroup.Done()
	}(taskProducts)
	syncGroup.Wait()

	// 使用php 序列化发送到kafka中
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

func Connect() *gorm.DB {
	dsn := "root:123@tcp(127.0.0.1:3306)/pms?charset=utf8&parseTime=True&loc=Local"
	//dsn := "root:crazy888@tcp(192.168.10.108:3306)/globaloutletcom?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
