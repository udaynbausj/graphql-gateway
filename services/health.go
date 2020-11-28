package services

import (
	"github.com/labstack/echo/v4"
	"github.com/udaynbausj/graphql-gateway/types"
	"net/http"
)

func Health(c echo.Context) error {

	response := &types.Health{
		Status: http.StatusOK,
		Message: "everything is alright!!",
	}

	return c.JSON(http.StatusOK, response)
}
