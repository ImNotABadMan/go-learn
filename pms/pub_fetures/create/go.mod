module pms/pub_fetures/create

go 1.15

require (
	gorm.io/gorm v1.21.7
	gorm.io/driver/mysql v1.0.5
	pms/pub_fetures/tables v1.0.0
)

replace pms/pub_fetures/tables => ../tables
