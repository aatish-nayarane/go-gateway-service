package routes

import (
	"api-gateway/token_manager"

	middlewares "api-gateway/middlewares"

	"github.com/labstack/echo/v4"

	// routes "api-gateway/routes"
	"api-gateway/routing"
)

func RegisterRoutes(e *echo.Echo) {
	Apis := e.Group("/login")
	Apis.POST("", token_manager.LoginHandler)
	//  1st param for incoming traffic and 2nd param for outgoing traffic and third param for middleware authorization
	// e.GET("networks/*", routing.ProxyRequest("http://localhost:8001"), middlewares.JWTMiddleware)
	e.Any("networks/*", routing.ProxyRequest("http://10.11.2.80:8001"), middlewares.JWTMiddleware)
	// e.POST("networks/*", routing.ProxyRequest("http://localhost:8001"))
	e.Any("instance/*", routing.ProxyRequest("http://localhost:8002"), middlewares.JWTMiddleware)
}
