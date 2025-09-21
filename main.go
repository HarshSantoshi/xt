package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"practice/configs"
	"practice/handlers" // Import the handlers package
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize the Echo server
	e := echo.New()

	// Add middleware directly
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define all GET and POST routes directly in main.go
	e.GET("/", handlers.GetRoot)
	e.GET("/text/change", handlers.TextChangeHandler)
	e.GET("/greet", handlers.GreetHandler)
	e.POST("/jsonPost", handlers.JsonPostHandler)

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
