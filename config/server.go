package config

import (
	"github.com/labstack/echo/v4")

func SetUpServer(e *echo.Echo){	
	if error := e.Start(":8080"); error != nil {
		e.Logger.Fatal("Failed to start server: ", error)
	}
	// e.Logger.Fatal(e.Start(":8080"))
}