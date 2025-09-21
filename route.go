package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitRoutes sets up all the API routes for the application.
func InitRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define all GET and POST routes
	e.GET("/", GetRoot)
	e.GET("/text/change", TextChangeHandler)
	e.GET("/greet", GreetHandler)
	e.POST("/jsonPost", JsonPostHandler)

}
