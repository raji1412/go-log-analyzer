package main

import (
	"loganalyzer/internal/api"
	"loganalyzer/internal/logger"
)

func main() {
	logger.Init()
	defer logger.Logger.Sync()
	logger.Logger.Info("Starting Server!!")
	api.StartServer()
}
