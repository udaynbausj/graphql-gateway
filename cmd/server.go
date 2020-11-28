package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/udaynbausj/graphql-gateway/constants"
	"os"
)

func main()  {

	// echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	err := e.Start(":" + os.Getenv(constants.PortKey))
	if  err != nil {
		e.Logger.Fatal(" Exception while initializing the server. Error : %s " , err.Error())
	}

}
