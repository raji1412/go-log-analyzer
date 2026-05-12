package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"loganalyzer/internal/analyzer"
	"loganalyzer/internal/api/middleware"
	"loganalyzer/internal/config"
	"loganalyzer/internal/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SuccessResponse struct {
	Status string         `json:"status"`
	Data   map[string]int `json:"data"`
}

type ErrorResponse struct {
	Status       string `json:"status"`
	ErrorMessage string `json:"error"`
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "log-analyzer",
	})
}

func LogCountingHandle(c *gin.Context) {
	file := c.Query("file")
	level := c.Query("level")
	if file == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Status: "error", ErrorMessage: "file is empty"})
		return
	}
	result := analyzer.LogCounter(file, level)
	c.JSON(http.StatusOK, SuccessResponse{Status: "success", Data: result})
}

func StartServer() {
	router := gin.Default()
	router.Use(middleware.LogMiddleware())
	router.GET("/health", healthCheck)
	cfg := config.LoadConfig()
	router.GET("/logs", LogCountingHandle)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	go func() {
		logger.Logger.Info("Server Running", zap.String("port", cfg.Port))

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error in listening to server!!")
		}
	}()

	stopSignal := make(chan os.Signal, 1)

	signal.Notify(stopSignal, os.Interrupt)
	<-stopSignal

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error in shutting down!!")
	}

	logger.Logger.Info("Server exited properly")

}
