package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	go test(ctx)

	ctx = test(ctx)

	go func(ctx context.Context) {
		fmt.Println("go func ", ctx.Value("test"))
	}(ctx)

	time.Sleep(time.Millisecond)
}

func test(ctx context.Context) context.Context {
	returnCtx := context.WithValue(ctx, "test", "test set test value")
	fmt.Println("test", ctx.Value("test"))
	fmt.Println("test", returnCtx.Value("test"))

	return returnCtx
}
