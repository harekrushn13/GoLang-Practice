package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go doTask(ctx)

	select {
	case <-ctx.Done():
		time.Sleep(time.Second)
		fmt.Println("Main context done")
	}
}

func doTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context cancel : ", ctx.Err())
			return
		default:
			fmt.Println("Do Task")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
