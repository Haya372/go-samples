package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// contextを使った方法
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// channelを使った方法
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("received context")
			fmt.Println("finish process")
			return
		case <-c:
			fmt.Println("received channel")
		}
	}
}
