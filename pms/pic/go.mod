module pms/pic

go 1.15

require (
	gorm.io/gorm v1.21.7
	gorm.io/driver/mysql v1.0.5
	pms/pic/tables v1.0.0
)

replace pms/pic/tables => ./tables
