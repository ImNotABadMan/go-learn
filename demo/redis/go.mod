module demo/redis

go 1.15

require github.com/go-redis/redis/v8 v8.3.3 // indirect

replace (
	github.com/go-redis/redis/v8 => ../../../pkg/mod/github.com/go-redis/redis/v8@v8.3.3
)