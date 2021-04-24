package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pms/pub_fetures/create"
	"sync"
)

func main() {
	db := Connect()
	syncGroup := sync.WaitGroup{}
	syncGroup.Add(2)
	//cond := sync.NewCond(&sync.RWMutex{})

	go func() {
		//cond.Signal()
		task := create.GetTasks(db)
		fmt.Println(task)
		syncGroup.Done()
	}()

	go func() {
		//cond.Wait()
		taskProducts := create.GetTaskProducts(db)
		fmt.Println(taskProducts)
		fmt.Println("Name:", taskProducts.Name)
		fmt.Println("Third_product_id:", taskProducts.Third_product_id)
		syncGroup.Done()
	}()

	syncGroup.Wait()
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
