package create

import (
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
