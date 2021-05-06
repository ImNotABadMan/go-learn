package create

import (
	"fmt"
	"gorm.io/gorm"
	"pms/pub_fetures/tables"
)

func GetTasks(db *gorm.DB) tables.Task {
	var task tables.Task
	db.Raw("select * from ss_publication_task where id = ?", 136).Scan(&task)

	return task
}

func GetTaskProducts(db *gorm.DB) tables.TaskProduct {
	var taskProduct tables.TaskProduct
	db.Raw("select * from ss_publication_products where publication_id = ?", 136).Scan(&taskProduct)

	return taskProduct
}

func InsertTask(db *gorm.DB, minProductID int, maxProductID int) uint {
	var (
		product tables.Products
		//additional tables.Product_additional
		//price tables.Price
	)

	db.Where("ss_products.productID between ? and ?", minProductID, maxProductID).
		Joins("Additional").Preload("Prices").Find(&product)

	fmt.Println(product.Additional.Supplier_name)
	fmt.Println(product.Product_code)
	fmt.Println(product.Prices)

	return product.ProductID

}
