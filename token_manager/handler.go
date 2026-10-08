package token_manager

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}




func LoginHandler(c echo.Context) error {
	var login Login
	
	if err := c.Bind(&login); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	// Validate user credentials (this is just a placeholder, implement your own logic)
	if login.Username != "senior" || login.Password != "aatish" {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}

	// Generate JWT token
	tokenString, err := CreateToken(login.Username)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create token")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}
