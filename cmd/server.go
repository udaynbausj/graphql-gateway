package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/udaynbausj/graphql-gateway/constants"
	"github.com/udaynbausj/graphql-gateway/routers"
	"os"
)

func main()  {

	// echo instance
	e := echo.New()

	// middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routers.InitializeHealthRouter(e)

	err := e.Start(":" + os.Getenv(constants.PortKey))
	if  err != nil {
		e.Logger.Fatal(" Exception while initializing the server. Error : %s " , err.Error())
	}

}
