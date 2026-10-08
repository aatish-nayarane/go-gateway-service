package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

var allowedIPs = make(map[string]bool)

func LoadAllowedIPs(ips []string) {
	for _, ip := range ips {
		allowedIPs[ip] = true
		fmt.Printf("Allowed IP: %s\n", ip)
		fmt.Println("Allowed IPs:", allowedIPs)
	}
}

func IpRateLimiting(ratePerSec int, burst int, expiresIn time.Duration) echo.MiddlewareFunc {

	store := middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{
			Rate:      rate.Limit(ratePerSec),
			Burst:     burst,
			ExpiresIn: expiresIn,
		})
	return middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: store,
		IdentifierExtractor: func(c echo.Context) (string, error) {
			ip := c.RealIP()
			if allowedIPs[ip] {
				return ip, nil // Rate limit this IP
			}
			fmt.Println("identifiew exteraptor:", ip)

			return "", nil // Skip rate limit for others
		},
		Skipper: func(c echo.Context) bool {
			ip := c.RealIP()
			fmt.Println("test ip skipper:", ip)
			return !allowedIPs[ip] // Skip middleware if IP not in allowed list
		},
	})
}

func UserRateLimiting(ratePerSec int, burst int, expiresIn time.Duration) echo.MiddlewareFunc {

	store := middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{
			Rate:      rate.Limit(ratePerSec),
			Burst:     burst,
			ExpiresIn: expiresIn,
		})
	return middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Store: store,
		IdentifierExtractor: func(c echo.Context) (string, error) {
			user := c.Get("user").(*jwt.Token)
			fmt.Println("identifier user details:", user)
			claims := user.Claims.(jwt.MapClaims)
			fmt.Println("identifier user claims:", claims)
			username, ok := claims["username"].(string)
			if !ok {
				return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
			}

			return username, nil // Use username as rate limit identifier
		},
		Skipper: func(c echo.Context) bool {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
		
			role, _ := claims["role"].(string)
			return role == "admin" // skip rate limiting for admins
		},
	})
}
