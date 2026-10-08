package middlewares

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetUpEntryMiddlewares(e *echo.Echo) {
	e.Use(middleware.Recover())
	//Initializing validator instance globally at the time of startup
	// validator.InitValidator()
	e.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			// Process request
			err := next(c)

			// Log response details
			status := c.Response().Status
			duration := time.Since(start)

			fmt.Printf("[RESPONSE] %s %s | Status: %d | Duration: %v\n", c.Request().Method, c.Request().URL.Path, status, duration)

			return err
		}
	}))
}
