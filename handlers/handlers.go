package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"practice/models"
	"strconv"
	"time"
)

// GetRoot is a simple handler for the root route.
func GetRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the complete API!")
}

// TextChangeHandler handles the /text/change endpoint, demonstrating the use of the new models.
func TextChangeHandler(c echo.Context) error {
	// Default values
	color := ""
	fontSize := 0

	// Check for query parameters and override defaults if present
	if c.QueryParam("color") != "" {
		color = c.QueryParam("color")
	}

	if c.QueryParam("fontSize") != "" {
		if size, err := strconv.ParseFloat(c.QueryParam("fontSize"), 64); err == nil {
			fontSize = size
		}
	}

	completedTag := models.OfferWidgetTag{
		Title: models.UikitText{
			Value: "Completed",
		},
		Style: models.UikitStyle{
			Color:    color,
			FontSize: fontSize,
		},
	}

	res := models.Response{
		Response: models.ResponseModel{
			Status:  "SUCCESS",
			Message: "",
			Code:    http.StatusOK,
		},
		Compete: completedTag,
	}

	// Returns the response as JSON
	return c.JSON(http.StatusOK, res)
}

// GreetHandler demonstrates handling query parameters.
func GreetHandler(c echo.Context) error {
	message := c.QueryParam("message")
	name := c.QueryParam("name")

	if message == "" || name == "" {
		return c.String(http.StatusBadRequest, "Both 'message' and 'name' query parameters are required.")
	}

	responseMessage := fmt.Sprintf("%s, %s!", message, name)
	return c.String(http.StatusOK, responseMessage)
}

// JsonPostHandler demonstrates handling a JSON POST request.
func JsonPostHandler(c echo.Context) error {
	greeting := new(models.Greeting)
	if err := c.Bind(greeting); err != nil {
		return c.String(http.StatusBadRequest, "Invalid JSON data")
	}

	greeting.Time = time.Now().Format(time.RFC3339)

	return c.JSON(http.StatusOK, greeting)
}
