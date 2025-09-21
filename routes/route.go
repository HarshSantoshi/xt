package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"practice/handlers"
)

// InitRoutes sets up all the API routes for the application.
func InitRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define all GET and POST routes
	e.GET("/", handlers.GetRoot)
	e.GET("/text/change", handlers.TextChangeHandler)
	e.GET("/greet", handlers.GreetHandler)
	e.POST("/jsonPost", handlers.JsonPostHandler)

}
