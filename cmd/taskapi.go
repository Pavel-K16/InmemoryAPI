package main

import (
	"taskapi/internal/app"
	"taskapi/internal/logger"
)

var (
	log = logger.LoggerInit()
)

func main() {
	log.Info("Starting application")
	app.Run()
}
