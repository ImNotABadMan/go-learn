package create

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"pms/pub_fetures/tables"
	"time"
)

func GetTasks(db *gorm.DB, id ...int) tables.Publication_task {
	var (
		task   tables.Publication_task
		taskID int
	)
	if len(id) > 0 {
		taskID = id[0]
	} else {
		taskID = 136
	}
	db.Raw("select * from ss_publication_task where id = ?", taskID).Scan(&task)

	return task
}

func GetTaskProducts(db *gorm.DB) tables.Publication_products {
	var taskProduct tables.Publication_products
	db.Raw("select * from ss_publication_products where publication_id = ?", 136).Scan(&taskProduct)

	return taskProduct
}

// 获取产品
func GetProducts(db *gorm.DB, minProductID int, maxProductID int) []tables.Products {
	var (
		products []tables.Products
	)

	db.Where("ss_products.productID between ? and ?", minProductID, maxProductID).
		Joins("Additional").Preload("Prices").Find(&products)

	//fmt.Println(len(products))

	return products
}

func GetSiteProducts(db *gorm.DB, productID int, siteID int) bool {
	type test struct {
		id int
	}
	var tmpRes = test{}

	db.Raw("select * from ss_publication_products "+
		"join ss_publication_task on ss_publication_task.id = ss_publication_products.publication_id "+
		" where productID = ? and ss_publication_task.site_id = ? ", productID, siteID).Scan(&tmpRes)

	//fmt.Println(len(products))

	return tmpRes.id > 0
}

// 获取全部站点
func GetSites(db *gorm.DB) []tables.Site_config {
	var (
		sites []tables.Site_config
	)

	db.Find(&sites)

	//fmt.Println(len(products))

	return sites
}

// 获取全部分类
func GetSiteCategory(db *gorm.DB) map[int][]tables.Site_category {
	var (
		categories    []tables.Site_category
		resCategories = make(map[int][]tables.Site_category)
	)

	db.Where("type != ?", "smart").Find(&categories)

	for _, category := range categories {
		_, ok := resCategories[category.Site_id]
		if ok {
			resCategories[category.Site_id] = append(resCategories[category.Site_id], category)
		} else {
			resCategories[category.Site_id] = []tables.Site_category{category}
		}
	}

	//fmt.Println(len(products))

	return resCategories
}

func InsertTask(db *gorm.DB, product tables.Products, categories map[int][]tables.Site_category, siteID int) int {
	task := tables.Publication_task{
		Site_id:    siteID,
		Admin_id:   1221,
		Updated_at: time.Now(),
		Created_at: time.Now(),
	}

	result := db.Debug().Create(&task)

	if result.Error != nil {
		fmt.Println(result.Error)
		return 0
	}

	insertProduct(db, product, categories, task.Id)

	return task.Id
}

func insertProduct(db *gorm.DB, product tables.Products, categories map[int][]tables.Site_category, taskID int) {
	var (
		cateMap = make(map[int]int)
	)

	currencyID := rand.Intn(len(product.Prices))

	pubProduct := tables.Publication_products{
		Publication_id: taskID,
		ProductID:      product.ProductID,
		Name:           product.Name,
		Is_logistics:   product.Additional.Out_stock_strategy,
		Is_taxable:     rand.Intn(2),
		Is_group:       product.Isgroup,
		Product_code:   product.Product_code,
		Currency_id:    product.Prices[currencyID].CurrencyID,
		Created_at:     time.Now(),
		Updated_at:     time.Now(),
	}

	result := db.Debug().Create(&pubProduct)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}

	task := GetTasks(db)

	// 插入category
	cateCount := rand.Intn(len(categories[task.Site_id]))

	for i := 0; i < cateCount; i++ {
		index := rand.Intn(len(categories[task.Site_id]))
		_, ok := cateMap[categories[task.Site_id][index].Id]
		if ok {
			i--
		} else {
			cateMap[categories[task.Site_id][index].Id] = categories[task.Site_id][index].Id
		}
	}

	for _, randCate := range cateMap {
		db.Debug().Create(&tables.Publication_product_category{
			Publication_id:   taskID,
			Site_category_id: randCate,
			Created_at:       time.Now(),
			Updated_at:       time.Now(),
		})
	}

	insertSub(db, product, taskID)

}

func insertSub(db *gorm.DB, product tables.Products, taskID int) {
	if product.Isgroup == 0 {
		sub := tables.Publication_product_sub{
			Publication_id:     taskID,
			Co_productID:       product.ProductID,
			Co_product_code:    product.Product_code,
			Price:              product.Prices[0].Price,
			Original_price:     product.Prices[0].Original_price,
			In_stock:           product.In_stock,
			Out_stock_strategy: product.Additional.Out_stock_strategy,
			Created_at:         time.Now(),
			Updated_at:         time.Now(),
		}

		db.Debug().Create(&sub)
		return
	}

	var subs []tables.Products

	db.Where("group_parent = ?", product.ProductID).
		Preload("Additional").Preload("Prices").Find(&subs)

	for _, sub := range subs {
		insertSub := tables.Publication_product_sub{
			Publication_id:     taskID,
			Co_productID:       sub.ProductID,
			Co_product_code:    sub.Product_code,
			Price:              sub.Prices[0].Price,
			Original_price:     sub.Prices[0].Original_price,
			In_stock:           sub.In_stock,
			Out_stock_strategy: sub.Additional.Out_stock_strategy,
			Created_at:         time.Now(),
			Updated_at:         time.Now(),
		}
		db.Debug().Create(&insertSub)

		db.Model(&tables.Products{}).
			Where("taskID = ?", taskID).Update("currency_id", sub.Prices[0].CurrencyID)
	}

}
