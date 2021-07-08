package main

import (
	"encoding/json"
	"fmt"
	"github.com/techoner/gophp"
	"github.com/tidwall/pretty"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"pms/pub_fetures/create"
	"pms/pub_fetures/kafka"
	"pms/pub_fetures/tables"
	"strconv"
	"sync"
)

func main() {
	db := Connect()

	pubIDs := createTask(db, 56556)
	if len(pubIDs) == 0 {
		PrettyPrint("没有可创建的刊登")
		return
	}

	// 发送到kafka
	for _, pubID := range pubIDs {
		task, taskProducts := getTask(db, pubID, 0)
		sendToKafka(task, taskProducts)
	}
}

func Connect() *gorm.DB {
	dsn := "root:123@tcp(172.22.0.2:3306)/globaloutletcom?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:crazy888@tcp(192.168.10.108:3306)/globaloutletcom?charset=utf8&parseTime=True&loc=Local"
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

func createTask(db *gorm.DB, minProductID int, maxProductID ...int) []int {
	insertedPubID := []int{}
	var (
		products   []tables.Products
		insertedID int
	)

	if len(maxProductID) > 0 && maxProductID[0] > minProductID {
		products = create.GetProducts(db, minProductID, maxProductID[0])
	} else {
		products = create.GetProducts(db, minProductID)
	}
	//sites := create.GetSites(db)
	categories := create.GetSiteCategory(db)
	//PrettyPrint(categories)
	//PrettyPrint(sites)
	//PrettyPrint(products)

	for _, product := range products {
		//siteID := sites[rand.Intn(len(sites))].Id
		siteID := 6
		fmt.Println("Product: ", product.ProductCode)
		if create.GetSiteProducts(db, product.ProductID, siteID) {
			log.Println("此产品在站点上已经刊登", siteID, " -- ", product.ProductID)
			continue
		}
		insertedID = create.InsertTask(db, product, categories, siteID)
		insertedPubID = append(insertedPubID, insertedID)
		fmt.Println(insertedPubID)
	}

	return insertedPubID
}

func getTask(db *gorm.DB, pubID int, maxPubID int) (tables.PublicationTask, tables.PublicationProducts) {
	var (
		task         tables.PublicationTask
		taskProducts tables.PublicationProducts
	)

	syncGroup := sync.WaitGroup{}
	syncGroup.Add(2)
	//cond := sync.NewCond(&sync.RWMutex{})

	go func(taskIn *tables.PublicationTask) {
		//cond.Signal()
		taskList := create.GetTasks(db, pubID, maxPubID)[0]
		*taskIn = taskList
		PrettyPrint("task:")
		fmt.Println("task:", taskIn)

		syncGroup.Done()
	}(&task)

	go func(taskProducts *tables.PublicationProducts) {
		//cond.Wait()
		*taskProducts = create.GetTaskProducts(db, pubID)
		//PrettyPrint(taskProducts)
		fmt.Println("Name:", taskProducts.Name)
		fmt.Println("ThirdProductID:", taskProducts.ThirdProductID)
		syncGroup.Done()
	}(&taskProducts)
	syncGroup.Wait()

	return task, taskProducts
}

// 使用php 序列化发送到kafka中
func sendToKafka(task tables.PublicationTask, taskProducts tables.PublicationProducts) {
	var (
		phpMap = map[string]string{
			"task_id":    strconv.FormatInt(int64(task.ID), 10),
			"site_id":    strconv.FormatInt(int64(task.SiteID), 10),
			"admin_id":   strconv.FormatInt(int64(task.AdminID), 10),
			"created_at": task.CreatedAt.String(),
		}
	)

	PrettyPrint(phpMap)

	phpSerialize, _ := gophp.Serialize(phpMap)

	fmt.Println(string(phpSerialize))

	kafkaQueue := kafka.Kafka{}

	kafkaQueue.Producer().Produce(kafka.TopicShopifyPublish, []string{string(phpSerialize)})
}

func PrettyPrint(v interface{}) {
	body, _ := json.MarshalIndent(v, "", "\t")
	log.Printf("%s\n", pretty.Color(body, pretty.TerminalStyle))
}
