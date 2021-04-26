module pms/pub_fetures

go 1.15

require (
	github.com/Shopify/sarama v1.28.0 // indirect
	github.com/klauspost/compress v1.12.1 // indirect
	github.com/techoner/gophp v0.2.0
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/net v0.0.0-20210423184538-5f58ad60dda6 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.7
	pms/pub_fetures/create v1.0.0
	pms/pub_fetures/tables v1.0.0
)

replace (
	pms/pub_fetures/create => ./create
	pms/pub_fetures/tables => ./tables
)
