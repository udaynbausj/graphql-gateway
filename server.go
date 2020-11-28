package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main()  {

	// echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if err := e.Start(":3333"); err != nil {
		e.Logger.Fatal(" Exception while initializing the server. Error : %s " , err.Error())
	}


}
