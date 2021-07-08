package create

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"pms/pub_fetures/tables"
	"time"
)

func GetTasks(db *gorm.DB, id ...int) []tables.PublicationTask {
	var rows *sql.Rows
	var (
		task   tables.PublicationTask
		taskID int
	)
	taskList := []tables.PublicationTask{}

	if len(id) > 0 && id[0] > 0 {
		taskID = id[0]
	} else {
		taskID = 1
	}
	if len(id) > 0 && id[0] > 0 {
		rows, _ = db.Raw("select * from ss_publication_task where id between ? and ?", taskID, id[0]).Rows()
	} else {
		rows, _ = db.Raw("select * from ss_publication_task where id = ?", taskID).Rows()
	}
	defer rows.Close()

	for rows.Next() {
		db.ScanRows(rows, &task)
		taskList = append(taskList, task)
	}

	return taskList
}

func GetTaskProducts(db *gorm.DB, pubID int) tables.PublicationProducts {
	var taskProduct tables.PublicationProducts
	db.Raw("select * from ss_publication_products where publication_id = ?", pubID).Scan(&taskProduct)

	return taskProduct
}

// 获取产品
func GetProducts(db *gorm.DB, minProductID int, maxProductID ...int) []tables.Products {
	var (
		products []tables.Products
	)

	if len(maxProductID) > 0 {
		db.Where("ss_products.productID between ? and ?", minProductID, maxProductID[0]).
			Where("ss_products.group_parent is Null").
			Joins("Additional").Preload("Prices").Find(&products)
	} else {
		db.Where("ss_products.productID >= ?", minProductID).
			Where("ss_products.group_parent is Null").Limit(2).
			Joins("Additional").Preload("Prices").Find(&products)
	}

	//fmt.Println(len(products))

	return products
}

func GetSiteProducts(db *gorm.DB, productID int, siteID int) bool {
	type test struct {
		ID int
	}

	var tmpRes = test{}

	db.Raw("select * from ss_publication_products "+
		"join ss_publication_task on ss_publication_task.id = ss_publication_products.publication_id "+
		" where productID = ? and ss_publication_task.site_id = ? limit 1", productID, siteID).Scan(&tmpRes)

	return tmpRes.ID > 0
}

// 获取全部站点
func GetSites(db *gorm.DB) []tables.SiteConfig {
	var (
		sites []tables.SiteConfig
	)

	db.Find(&sites)

	//fmt.Println(len(products))

	return sites
}

// 获取全部分类
func GetSiteCategory(db *gorm.DB) map[int][]tables.SiteCategory {
	var (
		categories    []tables.SiteCategory
		resCategories = make(map[int][]tables.SiteCategory)
	)

	db.Where("type != ?", "smart").Find(&categories)

	for _, category := range categories {
		_, ok := resCategories[category.SiteID]
		if ok {
			resCategories[category.SiteID] = append(resCategories[category.SiteID], category)
		} else {
			resCategories[category.SiteID] = []tables.SiteCategory{category}
		}
	}

	//fmt.Println(len(products))

	return resCategories
}

func InsertTask(db *gorm.DB, product tables.Products, categories map[int][]tables.SiteCategory, siteID int) int {
	task := tables.PublicationTask{
		SiteID:    siteID,
		AdminID:   1221,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	result := db.Create(&task)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	insertProduct(db, product, categories, task.ID)

	return task.ID
}

func insertProduct(db *gorm.DB, product tables.Products, categories map[int][]tables.SiteCategory, taskID int) {
	var (
		cateMap = make(map[int]int)
	)

	currencyID := rand.Intn(len(product.Prices))

	pubProduct := tables.PublicationProducts{
		PublicationID: taskID,
		ProductID:     product.ProductID,
		Name:          product.Name,
		IsLogistics:   product.Additional.OutStockStrategy,
		IsTaxable:     rand.Intn(2),
		IsGroup:       product.Isgroup,
		Enabled:       1,
		ProductCode:   product.ProductCode,
		CurrencyID:    product.Prices[currencyID].CurrencyID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result := db.Create(&pubProduct)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}

	task := GetTasks(db, taskID)[0]

	// 插入category
	cateCount := rand.Intn(len(categories[task.SiteID]))

	for i := 0; i < cateCount; i++ {
		index := rand.Intn(len(categories[task.SiteID]))
		_, ok := cateMap[categories[task.SiteID][index].Id]
		if ok {
			i--
		} else {
			cateMap[categories[task.SiteID][index].Id] = categories[task.SiteID][index].Id
		}
	}

	for _, randCate := range cateMap {
		db.Create(&tables.PublicationProductCategory{
			PublicationID:  taskID,
			SiteCategoryID: randCate,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		})
	}

	insertSub(db, product, taskID)

}

func insertSub(db *gorm.DB, product tables.Products, taskID int) {
	if product.Isgroup == 0 {
		sub := tables.PublicationProductSub{
			PublicationID:    taskID,
			CoProductID:      product.ProductID,
			CoProductCode:    product.ProductCode,
			Price:            product.Prices[0].Price,
			OriginalPrice:    product.Prices[0].OriginalPrice,
			InStock:          product.InStock,
			OutStockStrategy: product.Additional.OutStockStrategy,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}

		db.Create(&sub)
		return
	}

	var subs []tables.Products

	db.Where("group_parent = ?", product.ProductID).
		Preload("Additional").Preload("Prices").Find(&subs)

	for _, sub := range subs {
		insertSub := tables.PublicationProductSub{
			PublicationID:    taskID,
			CoProductID:      sub.ProductID,
			CoProductCode:    sub.ProductCode,
			Price:            sub.Prices[0].Price,
			OriginalPrice:    sub.Prices[0].OriginalPrice,
			InStock:          sub.InStock,
			OutStockStrategy: sub.Additional.OutStockStrategy,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		}
		db.Create(&insertSub)
	}

}
