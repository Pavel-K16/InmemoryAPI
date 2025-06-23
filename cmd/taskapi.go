package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"taskapi/internal/app"
	"taskapi/internal/logger"
)

var (
	log = logger.LoggerInit()
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	if err := app.Run(ctx); err != nil {
		log.Errorf("Error starting server: %s", err)

		stop()
	}
}
