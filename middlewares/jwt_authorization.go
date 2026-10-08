package middlewares

import (
	"fmt"
	"net/http"

	// "os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)




// JWT middleware function
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		fmt.Println("this is your auth header: ", authHeader)
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("Tm9HbNPPzW946bzJqGSv+opLIKtVqUCFSKq/NbmvgZw="), nil
		})
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims")
		}

		c.Set("username", claims["username"])
		c.Set("role", claims["aud"])
		fmt.Println("this is your data", c)
		return next(c)
	}
}
