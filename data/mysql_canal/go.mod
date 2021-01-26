module "data/mysql_canal"

go 1.15 

require (
	data/mysql_canal/canal_client v1.0.0
	github.com/golang/protobuf v1.4.3

)

replace (
	data/mysql_canal/canal_client => ./canal_client
)
