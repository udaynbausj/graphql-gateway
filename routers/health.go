package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/udaynbausj/graphql-gateway/constants"
	"github.com/udaynbausj/graphql-gateway/services"
)

func InitializeHealthRouter(e * echo.Echo)  {

	e.GET(constants.HealthRoutePath, services.Health)

}