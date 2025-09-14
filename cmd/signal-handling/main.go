package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// contextを使った方法
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// channelを使った方法
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGUSR1)

	for {
		select {
		case <-ctx.Done():
			slog.Info("finish process")

			return
		case <-signalCh:
			slog.Debug("received channel")
		}
	}
}
