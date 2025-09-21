package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"practice/configs"
	"practice/routes"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize the Echo server
	e := echo.New()

	// Register all the API routes
	routes.InitRoutes(e)

	// Start the server in a goroutine
	go func() {
		if err := e.Start(configs.AppConfig.ServerPort); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Shutting down the server: %v", err)
		}
	}()

	// Graceful Shutdown: Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited gracefully")
}
