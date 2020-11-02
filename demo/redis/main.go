package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:16379",
	})
	strSlice := rdb.Keys(ctx, "*")
	fmt.Println(strSlice)
}
