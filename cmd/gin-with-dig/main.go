package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/Haya372/go-samples/internal/di"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := di.NewDiContainer()
	if err != nil {
		panic(err)
	}

	if err := c.Invoke(func(r *gin.Engine) {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
		defer stop()
		go r.Run()
		<-ctx.Done()
		fmt.Println("shutdown server...")
	}); err != nil {
		panic(err)
	}
}
