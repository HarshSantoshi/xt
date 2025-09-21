package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"practice/models"
	"strconv"
)

// GetRoot is a simple handler for the root route.
func GetRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the complete API!")
}

// TextChangeHandler handles the /text/change endpoint, demonstrating the use of the new models.
func TextChangeHandler(c echo.Context) error {
	// Default values
	color := ""
	fontSize := float64(0)

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
		ButtonStyle: models.UikitStyle{
			Color:    "",
			FontSize: 0.0,
			Value:    "",
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
