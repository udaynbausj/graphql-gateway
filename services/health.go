package services

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Health(c echo.Context) error {

	response := &map[string]interface{}{
		"status" : http.StatusOK,
		"message": "Everything is alright!",
	}

	return c.JSON(http.StatusOK, response)
}
